package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"log"
	"time"
)

type RefreshToken struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

func NewRefreshToken(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &RefreshToken{ginx.New(c), c, appctx}
}

// Do
// @api     | me | 6 | 刷新token
// @path 	| /api/token/refresh
// @method 	| POST
// @header   |n Authorization |d token  |t string |c 鉴权
// @response | memodel.LogRep |200 Response
// @response | ginx.ErrResponse | 401 Response
// @tbtitle | Msg 数据
// @tbrow   |n msg |d refreshtoken已失效
func (this *RefreshToken) Do() error {

	//form := &mymodel.RefreshTokeForm{}
	//if err := this.Bind(form); err != nil {
	//	return err
	//}
	log.Println("refresh")
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
