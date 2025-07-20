package lg

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Format struct {
	l    *zap.Logger
	msg  string
	data []zapcore.Field
}

func (this *Format) Msg(msg string) *Format {
	this.msg = msg
	return this
}

func (this *Format) FmtData(format string, args ...interface{}) *Format {
	this.data = append(this.data, zap.String(DATA, fmt.Sprintf(format, args...)))
	return this
}

func (this *Format) Data(data string) *Format {
	this.data = append(this.data, zap.String(DATA, data))
	return this
}

func (this *Format) DataEx(data string) *Format {
	this.data = append(this.data, zap.String(DATAEX, data))
	return this
}

//func (this *Format) Kind(data string) *Format {
//	this.data = append(this.data, zap.String(KIND, data))
//	return this
//}

func (this *Format) Trace(data string) *Format {
	this.data = append(this.data, zap.String(TRACEID, data))
	return this
}

func (this *Format) ErrData(err error) *Format {
	this.data = append(this.data, zap.Error(err))
	return this
}

func (this *Format) JSON(data interface{}) {
	d, _ := json.Marshal(data)
	this.data = append(this.data, zap.String(JsonFormat, string(d)))
}

func (this *Format) Info() {
	this.l.Info(this.msg, this.data...)
}

func (this *Format) Err() {
	this.l.Error(this.msg, this.data...)
}

func (this *Format) Debug() {
	this.l.Log(zapcore.DebugLevel, this.msg, this.data...)
}

func (this *Format) Panic() {
	this.l.Panic(this.msg, this.data...)
}
