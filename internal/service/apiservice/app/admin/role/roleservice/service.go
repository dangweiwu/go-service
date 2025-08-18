package roleservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/auth/authmodel"
	"go-service/internal/service/apiservice/app/admin/role/rolemodel"
	"log"
	"time"

	"github.com/dangweiwu/microkit/casbinx"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func cacheName(name string) string {
	return "role:" + name
}

type RoleService struct {
	appctx *appctx.AppCtx
}

func NewRoleService(appctx *appctx.AppCtx) *RoleService {
	return &RoleService{appctx: appctx}
}

func (this *RoleService) GetRole(rolecode string) (*rolemodel.RolePo, error) {
	cacheKey := cacheName(rolecode)
	cacheResult, err := this.appctx.Redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var rolePo rolemodel.RolePo
		if err := json.Unmarshal([]byte(cacheResult), &rolePo); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return &rolePo, nil
	} else if err != redis.Nil {
		return nil, fmt.Errorf("failed to get from cache: %w", err)
	}

	// 缓存中没有数据，从数据库中获取
	rolePo := &rolemodel.RolePo{}
	result := this.appctx.Db.Where("code = ?", rolecode).First(rolePo)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("not found record") // 或者返回特定的错误，表示未找到记录
		}
		return nil, fmt.Errorf("failed to get from database: %w", result.Error)
	}

	// 将数据库中的数据序列化并存入缓存
	cacheData, err := json.Marshal(rolePo)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data for cache: %w", err)
	}
	if err := this.appctx.Redis.Set(context.Background(), cacheKey, cacheData, time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	return rolePo, nil

}

func (this *RoleService) UpdateRole(po *rolemodel.RoleUpdate) (err error) {
	backpo := &rolemodel.RolePo{}
	err = this.appctx.Db.Transaction(func(tx *gorm.DB) error {

		if r := tx.Model(po).Select([]string{"id", "code", "status", "auth"}).Where("id=?", po.ID).First(backpo); r.Error != nil {
			if r.Error == gorm.ErrRecordNotFound {
				return errors.New("not found record")
			} else {
				return r.Error
			}
		}
		if r := tx.Model(po).Updates(po); r.Error != nil {
			return r.Error
		}
		if err := this.appctx.Redis.Del(context.Background(), cacheName(backpo.Code)).Err(); err != nil {
			return err
		}
		if po.Status != backpo.Status {
			if po.Status == "0" {
				if _, err := this.appctx.Casbin.RemoveRolePolicy(backpo.Code); err != nil {
					return err
				}
			} else {
				if err := this._updateCasbinRule(tx, backpo.Code, backpo.Auth); err != nil {
					return err
				}
			}
		}

		return nil
	})
	return err
}

func (this *RoleService) Save(po *rolemodel.RoleForm) error {
	return this.appctx.Db.Transaction(func(tx *gorm.DB) error {
		backpo := &rolemodel.RolePo{}
		if r := tx.Model(po).Select(rolemodel.ViewFields).Where("code=?", po.Code).First(backpo); r.Error != nil {
			if r.Error == gorm.ErrRecordNotFound {
				if r := tx.Create(po); r.Error != nil {
					return r.Error
				} else {
					if err := this.appctx.Redis.Del(context.Background(), cacheName(po.Code)).Err(); err != nil {
						return fmt.Errorf("redis-err:%w", err)
					}
					log.Println("save-role-success", po.Code)
					if _, err := this.appctx.Casbin.RemoveRolePolicy(po.Code); err != nil {

						return fmt.Errorf("casbin-err:%w", err)
					}
				}
			} else {
				return r.Error
			}
		} else {
			return errors.New("角色编码已存在")
		}
		return nil
	})
}

func (this *RoleService) Delete(id int64) (err error) {
	backpo := &rolemodel.RolePo{}
	err = this.appctx.Db.Transaction(func(tx *gorm.DB) error {
		if r := tx.Model(backpo).Select(rolemodel.ViewFields).Where("id=?", id).First(backpo); r.Error != nil {
			if r.Error == gorm.ErrRecordNotFound {
				return errors.New("记录不存在")
			} else {
				return r.Error
			}
		}

		if r := tx.Model(backpo).Delete(backpo); r.Error != nil {
			return r.Error
		}

		if r := this.appctx.Redis.Del(context.Background(), cacheName(backpo.Code)); r.Err() != nil {
			return r.Err()
		}

		if _, err := this.appctx.Casbin.RemoveRolePolicy(backpo.Code); err != nil {
			return err
		}

		return nil
	})
	return err
}

// 设置权限auth
func (this *RoleService) SetAuth(roleid int64, auth []string) error {
	po := &rolemodel.RolePo{}
	po.ID = roleid
	if r := this.appctx.Db.Model(po).Select(rolemodel.ViewFields).First(po); r.Error != nil {
		return r.Error
	}
	po.Auth = auth
	err := this.appctx.Db.Transaction(func(tx *gorm.DB) error {
		if r := tx.Model(po).Select("auth").Updates(po); r.Error != nil {
			return r.Error
		}
		if r := this.appctx.Redis.Del(context.Background(), cacheName(po.Code)); r.Err() != nil {
			return r.Err()
		}

		if len(auth) == 0 {
			if _, err := this.appctx.Casbin.RemoveRolePolicy(po.Code); err != nil {
				return err
			}
		} else {

			if err := this._updateCasbinRule(tx, po.Code, auth); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (this *RoleService) _updateCasbinRule(tx *gorm.DB, code string, auth []string) error {
	authpo := &authmodel.AuthPo{}
	authpos := []authmodel.AuthPo{}
	if r := tx.Model(authpo).Select([]string{"api", "method"}).Where("kind = ? and code in ?", "0", auth).Find(&authpos); r.Error != nil {
		return r.Error
	}

	ps := []casbinx.ApiPolice{}
	for _, v := range authpos {
		ps = append(ps, casbinx.ApiPolice{Api: v.Api, Method: v.Method})
	}
	if len(ps) > 0 {
		if _, err := this.appctx.Casbin.AddPolicy(code, ps); err != nil {
			return err
		}
	} else {
		if _, err := this.appctx.Casbin.RemoveRolePolicy(code); err != nil {
			return err
		}
	}
	return nil
}
