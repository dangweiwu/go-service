package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleUpdate struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewRoleUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleUpdate{ginx.New(c), c, appctx}
}

// Do
// @api 	| role | 6 | 修改角色
// @path 	| /api/role/:id
// @method 	| PUT
// @urlparam |n id |d 角色ID   |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | rolemodel.RoleUpdate
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *RoleUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &rolemodel.RoleUpdate{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	if err := roleservice.NewRoleService(this.appctx).UpdateRole(po); err != nil {
		return err
	}

	this.RepOk()
	return nil
}
