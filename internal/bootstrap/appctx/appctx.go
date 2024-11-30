package appctx

// api与grpc依赖
import (
	"context"
	"go-service/internal/bootstrap/basectx"
	"go-service/internal/config"
	"go-service/internal/service/serviceboot"
)

type AppCtx struct {
	*basectx.BaseCtx
	*serviceboot.ServiceCtx
}

func NewAppCtx(ctx context.Context, cancel context.CancelFunc, cfg config.Config) (*AppCtx, error) {
	//基础依赖
	a := &AppCtx{}
	bctx, err := basectx.BaseBoot(ctx, cancel, cfg)
	if err != nil {
		return nil, err
	}
	a.BaseCtx = bctx

	//服务依赖
	sctx, err := serviceboot.Start(bctx)
	if err != nil {
		return nil, err
	}
	a.ServiceCtx = sctx
	return a, nil
}
