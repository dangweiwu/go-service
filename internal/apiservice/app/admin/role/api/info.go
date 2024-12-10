package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleInfo struct {
	ginx.Ginx
	appctx *appctx.AppCtx
	ctx    *gin.Context
}

func NewRoleInfo(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleInfo{
		Ginx:   *ginx.New(c),
		appctx: appctx,
		ctx:    c,
	}
}

// Do
// @api    | role | 3 | 角色详情
// @path   | /api/role/:code
// @method | GET
// @header |n Authorization |d token |t string |c 鉴权
// @url    |n code |d 角色代码 |t string |c 角色代码
// @response | rolemodel.RolePo | 200 Response
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
