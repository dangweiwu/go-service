package api

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/role/roleservice"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

type RoleInfo struct {
	ginx.Ginx
	appctx *appctx.AppCtx
	ctx    *gin.Context
}

// NewRoleInfo doc
// @tags 4-系统-角色管理
// @summary 角色详情
// @security		ApiKeyAuth
// @router /api/role/{code} [get]
// @param code path string true "角色编码"
// @success 200 {object} rolemodel.RolePo "角色详情"
func NewRoleInfo(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleInfo{
		Ginx:   *ginx.New(c),
		appctx: appctx,
		ctx:    c,
	}
}

func (this *RoleInfo) Do() error {
	code, err := this.GetUrlkey("code")
	if err != nil {
		return err
	}
	s := roleservice.NewRoleService(this.appctx)
	ro, err := s.GetRole(code)
	if err != nil {
		return err
	}

	this.Rep(ro)
	return nil

}
