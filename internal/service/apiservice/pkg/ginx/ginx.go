package ginx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var _ Giner = (*Ginx)(nil)

type Giner interface {
	GetId() (int64, error)
	GetUrlkey(name string) (string, error)
	Bind(po interface{}) error
	ShouldBindQuery(obj any) error
	Query(key string) (value string)
	Rep(data interface{})
	RepOk()
	ErrCode(msg string, data string) ErrResponse
	ErrMsg(msg string, data string) ErrResponse
}

type EmptyGinx struct {
}

func (e EmptyGinx) GetId() (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) GetUrlkey(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) Bind(po interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) ShouldBindQuery(obj any) error {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) Query(key string) (value string) {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) Rep(data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) RepOk() {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) ErrCode(msg string, data string) ErrResponse {
	//TODO implement me
	panic("implement me")
}

func (e EmptyGinx) ErrMsg(msg string, data string) ErrResponse {
	//TODO implement me
	panic("implement me")
}

type Ginx struct {
	Gctx *gin.Context
}

func New(ctx *gin.Context) *Ginx {
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

func (this *Ginx) ShouldBindQuery(obj any) error {
	return this.Gctx.ShouldBindQuery(obj)
}

func (this *Ginx) Query(key string) (value string) {
	return this.Gctx.Query(key)
}
