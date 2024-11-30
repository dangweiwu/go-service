package hellogoservice

import (
	"go-service/internal/apiservice/app/hellogoservice/api"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

func Router(actx *appctx.AppCtx, r *router.Router) {
	r.Root.GET("/hello", router.Do(actx, api.NewHelloGoService))
}
