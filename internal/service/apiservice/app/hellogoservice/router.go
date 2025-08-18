package hellogoservice

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/hellogoservice/api"
	"go-service/internal/service/apiservice/router"
)

func Route(actx *appctx.AppCtx, r *router.Router) {
	r.Root.GET("/hello", router.Do(actx, api.NewHelloGoService))
	r.Root.GET("/rand", router.Do(actx, api.NewHelloGoService2))
}
