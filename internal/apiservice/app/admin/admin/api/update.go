package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
)

type AdminUpdate struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAdminUpdate(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminUpdate{ginx.New(c), c, appctx}
}

// Do
// @api 	| admin | 2 | 修改用户
// @path 	| /api/admin/:id
// @method 	| PUT
// @urlparam |n id |d 用户ID |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | adminmodel.AdminUpdateForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
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
	if r := db.Model(tmpPo).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	//uid, err :=  this.appctx.GetUid(this.ctx)
	//if err != nil {
	//	return err
	//}
	//
	//if uid == po.ID {
	//	return errors.New("禁止修改自己")
	//}

	//更新
	if r := db.Select(adminmodel.AdminUpdateField).Updates(po); r.Error != nil {
		return r.Error
	}
	//修改人员下线
	//if (tmpPo.Status == "1" && po.Status == "0") || tmpPo.Role != po.Role {
	//	this.appctx.Redis.Del(context.Background(), mymodel.GetAdminRedisLoginId(int(tmpPo.ID)))
	//}

	return nil
}
