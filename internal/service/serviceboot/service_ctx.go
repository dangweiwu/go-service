package serviceboot

import "go-service/internal/bootstrap/basectx"

/*
- @Author: dang
*/
type ServiceCtx struct {
}

// 所有服务在这完成依赖注入
func Start(basectx *basectx.BaseCtx) (*ServiceCtx, error) {
	return &ServiceCtx{}, nil
}
