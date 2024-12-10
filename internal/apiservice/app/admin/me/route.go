package me

import (
	"go-service/internal/apiservice/app/admin/me/api"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

// @group | me | 1 | 系统我的 | 包括基本信息获取修改 登录登出 token刷新
func Route(appctx *appctx.AppCtx, r *router.Router) {
	r.Root.POST("/login", router.Do(appctx, api.NewLogin))
	r.Jwt.POST("/logout", router.Do(appctx, api.NewLogOut))
	r.TokenReflsh.POST("/token/refresh", router.Do(appctx, api.NewRefreshToken))
	r.Jwt.GET("/me", router.Do(appctx, api.NewMeInfo))
	r.Jwt.PUT("/me", router.Do(appctx, api.NewMeUpdate))
	r.Jwt.PUT("/me/password", router.Do(appctx, api.NewMeUpdatePwd))
}
