package option

import (
	"context"
	"go-service/internal/apiservice/app"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/config"
	"log"

	"github.com/dangweiwu/microkit/yamlconfig"
)

type InitTable struct {
}

func (*InitTable) Usage() string {
	return `//迁移数据库结构，但不会删除未使用的列。`
}

func (this *InitTable) Execute(args []string) error {
	var c config.Config
	loadEnvFile()
	yamlconfig.MustLoad(Opt.ConfigPath, &c)

	ctx, cf := context.WithCancel(context.Background())

	ctx2, err := appctx.NewAppCtx(ctx, cf, c)
	if err != nil {
		log.Panic(err)
		return err
	}
	err = app.Regdb(ctx2)
	if err != nil {
		log.Panic(err)
	}

	return nil
}
