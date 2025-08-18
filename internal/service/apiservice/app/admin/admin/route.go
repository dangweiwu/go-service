package admin

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin/api"
	"go-service/internal/service/apiservice/router"
)

func Route(appctx *appctx.AppCtx, r *router.Router) {
	r.Auth.POST("/admin", router.Do(appctx, api.NewAdminCreate))
	r.Auth.GET("/admin", router.Do(appctx, api.NewAdminQuery))
	r.Auth.PUT("/admin/:id", router.Do(appctx, api.NewAdminUpdate))
	r.Auth.DELETE("/admin/:id", router.Do(appctx, api.NewAdminDel))
	r.Auth.PUT("/admin/resetpwd/:id", router.Do(appctx, api.NewResetPassword))
}
