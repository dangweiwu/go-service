package doc

// rapidoc文档 https://rapidocweb.com/api.html#events
// swag文档 https://github.com/swaggo/swag/blob/master/README_zh-CN.md
import (
	"github.com/gin-gonic/gin"
	"go-service/internal/bootstrap/appctx"
	"html/template"
	"math/rand"
	"strconv"
	"time"
)

var tpl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <script
            type="module"
            src="/view/doc/rapidoc-min.js"
    ></script>
</head>
<body>
<rapi-doc
        spec-url = "/doc/swag.yaml?a={{.Notice}}"
        style = "height:100vh; width:100%"
        font-size="largest"
        show-method-in-nav-bar="as-colored-text"
        render-style="focused"
        show-header="false"
        persist-auth="true"
        sort-tags="true"
>
</rapi-doc>
</body>
</html>
`

// 定义一个结构体来传递给模板
type TemplateData struct {
	Notice string
}

func InitDoc(appctx *appctx.AppCtx, engine *gin.Engine) {
	engine.GET("/doc/swag.yaml", func(c *gin.Context) {
		c.File(appctx.Config.Api.DocDir + "/swagger.yaml")
	})
	engine.GET("/doc", gin.BasicAuth(gin.Accounts{
		appctx.Config.Api.DocUser: appctx.Config.Api.DocPassword,
	}), func(c *gin.Context) {
		t, err := template.New("doc").Parse(tpl)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		// 生成一个随机数作为notice参数
		rand.Seed(time.Now().UnixNano())
		notice := rand.Intn(100000)

		// 创建模板数据
		data := TemplateData{
			Notice: strconv.Itoa(notice),
		}

		err = t.Execute(c.Writer, data)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
	})
}
