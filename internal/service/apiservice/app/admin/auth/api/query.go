package api

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/auth/authmodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/ginx/query"
	"go-service/internal/service/apiservice/router"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthQuery struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAuthQuery doc
// @tags 3-系统-权限管理
// @summary 查询权限
// @Security		ApiKeyAuth
// @router /api/auth [get]
// @param key query string false "关键字，可以是name或者code，进行模糊匹配。"
// @success 200 {object} ginx.Response{data=[]authmodel.AuthVo}  "权限树列表，所有的权限数据"
func NewAuthQuery(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthQuery{ginx.New(c), c, appctx}
}

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
	qa, _ := this.ctx.GetQuery("key")

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
		if strings.Contains(v.Name, sub) || strings.Contains(v.Code, sub) || strings.Contains(v.Api, sub) {
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
