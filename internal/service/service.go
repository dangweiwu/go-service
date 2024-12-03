package service

import (
	"go-service/internal/bootstrap/basectx"
	"go-service/internal/service/pprofservice"
)

type ServiceCtx struct {
}

// 所有服务在这完成依赖注入
func Start(basectx *basectx.BaseCtx) (*ServiceCtx, error) {

	pprofservice.PprofStart(basectx)

	return &ServiceCtx{}, nil
}
