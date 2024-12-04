package middler

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go-service/internal/bootstrap/appctx"
)

func RegMiddler(actx *appctx.AppCtx) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		requestid.New(),
		Cors(),
		Recovery(actx),
		ApiLog(actx),
	}
}
