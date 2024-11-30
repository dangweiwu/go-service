package ginx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Ginx struct {
	Gctx *gin.Context
}

func NewHd(ctx *gin.Context) *Ginx {
	return &Ginx{ctx}
}

func (this *Ginx) GetId() (int64, error) {
	_id, has := this.Gctx.Params.Get("id")
	if !has {
		return 0, errors.New("缺少ID")
	}
	id, err := strconv.Atoi(_id)
	if err != nil {
		return 0, errors.New("无效ID")
	}
	return int64(id), nil
}

func (this *Ginx) GetUrlkey(name string) (string, error) {
	key, has := this.Gctx.Params.Get(name)
	if !has {
		return "", fmt.Errorf("缺少%s", name)
	}
	return key, nil
}

func (this *Ginx) Bind(po interface{}) error {
	return this.Gctx.ShouldBindJSON(po)
}

func (this *Ginx) Rep(data interface{}) {
	this.Gctx.JSON(200, data)
}

func (this *Ginx) RepOk() {
	this.Gctx.JSON(200, Response{"ok"})
}

func (this *Ginx) ErrCode(msg string, data string) ErrResponse {
	return ErrResponse{CODE, data, msg}
}

func (this *Ginx) ErrMsg(msg string, data string) ErrResponse {
	return ErrResponse{MSG, data, msg}
}
