package lg

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HttpLog struct {
	l    *zap.Logger
	msg  string
	data []zapcore.Field
}

func (this *HttpLog) Status(data int) *HttpLog {
	this.data = append(this.data, zap.Int("status", data))
	return this
}

func (this *HttpLog) Latency(data int64) *HttpLog {
	this.data = append(this.data, zap.Int64("latency", data))
	return this
}

func (this *HttpLog) Path(data string) *HttpLog {
	this.data = append(this.data, zap.String("path", data))
	return this
}

func (this *HttpLog) Size(data int) *HttpLog {
	this.data = append(this.data, zap.Int("size", data))
	return this
}

func (this *HttpLog) UserId(data int64) *HttpLog {
	this.data = append(this.data, zap.Int64("userid", data))
	return this
}

func (this *HttpLog) Trace(data string) *HttpLog {
	this.data = append(this.data, zap.String(TRACEID, data))
	return this
}

func (this *HttpLog) ErrData(err error) *HttpLog {
	this.data = append(this.data, zap.Error(err))
	return this
}

func (this *HttpLog) JSON(data interface{}) {
	d, _ := json.Marshal(data)
	this.data = append(this.data, zap.String(JsonFormat, string(d)))
}

func (this *HttpLog) Info() {
	this.l.Info(this.msg, this.data...)
}

func (this *HttpLog) Err() {
	this.l.Error(this.msg, this.data...)
}

func (this *HttpLog) Debug() {
	this.l.Log(zapcore.DebugLevel, this.msg, this.data...)
}

func (this *HttpLog) Panic() {
	this.l.Panic(this.msg, this.data...)
}
