package middler

/*校验code是否有效 无效则退出登陆
1. 放在token中间件之后
2. 必须有redis
*/
import (
	"context"
	"errors"
	"fmt"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/me/memodel"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/jwtx"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LoginCodeErrResponse(c *gin.Context, err error) {
	c.JSON(401, ginx.ErrResponse{Kind: ginx.MSG, Data: err.Error(), Msg: "重新登录"})
	c.Abort()
}

func CheckLoginCode(appctx *appctx.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		code, err := jwtx.GetLoginCode(c)
		if err != nil {
			LoginCodeErrResponse(c, fmt.Errorf("code:%w", err))
			return
		}
		uid, err := jwtx.GetUserid(c)
		if err != nil {
			LoginCodeErrResponse(c, fmt.Errorf("jwt_get_id:%w", err))
			return
		}
		logincode, err := appctx.Redis.Get(context.Background(), memodel.GetAdminRedisLoginId(int(uid))).Result()
		if err != nil {
			if err == redis.Nil {
				LoginCodeErrResponse(c, fmt.Errorf("redis_code:%w", err))
			} else {
				LoginCodeErrResponse(c, err)
			}
			return
		}

		if logincode != code {
			LoginCodeErrResponse(c, errors.New("invalid_login_code"))
			return
		}
		c.Next()
	}
}
