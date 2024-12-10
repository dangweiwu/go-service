package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/role/roleservice"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type RoleDel struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewRoleDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RoleDel{ginx.New(c), c, appctx}
}

// Do
// @api 	| admin | 2 | 删除角色
// @path 	| /api/role/:id
// @method 	| DELETE
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t type
// @response | ginx.ErrResponse | 500 RESPONSE
// @tbtitle  | msg 数据
// @tbrow    |n msg |e 记录不存在
func (this *RoleDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	if err := roleservice.NewRoleService(this.appctx).Delete(id); err != nil {
		return err
	}

	this.RepOk()
	return nil

}
