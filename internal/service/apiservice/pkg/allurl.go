package pkg

import (
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

/*
树结构展示部分
不需要进行鉴权的地址
与auth保持一致
*/

var skipurl = map[string]struct{}{
	"/api/allurl":        {},
	"/api/login":         {},
	"/api/my":            {},
	"/api/logout":        {},
	"/api/auth/tree":     {},
	"/api/my/password":   {},
	"/api/token/refresh": {},
}

var once sync.Once
var AllUrl *allUrl

type allUrl struct {
	sync.RWMutex
	Url []string
	e   *gin.Engine
}

// 单例模式
func NewAllUrl(e *gin.Engine) *allUrl {
	if AllUrl == nil {
		once.Do(func() {
			if AllUrl == nil {
				AllUrl = &allUrl{e: e}
			}
		})
	}
	return AllUrl
}

func (this *allUrl) GetUrl() []string {
	this.RLock()
	defer this.RUnlock()
	return this.Url
}

func (this *allUrl) InitUrl() {
	this.Lock()
	defer this.Unlock()
	a := this.e.Routes()

	uniqueRole := map[string]struct{}{}
	this.Url = []string{}
	for _, v := range a {
		if strings.Index(v.Path, "/api") == 0 {
			if _, has := uniqueRole[v.Path]; !has {
				if _, has := skipurl[v.Path]; !has {
					this.Url = append(this.Url, v.Path)
				}
				uniqueRole[v.Path] = struct{}{}
			}
		}
	}
}
