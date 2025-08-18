package api

import (
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/auth/authmodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

type AuthCreate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAuthCreate doc
// @tags 3-系统-权限管理
// @summary 创建权限
// @router /api/auth [post]
// @Security		ApiKeyAuth
// @param body body authmodel.AuthForm true "权限"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse ""
func NewAuthCreate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthCreate{ginx.New(c), c, appctx}
}

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
