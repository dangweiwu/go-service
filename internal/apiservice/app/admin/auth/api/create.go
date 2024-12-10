package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/auth/authmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type AuthCreate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAuthCreate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthCreate{ginx.New(c), c, appctx}
}

// Do
// @api | auth | 1 | 创建权限
// @path    | /api/auth
// @method  | POST
// @header  |n Authorization |d token |t string |c 鉴权
// @form    | authmodel.AuthForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *AuthCreate) Do() error {
	//数据源
	po := &authmodel.AuthForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	err = this.Create(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *AuthCreate) Create(po *authmodel.AuthForm) error {
	db := this.appctx.Db
	//验证是否已创建 或者其他检查
	tmpPo := &authmodel.AuthPo{}
	if r := db.Model(po).Where("code = ?", po.Code).Take(tmpPo); r.Error == nil {
		return errors.New("权限编码已存在")
	}

	if r := db.Create(po); r.Error != nil {
		return r.Error
	}
	return nil
}
