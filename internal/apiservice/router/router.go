package router

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/bootstrap/appctx"
)

type (
	Handler interface {
		Do() error
	}
	HandlerFunc func(ctx *appctx.AppCtx, c *gin.Context) Handler
)

func Do(actx *appctx.AppCtx, f HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := f(actx, c).Do(); err != nil {
			switch err.(type) {
			case ginx.ErrResponse:
				c.JSON(400, err)
			default:
				c.JSON(400, &ginx.ErrResponse{Kind: ginx.MSG, Msg: err.Error()})
			}
		}
	}
}

type Router struct {
	Root *gin.RouterGroup
	Jwt  *gin.RouterGroup //jwt登陆
	Auth *gin.RouterGroup //权限
}

func NewRouter(actx *appctx.AppCtx, g *gin.Engine) *Router {
	return &Router{
		Root: g.Group("/api"),
		//Jwt:  g.Group("/api", middler.TokenParse(actx), middler.CheckLoginCode(actx)),
		//Auth: g.Group("/api", middler.TokenParse(actx), middler.CheckLoginCode(actx), middler.CheckAuth(actx)),
		Auth: g.Group("/api"),
	}
}
