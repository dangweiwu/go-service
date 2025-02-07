package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/auth/authmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type AuthDel struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAuthDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthDel{ginx.New(c), c, appctx}
}

// Do
// @api | auth | 4 | 删除权限
// @path | /api/auth/:id
// @method | DELETE
// @headers 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t type
// @responses | ginx.ErrResponse | 500 RESPONSE
// @tbtitle  | msg 数据
// @tbrow    |n msg |e 该权限下包含其他权限，禁止删除！
// @tbrow    |n msg |e 记录不存在
func (this *AuthDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	if err := this.Delete(id); err != nil {
		return err
	} else {
		this.RepOk()
		return nil
	}
}

func (this *AuthDel) Delete(id int64) error {
	db := this.appctx.Db
	po := &authmodel.AuthPo{}
	po.ID = id
	if r := db.Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	ct := int64(0)
	if r := db.Model(&authmodel.AuthPo{}).Where("parent_id=?", id).Count(&ct); r.Error != nil {
		return r.Error
	}

	if ct != 0 {
		return errors.New("该权限下包含其他权限，禁止删除！")
	}

	if r := db.Delete(po); r.Error != nil {
		return r.Error
	}
	return nil
}
