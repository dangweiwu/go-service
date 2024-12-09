package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type MeInfo struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewMeInfo(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &MeInfo{ginx.New(c), c, appctx}
}

// Do
// @api | me | 3 | 我的详情
// @path | /api/me
// @method | GET
// @header |n Authorization |d token |t string |c 鉴权
// @response | memodel.MeInfo | 200 Response
func (this *MeInfo) Do() error {

	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}

	po := &memodel.MeInfo{}
	if r := this.appctx.Db.Model(po).Select(memodel.MeViewField).Where("id = ?", uid).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	this.Rep(po)
	return nil
}
