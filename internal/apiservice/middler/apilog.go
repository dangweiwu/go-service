package middler

import (
	"go-service/internal/apiservice/pkg/jwtx"
	"go-service/internal/bootstrap/appctx"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func ApiLog(appctx *appctx.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		_path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		userid, _ := jwtx.GetUserid(c)
		TimeStamp := time.Now()
		Latency := TimeStamp.Sub(start)
		if raw != "" {
			_path = _path + "?" + raw
		}

		// GET请求打印到控制台，不打印到日志中
		// if c.Request.Method != "GET" {
		// 	appctx.ApiLog.Msg("request").Data(fmt.Sprintf("userid:%d method:%s path:%s status:%d size:%d latency:%d",
		// 		userid, c.Request.Method, _path, c.Writer.Status(), c.Writer.Size(), int(Latency.Milliseconds()),
		// 	)).DataEx("requestid:" + requestid.Get(c)).Info()
		// 	if len(c.Errors) != 0 {
		// 		appctx.ApiLog.Msg("request").ErrData(c.Errors[0]).Err()
		// 	}
		// } else {
		// 	log.Printf("userid:%d method:%s path:%s status:%d size:%d latency:%d requestid: %s",
		// 		userid, c.Request.Method, _path, c.Writer.Status(), c.Writer.Size(), int(Latency.Milliseconds()), requestid.Get(c))
		// 	if len(c.Errors) != 0 {
		// 		log.Println(c.Errors)
		// 	}

		// }
		appctx.ApiLog.ApiLog("").Path(c.Request.Method + ":" + _path).
			Status(c.Writer.Status()).
			Size(c.Writer.Size()).
			Latency(Latency.Milliseconds()).
			UserId(userid).
			Trace(requestid.Get(c)).
			Info()

		if len(c.Errors) != 0 {
			appctx.Log.Msg("request").ErrData(c.Errors[0]).Err()
		}
	}
}
