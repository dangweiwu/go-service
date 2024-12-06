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

type AdminDel struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewAdminDel(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &AdminDel{ginx.New(c), c, appctx}
}

// Do
// @api | admin | 5 | 删除用户
// @path | /api/admin/:id
// @method | DELETE
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @urlparam |n id |d 用户ID |v required |t int    |e 1
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t type
// @response | hd.ErrResponse | 500 RESPONSE
// @tbtitle  | msg 数据
// @tbrow    |n msg |e 禁止删除自己
// @tbrow    |n msg |e 记录不存在
func (this *AdminDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	//uid, err := this.appctx.GetUid(this.ctx)
	//if err != nil {
	//	return err
	//}
	//
	//if id == uid {
	//	return errors.New("禁止删除自己")
	//}

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
	if r := db.Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	if r := db.Delete(po); r.Error != nil {
		return r.Error
	}
	return nil
}
