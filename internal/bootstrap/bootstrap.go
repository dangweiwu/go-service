package bootstrap

import (
	"context"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/config"
	"go-service/internal/service"
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

	//启动服务
	err = service.Start(actx)
	if err != nil {
		actx.Log.Msg("service start failed").ErrData(err).Err()
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
	<-shutdownCtx.Done()

	os.Exit(0)
}
