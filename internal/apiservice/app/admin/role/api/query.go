package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/ginx/query"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleQuery struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewRoleQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleQuery{ginx.New(c), c, appctx}
}

var likeRule = []string{"code", "name"}

// Do
// @api 	| role | 4 | 角色查询
// @path 	| /api/role
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query   |n limit   |d 条数 |e 10 |t int
// @query   |n current |d 页码 |e 1  |t int
// @query 	|n code |d 角色编码 |t string
// @query   |n name |d 角色名称 |t string
// @response | query.QueryResult | 200 Response
// @response | query.Page | Page定义
// @response | rolemodel.RolePo | []Data 定义
func (this *RoleQuery) Do() error {

	q := query.NewQuery(this.appctx.Db, this.Giner).SetTable(&rolemodel.RolePo{}).
		Select(rolemodel.ViewFields).WhereLike(likeRule)
	q.Db = q.Db.Order("order_num")
	r, err := q.PageFind(&[]rolemodel.RolePo{})
	if err != nil {
		return err
	}
	this.Rep(r)
	return nil
}
