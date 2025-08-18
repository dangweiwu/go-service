package api

import (
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/auth/authmodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthUpdate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAuthUpdate doc
// @tags 3-系统-权限管理
// @summary 修改权限
// @Security		ApiKeyAuth
// @router /api/auth/{id} [put]
// @param id path int true "权限ID"
// @param body body authmodel.AuthUpdateForm true "用户信息"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse ""
func NewAuthUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthUpdate{ginx.New(c), c, appctx}
}

func (this *AuthUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &authmodel.AuthUpdateForm{}
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

func (this *AuthUpdate) Update(po *authmodel.AuthUpdateForm) error {
	db := this.appctx.Db
	tmpPo := &authmodel.AuthPo{}
	tmpPo.ID = po.ID
	if r := db.Model(tmpPo).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	//更新
	if r := db.Save(po); r.Error != nil {
		return r.Error
	}
	return nil
}
