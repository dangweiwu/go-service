package me

import (
	"go-service/internal/apiservice/app/admin/me/api"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

// @group | me | 2 | 系统我的 | 包括基本信息获取修改 登录登出 token刷新
func Route(appctx *appctx.AppCtx, r *router.Router) {

	//r.Jwt.GET("/my", router.Do(appctx, api.NewLogin))
	//
	//r.Jwt.PUT("/my", router.Do(appctx, api.NewMyUpdate))
	//
	//r.Jwt.PUT("/my/password", router.Do(appctx, api.NewMyUpdatePwd))

	r.Root.POST("/login", router.Do(appctx, api.NewLogin))

	//r.Jwt.POST("/logout", router.Do(appctx, api.NewLogOut))
	//
	//r.Jwt.POST("/token/refresh", router.Do(appctx, api.NewRefreshToken))
	//
	//r.Jwt.GET("/my-auth", router.Do(appctx, api.NewMyAuth))
}
