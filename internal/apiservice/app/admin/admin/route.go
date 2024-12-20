package admin

import (
	"go-service/internal/apiservice/app/admin/admin/api"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

// @group | admin | 2 | 用户管理 | 系统用户管理 增删查改
func Route(appctx *appctx.AppCtx, r *router.Router) {
	r.Auth.POST("/admin", router.Do(appctx, api.NewAdminCreate))
	r.Auth.GET("/admin", router.Do(appctx, api.NewAdminQuery))
	r.Auth.PUT("/admin/:id", router.Do(appctx, api.NewAdminUpdate))
	r.Auth.DELETE("/admin/:id", router.Do(appctx, api.NewAdminDel))
	r.Auth.PUT("/admin/resetpwd/:id", router.Do(appctx, api.NewResetPassword))
}
