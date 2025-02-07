package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
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

// NewMeInfo
// @tags 1-系统-我的
// @summary 我的详情
// @Description	获取我的详情
// @router /api/me [get]
// @Security		ApiKeyAuth
// @success 200 {object} memodel.MeInfo "个人信息详情"
func NewMeInfo(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &MeInfo{ginx.New(c), c, appctx}
}

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
	role := &rolemodel.RolePo{}
	if po.IsSuperAdmin != "1" {
		this.appctx.Db.Model(role).Select("code").Where("id = ?", po.Role).Take(role)
	}

	vo := &memodel.MeInfoVo{}
	vo.Name = po.Name
	vo.Account = po.Account
	vo.Role = po.Role
	vo.RoleName = role.Name
	vo.Phone = po.Phone
	vo.IsSuperAdmin = po.IsSuperAdmin
	vo.Email = po.Email

	this.Rep(vo)
	return nil
}
