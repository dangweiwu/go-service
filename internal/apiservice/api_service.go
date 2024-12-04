package apiservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app"
	"go-service/internal/apiservice/middler"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/bootstrap/appctx"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"path"
	"sync"
)

func Start(appctx *appctx.AppCtx) (err error) {
	appctx.ApiLog.Msg("init").Info()
	gin.SetMode(appctx.Config.Api.Mode)
	engine := gin.New()

	//注册所有中间件
	for _, mid := range middler.RegMiddler(appctx) {
		engine.Use(mid)
	}

	//设置静态文件
	engine.Static("/view", appctx.Config.Api.ViewDir)
	//重定向到主页面
	engine.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/view"
		engine.HandleContext(c)
	})

	//注册所有路由
	app.InitRouter(appctx, engine)
	pkg.NewAllUrl(engine).InitUrl()
	//注册所有数据库
	err = app.Regdb(appctx)
	if err != nil {
		appctx.ApiLog.Msg("注册数据库失败").ErrData(err).Err()
		return
	}
	go func() {
		err := ginRun(appctx, engine)

		if err != nil {
			if err == http.ErrServerClosed {
				appctx.ApiLog.Msg("api服务安全关闭").Info()
			} else {
				appctx.ApiLog.Msg("启动服务失败").ErrData(err).Err()
			}
			appctx.Cancel()
		}
	}()

	appctx.ApiLog.Msg("服务启动").Data(fmt.Sprintf("host:%s domain:%s", appctx.Config.Api.Host, appctx.Config.Api.Domain)).Info()

	return nil

}

func ginRun(actx *appctx.AppCtx, engin *gin.Engine) (err error) {
	var (
		server *http.Server
		m      *autocert.Manager
		mu     sync.Mutex // 互斥锁，用于保护 server 变量
	)

	go func() {
		//安全关闭
		<-actx.Ctx.Done()
		mu.Lock()
		if server != nil {
			err := server.Shutdown(actx.Ctx)
			if err != nil {
				actx.ApiLog.Msg("api服务关闭失败").ErrData(err).Err()
			}
		}
		mu.Unlock()
	}()

	if len(actx.Config.Api.Domain) == 0 {
		mu.Lock()
		server = &http.Server{
			Addr:    actx.Config.Api.Host,
			Handler: engin,
		}
		mu.Unlock()
		return server.ListenAndServe()

	} else {
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,                                      // 接受 Let's Encrypt 的服务条款
			Cache:      autocert.DirCache(path.Join(actx.Config.Root, "cache")), // 存储证书的位置
			HostPolicy: autocert.HostWhitelist(actx.Config.Api.Domain),          // 允许的域名
		}
		mu.Lock()
		server = &http.Server{
			Handler: engin,
		}
		mu.Unlock()
		return server.Serve(m.Listener())
	}

}
