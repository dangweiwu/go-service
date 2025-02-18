package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/auth/authmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type AuthDel struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAuthDel doc
// @tags 3-系统-权限管理
// @summary 删除权限
// @security		ApiKeyAuth
// @description 该权限下包含其他权限，禁止删除！
// @router /api/auth/{id} [delete]
// @param id path int true "角色ID"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=该权限下包含其他权限，禁止删除！"
func NewAuthDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AuthDel{ginx.New(c), c, appctx}
}

func (this *AuthDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	if err := this.Delete(id); err != nil {
		return err
	} else {
		this.RepOk()
		return nil
	}
}

func (this *AuthDel) Delete(id int64) error {
	db := this.appctx.Db
	po := &authmodel.AuthPo{}
	po.ID = id
	if r := db.Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	ct := int64(0)
	if r := db.Model(&authmodel.AuthPo{}).Where("parent_id=?", id).Count(&ct); r.Error != nil {
		return r.Error
	}

	if ct != 0 {
		return errors.New("该权限下包含其他权限，禁止删除！")
	}

	if r := db.Delete(po); r.Error != nil {
		return r.Error
	}
	return nil
}
