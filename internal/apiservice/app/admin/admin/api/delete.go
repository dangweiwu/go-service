package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type AdminDel struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAdminDel doc
// @tags 2-系统-用户管理
// @summary 删除用户
// @Description 禁止删除自己。
// @Security		ApiKeyAuth
// @router /api/admin/{id} [delete]
// @param id path int true "用户ID"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=禁止删除自己"
func NewAdminDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminDel{ginx.New(c), c, appctx}
}

func (this *AdminDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}
	//
	if id == uid {
		return errors.New("禁止删除自己")
	}

	if err := this.Delete(id); err != nil {
		return err
	} else {
		this.RepOk()
		return nil
	}
}

func (this *AdminDel) Delete(id int64) error {
	db := this.appctx.Db
	po := &adminmodel.AdminPo{}
	po.ID = id
	if r := db.Select(adminmodel.AdminViewField).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	if r := db.Delete(po); r.Error != nil {
		return r.Error
	}
	this.appctx.Redis.Del(context.Background(), memodel.GetAdminRedisLoginId(int(po.ID)))

	return nil
}
