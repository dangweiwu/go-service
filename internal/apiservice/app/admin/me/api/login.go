package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/me/memodel"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/apiservice/router"
	"go-service/internal/bootstrap/appctx"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Login struct {
	*ginx.Ginx
	ctx    *gin.Context
	appctx *appctx.AppCtx
}

// NewLogin 登录系统
// @tags 1-系统-我的
// @summary 登录系统
// @Description	用户登录
// @router /api/login [post]
// @param body body memodel.LoginForm true "登录"
// @success 200 {object} memodel.LogRep
// @failure 400 {object} ginx.ErrResponse "msg=密码错误|账号不存在|账号被禁用"
func NewLogin(appctx *appctx.AppCtx, c *gin.Context) router.Handler {
	return &Login{ginx.New(c), c, appctx}
}

func (this *Login) Do() error {

	//数据源
	po := &memodel.LoginForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	data, err := this.Login(po)
	if err != nil {
		return err
	}
	this.Rep(data)
	return nil
}
func (this *Login) Login(form *memodel.LoginForm) (*memodel.LogRep, error) {
	var (
		accessToken  string
		refreshToken string
		logcode      string
		err          error
	)
	po, err := this.Valid(form)
	if err != nil {
		return nil, err
	}

	pwd, err := pkg.GetPassword(form.Password)
	if err != nil {
		return nil, err
	}
	if pwd != po.Password {
		return nil, errors.New("密码错误")
	}

	logcode, err = newLogCode(po.ID, this.appctx.Redis)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	exp := now + this.appctx.Config.Jwt.Exp

	accessToken, err = jwtx.Token{
		SecretKey: this.appctx.Config.Jwt.Secret,
		Exp:       exp,
		UserId:    po.ID,
		IsSuper:   po.IsSuperAdmin,
		LoginCode: logcode,
		Kind:      jwtx.ACCESS,
		Role:      po.Role,
	}.Gen()
	if err != nil {
		return nil, err
	}

	refreshToken, err = jwtx.Token{
		SecretKey: this.appctx.Config.Jwt.Secret,
		Exp:       now + this.appctx.Config.Jwt.Exp*3,
		UserId:    po.ID,
		IsSuper:   po.IsSuperAdmin,
		LoginCode: logcode,
		Kind:      jwtx.REFRESH,
		Role:      po.Role,
	}.Gen()
	if err != nil {
		return nil, err
	}

	return &memodel.LogRep{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenExp:     exp - 10,
	}, nil

}

func (this *Login) Valid(form *memodel.LoginForm) (*adminmodel.AdminPo, error) {
	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Select([]string{"id", "is_super_admin", "role", "password", "status"}).Where("account=?", form.Account).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return po, errors.New("账号不存在")
		} else {
			return po, r.Error
		}
	}
	pwd, err := pkg.GetPassword(form.Password)
	if err != nil {
		return nil, err
	}
	if pwd != po.Password {
		return nil, errors.New("密码错误")
	}

	if po.Status == "0" {
		return nil, errors.New("账号被禁用")
	}

	return po, nil
}
func newLogCode(userid int64, rd *redis.Client) (logincode string, err error) {
	//登陆处理
	//登陆code 控制唯一登陆有效及踢人
	if logincode = uuid.New().String(); logincode == "" {
		err = errors.New("logincode is empty")
		return
	} else {
		logincode = strings.Split(logincode, "-")[0]
		if r := rd.Set(context.Background(), memodel.GetAdminRedisLoginId(int(userid)), logincode, 0); r.Err() != nil {
			err = errors.New("login code redis:" + r.Err().Error())
			return
		}
	}
	return logincode, nil

}
