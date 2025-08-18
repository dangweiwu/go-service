package api

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

type HelloGoService struct {
	ginx.Giner
	ctx *appctx.AppCtx
}

func NewHelloGoService(ctx *appctx.AppCtx, gctx *gin.Context) router.Handler {
	return &HelloGoService{
		Giner: ginx.New(gctx),
		ctx:   ctx,
	}
}

func (this *HelloGoService) Do() error {
	this.Rep("hello go-service")
	return nil
}
