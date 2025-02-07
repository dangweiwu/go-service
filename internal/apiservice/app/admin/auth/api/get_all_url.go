package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

/*
获取全部url
*/
type AllUrl struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAllUrl(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AllUrl{ginx.New(c), c, appctx}
}

// Do
// @api 	| auth | 5 | 获取所有URL | 创建修改权限时url参数从这获取
// @path 	| /api/allurl
// @method 	| GET
// @headers 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @tbtitle | 200 Response
// @tbrow |n data |d 权限列表 |t []string |c 列表数据 |e ['/api/admin']
func (this *AllUrl) Do() error {
	this.Rep(ginx.Response{Data: pkg.AllUrl.GetUrl()})
	return nil
}
