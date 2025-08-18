package app

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin"
	"go-service/internal/service/apiservice/app/admin/auth"
	"go-service/internal/service/apiservice/app/admin/me"
	"go-service/internal/service/apiservice/app/admin/role"
	"go-service/internal/service/apiservice/app/hellogoservice"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

var routes = []func(actx *appctx.AppCtx, r *router.Router){
	hellogoservice.Route,
	admin.Route,
	me.Route,
	auth.Route,
	role.Route,
}

// InitRouter 初始化路由

func InitRouter(actx *appctx.AppCtx, g *gin.Engine) {
	r := router.NewRouter(actx, g)
	for _, f := range routes {
		f(actx, r)
	}
}
