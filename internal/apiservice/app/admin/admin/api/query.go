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

// NewAdminQuery doc
// @tags 2-系统-用户管理
// @summary 查询用户
// @security		ApiKeyAuth
// @router /api/admin [get]
// @param limit query int false "条数 默认10"
// @param current query int false "当前页码 默认1"
// @param account query string false "账号 模糊查询"
// @param phone query string false "手机号 模糊查询"
// @param email query string false "Email 模糊查询"
// @param name query string false "姓名 模糊查询"
// @success 200 {object} query.QueryPageResult{data=[]adminmodel.AdminVo}  "数据"
func NewAdminQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminQuery{ginx.New(c), c, appctx}
}

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
