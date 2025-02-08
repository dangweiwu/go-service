package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleCreate struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewRoleCreate doc
// @tags 4-系统-角色管理
// @summary 创建角色
// @router /api/role [post]
// @security		ApiKeyAuth
// @param body body rolemodel.RoleForm true "用户信息"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=账号已存在"
func NewRoleCreate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleCreate{ginx.New(c), c, appctx}
}

func (this *RoleCreate) Do() error {
	//数据源
	po := &rolemodel.RoleForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	if err := roleservice.NewRoleService(this.appctx).Save(po); err != nil {
		return err
	}
	this.RepOk()
	return nil
}
