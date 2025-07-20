package lg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志格式化
const (
	DATA       = "data"
	DATAEX     = "dataex"
	JsonFormat = "json"
	KIND       = "kind"
	TRACEID    = "traceid"
)

type BaseLog struct {
	*zap.Logger
	kind string
}

func NewBaseLog(l *zap.Logger, kind string) *BaseLog {
	return &BaseLog{l, kind}
}

func (b *BaseLog) Msg(msg ...string) *Format {
	return Msg(b.Logger, b.kind, msg...)
}

func Msg(log *zap.Logger, kind string, msg ...string) *Format {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Format{
		l:    log,
		data: []zapcore.Field{zap.String(KIND, kind)},
		msg:  m,
	}
}

func (b *BaseLog) ApiLog(msg string) *ApiLog {
	return NewApiLog(b.Logger, msg)
}
