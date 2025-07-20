package basectx

import (
	"context"
	"fmt"

	"github.com/dangweiwu/microkit/casbinx"
	"github.com/dangweiwu/microkit/db/mysqlx"
	"github.com/dangweiwu/microkit/db/redisx"

	"go-service/internal/config"
	"go-service/internal/pkg/lg"

	"github.com/dangweiwu/microkit/observe/logx"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BaseCtx struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Config config.Config
	Log    *lg.BaseLog
	ApiLog *lg.BaseLog
	RpcLog *lg.BaseLog
	SerLog *lg.BaseLog
	Db     *gorm.DB
	Redis  *redis.Client
	Casbin *casbinx.CasbinxGorm
}

func BaseBoot(ctx context.Context, cf context.CancelFunc, cfg config.Config) (*BaseCtx, error) {

	sctx := &BaseCtx{
		Ctx:    ctx,
		Cancel: cf,
		Config: cfg,
	}

	if _lg, err := logx.New(cfg.Log); err != nil {
		return nil, fmt.Errorf("new logx error :%w", err)
	} else {
		sctx.Log = lg.NewBaseLog(_lg, "apiser")
		sctx.ApiLog = lg.NewBaseLog(_lg, "api")
		sctx.RpcLog = lg.NewBaseLog(_lg, "rpc")
		sctx.SerLog = lg.NewBaseLog(_lg, "ser")
	}
	//mysql实现
	if db, err := mysqlx.NewClient(cfg.Mysql); err != nil {
		return nil, fmt.Errorf("new mysqlx error :%v", err)
	} else {
		sctx.Db = db
	}

	//redis实现
	if rd, err := redisx.NewClient(cfg.Redis); err != nil {
		return nil, fmt.Errorf("new Redisx error :%w", err)
	} else {
		sctx.Redis = rd
	}

	if c, err := casbinx.NewCasbinGorm(sctx.Db); err != nil {
		return nil, fmt.Errorf("new casbinx error :%w", err)
	} else {
		sctx.Casbin = c
	}

	return sctx, nil
}
