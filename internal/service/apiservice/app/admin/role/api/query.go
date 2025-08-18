package api

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/role/rolemodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/ginx/query"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
)

type RoleQuery struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewRoleQuery doc
// @tags 4-系统-角色管理
// @summary 角色查询
// @Security		ApiKeyAuth
// @router /api/role [get]
// @param limit query int false "条数 默认10"
// @param current query int false "当前页码 默认1"
// @param code query string false "编码 模糊查询"
// @param name query string false "名称 模糊查询"
// @success 200 {object} query.QueryPageResult{data=[]rolemodel.RolePo}  "数据"
func NewRoleQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleQuery{ginx.New(c), c, appctx}
}

var likeRule = []string{"code", "name"}

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
