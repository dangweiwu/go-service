package middler

import (
	"errors"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/pkg/ginx"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

func RefuseResponse(c *gin.Context) {
	c.JSON(401, ginx.ErrResponse{Kind: ginx.MSG, Data: "invalidauth", Msg: "鉴权失效"})
	c.Abort()
}

func ErrToken(c *gin.Context, err error) {
	c.JSON(400, ginx.ErrResponse{Kind: ginx.MSG, Data: err.Error(), Msg: "无效token"})
	c.Abort()
}

//token 中间件

func TokenParse(appctx *appctx.AppCtx) gin.HandlerFunc {
	sec := appctx.Config.Jwt.Secret
	return func(c *gin.Context) {
		tk, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (interface{}, error) {
			return []byte(sec), nil
		})

		if errors.Is(err, jwt.ErrTokenExpired) {
			RefuseResponse(c)
			return
		}
		if err != nil {
			ErrToken(c, err)
			return
		}

		claims, ok := tk.Claims.(jwt.MapClaims)
		if !ok {
			RefuseResponse(c)
			return
		}
		for k, v := range claims {
			switch k {
			case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
				// ignore the standard claims
			default:
				c.Set(k, v)
			}
		}
		c.Next()
	}
}
