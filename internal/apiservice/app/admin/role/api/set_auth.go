package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type SetAuth struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewSetAuth doc
// @tags 4-系统-角色管理
// @summary 设置auth列表
// @security		ApiKeyAuth
// @router /api/role/auth/{id} [put]
// @param id path int true "角色ID"
// @param body body rolemodel.RoleAuthForm true " "
// @success 200 {object} ginx.Response{data=string} "data=ok"
func NewSetAuth(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &SetAuth{ginx.New(c), c, appctx}
}

func (this *SetAuth) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	po := &rolemodel.RoleAuthForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	s := roleservice.NewRoleService(this.appctx)
	err = s.SetAuth(id, po.Auth)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}
