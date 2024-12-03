package pprofservice

import (
	"go-service/internal/bootstrap/basectx"
	"net/http"
	_ "net/http/pprof"
)

func PprofStart(ctx *basectx.BaseCtx) {

	if ctx.Config.Pprof.Enable {
		go func() {
			err := http.ListenAndServe(ctx.Config.Pprof.Host, nil)
			if err != nil {
				ctx.SerLog.Msg("pprof err").ErrData(err).Err()
			}
		}()
	}
}

//func _pprof(r *gin.Engine) {
//	// 注册 pprof 路由
//	r.GET("/debug/pprof/", gin.WrapH(http.HandlerFunc(pprof.Index)))
//	r.GET("/debug/pprof/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
//	r.GET("/debug/pprof/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
//	r.GET("/debug/pprof/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
//	r.POST("/debug/pprof/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
//	r.GET("/debug/pprof/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
//	r.GET("/debug/pprof/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP))) // 添加堆栈分析路由
//}
