package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type SetAuth struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewSetAuth(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &SetAuth{ginx.New(c), c, appctx}
}

// Do
// @api 	| role | 5 | 设定角色权限
// @path 	| /api/role/auth/:id
// @method 	| PUT
// @urlparam |n id |d 角色ID   |v required |t int    |e 1
// @headers   |n Authorization |d token  |t string |c 鉴权
// @form     | rolemodel.RoleAuthForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *SetAuth) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	po := &rolemodel.RoleAuthForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	s := roleservice.NewRoleService(this.appctx)
	err = s.SetAuth(id, po.Auth)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}
