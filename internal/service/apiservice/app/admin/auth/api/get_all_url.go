package api

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/pkg"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

/*
获取全部url
*/
type AllUrl struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAllUrl doc
// @tags 3-系统-权限管理
// @summary 获取所有URL
// @description url参数用于auth的api
// @Security		ApiKeyAuth
// @router /api/allurl [get]
// @success 200 {object} ginx.Response{data=[]string}  "系统所有的可用URL"
func NewAllUrl(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AllUrl{ginx.New(c), c, appctx}
}

func (this *AllUrl) Do() error {
	this.Rep(ginx.Response{Data: pkg.AllUrl.GetUrl()})
	return nil
}
