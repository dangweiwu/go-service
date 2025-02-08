package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleDel struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewRoleDel doc
// @tags 4-系统-角色管理
// @summary 删除角色
// @Security		ApiKeyAuth
// @router /api/role/{id} [delete]
// @param id path int true "用户ID"
// @success 200 {object} ginx.Response{data=string} "data=ok"
func NewRoleDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleDel{ginx.New(c), c, appctx}
}

func (this *RoleDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	if err := roleservice.NewRoleService(this.appctx).Delete(id); err != nil {
		return err
	}

	this.RepOk()
	return nil

}
