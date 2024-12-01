package middler

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/bootstrap/appctx"
)

func RegMiddler(actx *appctx.AppCtx) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		Cors(),
	}
}
