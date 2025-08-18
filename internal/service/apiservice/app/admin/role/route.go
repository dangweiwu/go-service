package role

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/role/api"
	"go-service/internal/service/apiservice/router"
)

// @group | role | 5 | 角色管理
func Route(appctx *appctx.AppCtx, r *router.Router) {
	r.Auth.GET("/role", router.Do(appctx, api.NewRoleQuery))
	r.Auth.POST("/role", router.Do(appctx, api.NewRoleCreate))
	r.Auth.PUT("/role/:id", router.Do(appctx, api.NewRoleUpdate))
	r.Auth.DELETE("/role/:id", router.Do(appctx, api.NewRoleDel))
	r.Auth.PUT("/role/auth/:id", router.Do(appctx, api.NewSetAuth))
	r.Auth.GET("/role/:code", router.Do(appctx, api.NewRoleInfo))
}
