package auth

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/auth/api"
	"go-service/internal/service/apiservice/router"
)

func Route(appctx *appctx.AppCtx, r *router.Router) {

	r.Auth.GET("/auth", router.Do(appctx, api.NewAuthQuery))

	r.Auth.POST("/auth", router.Do(appctx, api.NewAuthCreate))

	r.Auth.PUT("/auth/:id", router.Do(appctx, api.NewAuthUpdate))
	//
	r.Auth.DELETE("/auth/:id", router.Do(appctx, api.NewAuthDel))
	//
	r.Root.GET("/allurl", router.Do(appctx, api.NewAllUrl))

}
