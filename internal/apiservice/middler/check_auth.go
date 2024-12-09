package middler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/apiservice/pkg/ginx"
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/bootstrap/appctx"
)

func NoAuthErrResponse(c *gin.Context, err error) {
	c.JSON(403, ginx.ErrResponse{Kind: ginx.MSG, Data: fmt.Errorf("Forbidden:%w", err).Error(), Msg: "缺少权限"})
	c.Abort()
}

// 权限中间件
func CheckAuth(appctx *appctx.AppCtx) gin.HandlerFunc {
	return func(context *gin.Context) {

		yes, err := jwtx.GetIsSuper(context)
		if err != nil {
			NoAuthErrResponse(context, err)
			return
		}
		if yes {
			context.Next()
			return
		}
		role, err := jwtx.GetRole(context)
		if err != nil {
			NoAuthErrResponse(context, err)
			return
		}

		if ok, err := appctx.Casbin.Enforce(role, context.FullPath(), context.Request.Method); ok {
			context.Next()
			return
		} else if err != nil {
			NoAuthErrResponse(context, err)
		} else {
			NoAuthErrResponse(context, fmt.Errorf("%s:%s", context.Request.Method, context.FullPath()))
		}
	}
}
