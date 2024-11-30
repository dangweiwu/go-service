package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type HelloGoService struct {
	*ginx.Ginx
	ctx *appctx.AppCtx
}

func NewHelloGoService(ctx *appctx.AppCtx, gctx *gin.Context) router.Handler {
	return &HelloGoService{
		Ginx: ginx.NewHd(gctx),
		ctx:  ctx,
	}
}

func (this *HelloGoService) Do() error {
	this.Rep("hello go-service")
	return nil
}
