package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
)

type AdminCreate struct {
	ginx.Giner
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewAdminCreate doc
// @tags 2-系统-用户管理
// @summary 创建用户
// @router /api/admin [post]
// @Security		ApiKeyAuth
// @param body body adminmodel.AdminForm true "用户信息"
// @success 200 {object} ginx.Response{data=string} "data=ok"
// @failure 400 {object} ginx.ErrResponse "msg=账号已存在"
func NewAdminCreate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminCreate{ginx.New(c), c, appctx}
}

func (this *AdminCreate) Do() error {

	//数据源
	po := &adminmodel.AdminForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	err = this.Create(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *AdminCreate) Create(po *adminmodel.AdminForm) error {
	db := this.appctx.Db
	//验证是否已创建 或者其他检查
	if err := this.Valid(po); err != nil {
		return err
	}

	pwd, err := pkg.GetPassword(po.Password)
	if err != nil {
		return fmt.Errorf("创建账号失败:%w", err)
	}
	po.Password = pwd
	if po.IsSuperAdmin == "1" {
		po.Role = ""
	}
	if r := db.Create(po); r.Error != nil {
		return r.Error
	}
	return nil
}

func (this *AdminCreate) Valid(po *adminmodel.AdminForm) error {
	db := this.appctx.Db
	var ct = int64(0)
	if r := db.Model(po).Where("account = ?", po.Account).Count(&ct); r.Error != nil {
		return fmt.Errorf("校验失败:%w", r.Error)
	} else if ct != 0 {
		return errors.New("账号已存在")
	}

	/*
		if po.Phone != "" {
			if r := db.Model(po).Where("phone = ?", po.Phone).Count(&ct); r.Error != nil {
				return errs.WithMessage(r.Error, "校验失败")
			} else if ct != 0 {
				return errors.New("手机号已存在")
			}
		}

		if po.Email != "" {
			if r := db.Model(po).Where("email = ?", po.Email).Count(&ct); r.Error != nil {
				return errs.WithMessage(r.Error, "校验失败")
			} else if ct != 0 {
				return errors.New("Email已存在")
			}
		}
	*/
	return nil
}
