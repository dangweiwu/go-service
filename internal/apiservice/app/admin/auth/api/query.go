package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/auth/authmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/ginx/query"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
	"strings"
)

type AuthQuery struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAuthQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthQuery{ginx.New(c), c, appctx}
}

// Do
// @api 	| auth | 3 | 权限查询
// @path 	| /api/auth
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query 	|n name |d 名称 |e 0 |t string |c 权限名称
// @response | ginx.Response | 200 Response
// @response | authmodel.AuthVo | Data定义
func (this *AuthQuery) Do() error {

	data, err := this.Query()
	if err != nil {
		return err
	} else {
		this.Rep(data)
		return nil
	}
}

func (this *AuthQuery) Query() (interface{}, error) {
	po := &authmodel.AuthVo{}
	pos := []authmodel.AuthVo{}

	q := query.NewQuery(this.appctx.Db, this.Ginx).SetTable(po)
	qa, _ := this.ctx.GetQuery("name")

	q.Db = q.Db.Where("parent_id=0").Preload("Children", func(db *gorm.DB) *gorm.DB {
		db = db.Order("order_num")
		return db
	}).Order("order_num").Find(&pos)
	if len(qa) != 0 {
		pos = voloop(pos, qa)
	}

	return ginx.Response{Data: pos}, nil
}

func voloop(a []authmodel.AuthVo, sub string) []authmodel.AuthVo {
	var result []authmodel.AuthVo
	for _, v := range a {
		if strings.Contains(v.Name, sub) {
			v.Children = voloop(v.Children, sub)
			result = append(result, v)
		} else {
			// 否则，递归过滤子节点
			filteredChildren := voloop(v.Children, sub)
			if len(filteredChildren) > 0 {
				// 如果子节点中有符合条件的节点，则保留这些子节点
				result = append(result, authmodel.AuthVo{
					ID:       v.ID,
					Name:     v.Name,
					Code:     v.Code,
					OrderNum: v.OrderNum,
					Api:      v.Api,
					Method:   v.Method,
					Kind:     v.Kind,
					ParentId: v.ParentId,
					Children: filteredChildren,
				})
			}
		}

	}
	return result

}
