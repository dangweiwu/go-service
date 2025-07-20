package middler

import (
	"fmt"
	"go-service/internal/bootstrap/appctx"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Recovery(appctx *appctx.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				raw := c.Request.URL.RawQuery
				_path := c.Request.URL.Path

				if raw != "" {
					_path = _path + "?" + raw
				}
				if brokenPipe {
					appctx.Log.Msg("网络中断").ErrData(err.(error)).Data("path:" + _path).DataEx("requestid:" + requestid.Get(c)).Err()
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
					return
				}
				appctx.Log.Msg("系统异常").ErrData(err.(error)).Data(string(debug.Stack())).DataEx("requestid:" + requestid.Get(c)).Err()
				c.AbortWithStatus(http.StatusInternalServerError)
				c.String(500, fmt.Sprintf("%v", err))
			}
		}()
		c.Next()
	}
}
