package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type UpdatePwd struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewMeUpdatePwd(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &UpdatePwd{ginx.New(c), c, appctx}
}

// Do
// @api     | me | 5 | 修改密码
// @path 	| /api/me/password
// @method 	| PUT
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | memodel.PasswordForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
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
