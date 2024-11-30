package app

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/hellogoservice"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

var routes = []func(actx *appctx.AppCtx, r *router.Router){
	hellogoservice.Router,
}

func InitRouter(actx *appctx.AppCtx, g *gin.Engine) {
	r := router.NewRouter(actx, g)
	for _, f := range routes {
		f(actx, r)
	}
}
