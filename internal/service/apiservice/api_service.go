package apiservice

import (
	"fmt"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app"
	"go-service/internal/service/apiservice/app/doc"
	"go-service/internal/service/apiservice/middler"
	"go-service/internal/service/apiservice/pkg"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func Start(appctx *appctx.AppCtx) {
	appctx.ApiLog.Msg("init").Info()
	gin.SetMode(appctx.Config.Api.Mode)
	engine := gin.New()

	//注册所有中间件
	for _, mid := range middler.RegMiddler(appctx) {
		engine.Use(mid)
	}

	//设置静态文件
	engine.Static("/view", appctx.Config.Api.ViewDir)

	//文档
	doc.InitDoc(appctx, engine)

	//重定向到主页面
	engine.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/view"
		engine.HandleContext(c)
	})

	//注册所有路由
	app.InitRouter(appctx, engine)
	pkg.NewAllUrl(engine).InitUrl()

	server := &http.Server{
		Handler: engine,
	}

	if len(appctx.Config.Api.Domain) == 0 {
		server.Addr = appctx.Config.Api.Host
		go func() {
			if err := server.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					appctx.Log.Msg("启动服务失败").ErrData(err).Err()
				}
			}
		}()
	} else {
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,                                        // 接受 Let's Encrypt 的服务条款
			Cache:      autocert.DirCache(path.Join(appctx.Config.Root, "cache")), // 存储证书的位置
			HostPolicy: autocert.HostWhitelist(appctx.Config.Api.Domain),          // 允许的域名
			Email:      appctx.Config.Api.Email,
		}
		go func() {
			if err := server.Serve(m.Listener()); err != nil {
				if err != http.ErrServerClosed {
					appctx.Log.Msg("启动服务失败").ErrData(err).Err()
				}
			}
		}()
	}
	appctx.ApiLog.Msg("服务启动").Data(fmt.Sprintf("host:%s domain:%s", appctx.Config.Api.Host, appctx.Config.Api.Domain)).Info()

	<-appctx.Ctx.Done()
	if err := server.Shutdown(appctx.Ctx); err != nil {
		appctx.Log.Msg("api服务关闭失败").ErrData(err).Err()
	} else {
		appctx.ApiLog.Msg("api服务安全关闭").Info()
	}
}
