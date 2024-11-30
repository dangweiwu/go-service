package basectx

import (
	"context"
	"fmt"
	"github.com/dangweiwu/microkit/db/mysqlx"
	"github.com/dangweiwu/microkit/db/redisx"
	"github.com/dangweiwu/microkit/observe/logx"
	"github.com/go-redis/redis/v8"
	"go-service/internal/config"
	"go-service/internal/pkg/lg"
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
		sctx.Log = lg.NewBaseLog(_lg, "")
		sctx.ApiLog = lg.NewBaseLog(_lg, "api")
		sctx.RpcLog = lg.NewBaseLog(_lg, "rpc")
		sctx.RpcLog = lg.NewBaseLog(_lg, "ser")
	}

	if db, err := mysqlx.NewClient(cfg.Mysql); err != nil {
		return nil, fmt.Errorf("new mysqlx error :%w", err)
	} else {
		sctx.Db = db
	}

	if rd, err := redisx.NewClient(cfg.Redis); err != nil {
		return nil, fmt.Errorf("new Redisx error :%w", err)
	} else {
		sctx.Redis = rd

	}

	return sctx, nil
}
