package middler

import (
	"fmt"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/pkg/ginx"
	"go-service/internal/service/apiservice/pkg/jwtx"

	"github.com/gin-gonic/gin"
)

func ErrKind(c *gin.Context, err error) {
	c.JSON(400, ginx.ErrResponse{Kind: ginx.MSG, Data: err.Error(), Msg: "jwt类型错误"})
	c.Abort()
}

func CheckTokenKind(appctx *appctx.AppCtx, kind int) gin.HandlerFunc {
	return func(c *gin.Context) {
		now_kind, err := jwtx.GetKind(c)
		if err != nil {
			ErrKind(c, err)
			return
		}
		if kind != now_kind {
			ErrKind(c, fmt.Errorf("kind err :%d", now_kind))
			return
		}
		c.Next()
	}
}
