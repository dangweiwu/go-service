package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/ginx/query"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type AdminQuery struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAdminQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminQuery{ginx.New(c), c, appctx}
}

// Do
// @api 	| admin | 4 | 查询用户
// @path 	| /api/admin
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query   |n limit   |d 条数 |e 10 |t int
// @query   |n current |d 页码 |e 1  |t int
// @query 	|n account |d 账号 |e admin | t string
// @query   |n phone   |d 手机号 |e 12345678911 |t int
// @query   |n email   |d email
// @query   |n name    |d 姓名
// @response | query.QueryResult | 200 Response
// @response | query.Page | Page定义
// @response | adminmodel.AdminVo | []Data 定义
func (this *AdminQuery) Do() error {
	data, err := this.Query()
	if err != nil {
		return err
	} else {
		this.Rep(data)
		return nil
	}
}

var likeRule = []string{"account", "phone", "email", "name"}
var Rule = []string{"status"}

func (this *AdminQuery) Query() (interface{}, error) {

	po := &adminmodel.AdminVo{}
	pos := []adminmodel.AdminVo{}
	q := query.NewQuery(this.appctx.Db, this.Giner).
		SetTable(po).
		Select(adminmodel.AdminViewField).
		WhereLike(likeRule).
		Where(Rule).
		Order()
	return q.PageFind(&pos)
}
