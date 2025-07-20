package bootstrap

import (
	"context"
	"go-service/internal/apiservice"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// block
type BootStrap struct {
	ctx    context.Context
	cancel context.CancelFunc
	quit   chan os.Signal
}

func NewBootStrap() *BootStrap {
	_ctx, closecancel := context.WithCancel(context.Background())
	return &BootStrap{
		ctx:    _ctx,
		cancel: closecancel,
		quit:   make(chan os.Signal, 1),
	}
}

func (this *BootStrap) Init(cfg config.Config) error {
	var err error
	//依赖注入
	actx, err := appctx.NewAppCtx(this.ctx, this.cancel, cfg)
	if err != nil {
		return err
	}

	//启动api服务
	err = apiservice.Start(actx)
	if err != nil {
		return err
	}
	return nil
}

func (this *BootStrap) Run() {

	signal.Notify(this.quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case <-this.ctx.Done():
		return
	case <-this.quit:
		this.cancel()
	}

	shutdownCtx, cf := context.WithTimeout(context.Background(), 2*time.Second)
	defer cf()
	select {
	case <-shutdownCtx.Done():
	case <-time.After(10 * time.Second):
	}

	os.Exit(0)
}
