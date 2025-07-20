package api

import (
	"fmt"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type HelloGoService2 struct {
	ginx.Giner
	ctx *appctx.AppCtx
}

func NewHelloGoService2(ctx *appctx.AppCtx, gctx *gin.Context) router.Handler {
	return &HelloGoService2{
		Giner: ginx.New(gctx),
		ctx:   ctx,
	}
}

func (this *HelloGoService2) Do() error {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := time.Duration(100+r.Intn(900)) * time.Millisecond
	time.Sleep(delay)

	statusCodes := []int{200, 500}
	randomStatus := statusCodes[rand.Intn(len(statusCodes))]
	if randomStatus == 200 {
		this.Rep(fmt.Sprintf("hello go-service delay %d", delay))
	} else {
		return fmt.Errorf("hello go-service delay %d", delay)
	}

	return nil
}
