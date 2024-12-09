package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type MeUpdate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewMeUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &MeUpdate{ginx.New(c), c, appctx}
}

// Do
// @api     | me | 4 | 修改我的信息
// @path 	| /api/me
// @method 	| PUT
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | memodel.MeForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *MeUpdate) Do() error {
	var err error
	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}

	po := &memodel.MeForm{}

	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = uid

	if err := this.valid(po); err != nil {
		return err
	}

	err = this.Update(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *MeUpdate) Update(rawpo *memodel.MeForm) error {
	po := &memodel.MeForm{}
	//校验
	if r := this.appctx.Db.Model(po).Where("id=?", rawpo.ID).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}

	//更新
	if r := this.appctx.Db.Model(rawpo).Select("phone", "name", "memo", "email").Updates(rawpo); r.Error != nil {
		return r.Error
	}
	return nil

}

func (this *MeUpdate) valid(po *memodel.MeForm) error {
	var ct = int64(0)

	if po.Phone != "" {
		if r := this.appctx.Db.Model(po).Where("id != ? and phone = ?", po.ID, po.Phone).Count(&ct); r.Error != nil {
			return fmt.Errorf("校验失败:%w", r.Error)
		} else if ct != 0 {
			return errors.New("手机号已存在")
		}
	}

	if po.Email != "" {
		if r := this.appctx.Db.Model(po).Where("id!=? and email = ?", po.ID, po.Email).Count(&ct); r.Error != nil {
			return fmt.Errorf("校验失败:%w", r.Error)
		} else if ct != 0 {
			return errors.New("Email已存在")
		}
	}
	return nil
}
