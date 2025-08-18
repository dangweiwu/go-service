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

// 格式化日志
func (b *BaseLog) Msg(msg ...string) *Format {

	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Format{
		l:    b.Logger,
		data: []zapcore.Field{zap.String(KIND, b.kind)},
		msg:  m,
	}

}

// http协议专用日志
func (b *BaseLog) Http() *HttpLog {

	return &HttpLog{
		l:    b.Logger,
		msg:  "",
		data: []zapcore.Field{zap.String(KIND, b.kind)},
	}

}
