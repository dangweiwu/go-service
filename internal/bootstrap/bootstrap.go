package bootstrap

import (
	"context"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/config"
	"go-service/internal/service"
	"os/signal"
	"syscall"
	"time"
)

// block
type BootStrap struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewBootStrap() *BootStrap {
	ctx, stopSignal := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	return &BootStrap{
		Ctx:    ctx,
		Cancel: stopSignal,
	}
}

func (this *BootStrap) Init(cfg config.Config) error {
	var err error
	//依赖注入
	actx, err := appctx.NewAppCtx(this.Ctx, this.Cancel, cfg)
	if err != nil {
		return err
	}

	//启动服务
	if err := service.Start(actx); err != nil {
		actx.Log.Msg("service start failed").ErrData(err).Err()
		return err
	}
	return nil
}

func (this *BootStrap) Run() {
	<-this.Ctx.Done()
	this.Cancel()               //取消自定义信号处理器,启用运行时信号处理
	time.Sleep(5 * time.Second) //5s等待
}
