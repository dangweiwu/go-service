package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleCreate struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewRoleCreate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleCreate{ginx.New(c), c, appctx}
}

// Do
// @api | role | 1 | 创建角色
// @path    | /api/role
// @method  | POST
// @header  |n Authorization |d token |t string |c 鉴权
// @form    | rolemodel.RoleForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *RoleCreate) Do() error {
	//数据源
	po := &rolemodel.RoleForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	if err := roleservice.NewRoleService(this.appctx).Save(po); err != nil {
		return err
	}
	this.RepOk()
	return nil
}
