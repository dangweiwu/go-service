package service

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice"
	"go-service/internal/service/pprofservice"
)

type ServiceCtx struct {
}

// 所有服务在这完成依赖注入
func Start(appctx *appctx.AppCtx) error {
	// //启动api服务
	apiservice.Start(appctx)
	pprofservice.PprofStart(appctx)
	return nil
}
