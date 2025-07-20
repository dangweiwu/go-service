package lg

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ApiLog struct {
	l    *zap.Logger
	msg  string
	data []zapcore.Field
}

func NewApiLog(log *zap.Logger, msg string) *ApiLog {
	return &ApiLog{
		l:    log,
		msg:  msg,
		data: []zapcore.Field{zap.String(KIND, "api")},
	}
}

func (this *ApiLog) Status(data int) *ApiLog {
	this.data = append(this.data, zap.Int("status", data))
	return this
}

func (this *ApiLog) Latency(data int64) *ApiLog {
	this.data = append(this.data, zap.Int64("latency", data))
	return this
}

func (this *ApiLog) Path(data string) *ApiLog {
	this.data = append(this.data, zap.String("path", data))
	return this
}

func (this *ApiLog) Size(data int) *ApiLog {
	this.data = append(this.data, zap.Int("size", data))
	return this
}

func (this *ApiLog) UserId(data int64) *ApiLog {
	this.data = append(this.data, zap.Int64("userid", data))
	return this
}

func (this *ApiLog) Trace(data string) *ApiLog {
	this.data = append(this.data, zap.String(TRACEID, data))
	return this
}

func (this *ApiLog) ErrData(err error) *ApiLog {
	this.data = append(this.data, zap.Error(err))
	return this
}

func (this *ApiLog) JSON(data interface{}) {
	d, _ := json.Marshal(data)
	this.data = append(this.data, zap.String(JsonFormat, string(d)))
}

func (this *ApiLog) Info() {
	this.l.Info(this.msg, this.data...)
}

func (this *ApiLog) Err() {
	this.l.Error(this.msg, this.data...)
}

func (this *ApiLog) Debug() {
	this.l.Log(zapcore.DebugLevel, this.msg, this.data...)
}

func (this *ApiLog) Panic() {
	this.l.Panic(this.msg, this.data...)
}
