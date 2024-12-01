package service

import "go-service/internal/bootstrap/basectx"

type ServiceCtx struct {
}

// 所有服务在这完成依赖注入
func Start(basectx *basectx.BaseCtx) (*ServiceCtx, error) {
	return &ServiceCtx{}, nil
}
