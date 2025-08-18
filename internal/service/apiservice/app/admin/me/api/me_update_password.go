package api

import (
	"context"
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/service/apiservice/app/admin/me/memodel"
	"go-service/internal/service/apiservice/pkg"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/jwtx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

type UpdatePwd struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewMeUpdatePwd 修改我的密码
// @tags 1-系统-我的
// @summary 修改个人密码
// @Description	修改个人密码，修改完后token则会失效，需要进行重新登录。
// @router /api/me/password [put]
// @Security		ApiKeyAuth
// @param body body memodel.PasswordForm true "修改个人密码"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=原密码错误"
func NewMeUpdatePwd(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &UpdatePwd{ginx.New(c), c, appctx}
}

func (this *UpdatePwd) Do() error {
	var err error
	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}

	po := &memodel.PasswordForm{}

	err = this.Bind(po)
	if err != nil {
		return err
	}
	err = this.UpdatePwd(po, uid)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *UpdatePwd) UpdatePwd(form *memodel.PasswordForm, id int64) error {
	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Where("id=?", id).Take(po); r.Error != nil {
		return r.Error
	}
	pwd, err := pkg.GetPassword(form.Password)
	if err != nil {
		return err
	}

	//校验旧密码是否正确
	if pwd != po.Password {
		return errors.New("原密码错误")
	}
	pwd, err = pkg.GetPassword(form.NewPassword)
	if err != nil {
		return err
	}
	po.Password = pwd

	r := this.appctx.Db.Model(po).Select("password").Updates(po)

	this.appctx.Redis.Del(context.Background(), memodel.GetAdminRedisLoginId(int(po.ID)))
	return r.Error
}
