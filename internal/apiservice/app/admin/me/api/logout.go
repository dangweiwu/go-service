package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type LogOut struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewLogOut(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &LogOut{ginx.New(c), c, appctx}
}

// Do
// @api     | me | 2 | 登出
// @path 	| /api/logout
// @method 	| POST
// @header  |n Authorization |d token |t string |c 鉴权
// @tbtitle | 200 Response
// @tbrow   |n data |e ok |c 成功 |t string
func (this *LogOut) Do() error {
	err := this.Logout()
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *LogOut) Logout() error {
	//获取id
	id, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return nil
	}
	//删除logincode
	this.appctx.Redis.Del(context.Background(), memodel.GetAdminRedisLoginId(int(id)))
	return nil
}
