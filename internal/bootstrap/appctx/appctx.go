package appctx

// api与grpc依赖
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

type AppCtx struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	Config  config.Config
	Log     *lg.BaseLog
	ApiLog  *lg.BaseLog
	HttpLog *lg.BaseLog
	PproLog *lg.BaseLog
	Db      *gorm.DB
	Redis   *redis.Client
	Casbin  *casbinx.CasbinxGorm
}

func NewAppCtx(ctx context.Context, cancel context.CancelFunc, cfg config.Config) (*AppCtx, error) {
	//基础依赖
	a := &AppCtx{
		Ctx:    ctx,
		Cancel: cancel,
		Config: cfg,
	}

	if _lg, err := logx.New(cfg.Log); err != nil {
		return nil, fmt.Errorf("new logx error :%w", err)
	} else {
		a.Log = lg.NewBaseLog(_lg, "log")      //整体框架用
		a.ApiLog = lg.NewBaseLog(_lg, "api")   //api 用
		a.HttpLog = lg.NewBaseLog(_lg, "http") //api http用
		a.PproLog = lg.NewBaseLog(_lg, "ppro") //api http用

	}
	//mysql实现
	if db, err := mysqlx.NewClient(cfg.Mysql); err != nil {
		return nil, fmt.Errorf("new mysqlx error :%v", err)
	} else {
		a.Db = db
	}

	//redis实现
	if rd, err := redisx.NewClient(cfg.Redis); err != nil {
		return nil, fmt.Errorf("new Redisx error :%w", err)
	} else {
		a.Redis = rd
	}

	if c, err := casbinx.NewCasbinGorm(a.Db); err != nil {
		return nil, fmt.Errorf("new casbinx error :%w", err)
	} else {
		a.Casbin = c
	}

	return a, nil
}
