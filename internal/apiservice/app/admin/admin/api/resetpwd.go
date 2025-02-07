package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
	"math/rand"
)

type ResetPassword struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewResetPassword doc
// @tags 2-系统-用户管理
// @summary 修改密码
// @Description 生成数字与字母组合的随机6位密码，需要自己保存。生成密码后对应账号下线。不能重置自己密码
// @Security		ApiKeyAuth
// @router /api/admin/resetpwd/{id} [put]
// @Param id path int true "用户ID"
// @success 200  {object} ginx.Response{data=string} "data=新密码"
// @failure 400 {object} ginx.ErrResponse "msg=不能重置自己密码"
func NewResetPassword(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &ResetPassword{ginx.New(c), c, appctx}
}

func (this *ResetPassword) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &adminmodel.AdminPo{}
	po.ID = id
	pwd, err := this.ResetPassword(po)
	if err != nil {
		return err
	}
	this.Rep(ginx.Response{Data: pwd})
	return nil
}

func (this *ResetPassword) ResetPassword(rawPo *adminmodel.AdminPo) (string, error) {
	id, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return "", err
	}
	if id == rawPo.ID {
		return "", errors.New("不能重置自己密码")
	}

	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Select([]string{"id", "password"}).Where("id=?", rawPo.ID).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return "", errors.New("记录不存在")
		} else {
			return "", r.Error
		}
	}

	_pwd := this.newPwd()
	pwd, err := pkg.GetPassword(_pwd)
	if err != nil {
		return "", fmt.Errorf("生成密码异常:%w", err)
	}

	r := this.appctx.Db.Model(po).Update("password", pwd)

	//踢出在线
	this.appctx.Redis.Del(context.Background(), memodel.GetAdminRedisLoginId(int(po.ID)))

	return _pwd, r.Error
}

// 生成三位字符，三位数字的密码
var ltes = "abcdefghijkmnpqrstuvwxyz"
var nums = "0123456789"

func (this *ResetPassword) newPwd() string {

	rt := ""
	for i := 0; i < 3; i++ {
		rt += string(ltes[rand.Intn(len(ltes))])
	}
	for i := 0; i < 3; i++ {
		rt += string(nums[rand.Intn(len(nums))])
	}
	return rt
}
