package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleUpdate struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewRoleUpdate doc
// @tags 4-系统-角色管理
// @summary 修改角色
// @security		ApiKeyAuth
// @router /api/role/{id} [put]
// @param id path int true "角色ID"
// @param body body rolemodel.RoleUpdate true " "
// @success 200 {object} ginx.Response{data=string} "data=ok"
func NewRoleUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleUpdate{ginx.New(c), c, appctx}
}

func (this *RoleUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &rolemodel.RoleUpdate{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	if err := roleservice.NewRoleService(this.appctx).UpdateRole(po); err != nil {
		return err
	}

	this.RepOk()
	return nil
}
