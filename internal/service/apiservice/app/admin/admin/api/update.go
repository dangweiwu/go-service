package api

import (
	"context"
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/service/apiservice/app/admin/me/memodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/jwtx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminUpdate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAdminUpdate doc
// @tags 2-系统-用户管理
// @summary 修改用户
// @description status修改为0会导致对应账号下线，禁止修改自己。
// @security		ApiKeyAuth
// @router /api/admin/{id} [put]
// @param id path int true "用户ID"
// @param body body adminmodel.AdminUpdateForm true "用户信息"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=禁止修改自己"
func NewAdminUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminUpdate{ginx.New(c), c, appctx}
}

func (this *AdminUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &adminmodel.AdminUpdateForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	err = this.Update(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}
func (this *AdminUpdate) Update(po *adminmodel.AdminUpdateForm) error {
	db := this.appctx.Db
	tmpPo := &adminmodel.AdminUpdateForm{}
	tmpPo.ID = po.ID
	if r := db.Model(tmpPo).Select(adminmodel.AdminUpdateField).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}
	//
	if uid == po.ID {
		return errors.New("禁止修改自己")
	}

	//更新
	if r := db.Select(adminmodel.AdminUpdateField).Updates(po); r.Error != nil {
		return r.Error
	}
	//被修改人下线
	if (tmpPo.Status != po.Status) || tmpPo.Role != po.Role {
		this.appctx.Redis.Del(context.Background(), memodel.GetAdminRedisLoginId(int(tmpPo.ID)))
	}

	return nil
}
