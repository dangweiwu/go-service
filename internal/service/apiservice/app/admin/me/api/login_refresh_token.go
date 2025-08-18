package api

import (
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/service/apiservice/app/admin/me/memodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/jwtx"
	"go-service/internal/service/apiservice/router"
	"time"

	"github.com/gin-gonic/gin"
)

type RefreshToken struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewRefreshToken 刷新token
// @tags 1-系统-我的
// @summary TOKEN刷新
// @description	通过reflesh token获取access token,
// @description	reflesh token有效期是access token的3倍时长,
// @description	access token到期前进行续期。
// @router /api/token/refresh [post]
// @param Authorization header string true "reflesh_token"
// @success 200 {object} memodel.LogRep "与登录获取数据一致"
// @failure 400 {object} ginx.ErrResponse "msg=账号已禁用"
func NewRefreshToken(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RefreshToken{ginx.New(c), c, appctx}
}

func (this *RefreshToken) Do() error {

	uid, err := jwtx.GetUserid(this.ctx)
	if err != nil {
		return err
	}

	po := &adminmodel.AdminPo{}
	if err := this.appctx.Db.Select(adminmodel.AdminViewField).Where("id=?", uid).First(po).Error; err != nil {
		return err
	}

	if po.Status != "1" {
		return errors.New("账号已禁用")
	}
	logcode, err := jwtx.GetLoginCode(this.ctx)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	exp := now + this.appctx.Config.Jwt.Exp

	accessToken, err := jwtx.Token{
		SecretKey: this.appctx.Config.Jwt.Secret,
		Exp:       exp,
		UserId:    po.ID,
		IsSuper:   po.IsSuperAdmin,
		LoginCode: logcode,
		Kind:      jwtx.ACCESS,
		Role:      po.Role,
	}.Gen()
	if err != nil {
		return err
	}

	refreshToken, err := jwtx.Token{
		SecretKey: this.appctx.Config.Jwt.Secret,
		Exp:       now + this.appctx.Config.Jwt.Exp*3,
		UserId:    po.ID,
		IsSuper:   po.IsSuperAdmin,
		LoginCode: logcode,
		Kind:      jwtx.REFRESH,
		Role:      po.Role,
	}.Gen()
	if err != nil {
		return err
	}

	this.Rep(&memodel.LogRep{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenExp:     exp - 10,
	})
	return nil
}
