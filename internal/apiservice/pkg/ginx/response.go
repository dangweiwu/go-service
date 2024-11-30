package ginx

type ErrKind string

const (
	CODE ErrKind = "code" //需要映射 code映射成相关中文名 多语言支持
	MSG  ErrKind = "msg"  //无需映射 直接显示
)

// @doc | hd.ErrResponse
type ErrResponse struct {
	Kind ErrKind `json:"kind" doc:"|d 类型 |c code msg |t string"`
	Msg  string  `json:"msg" doc:"|d 消息 |c 说明语句或者异常代码 |t string"`
	Data string  `json:"data" doc:"|d 代码 |c data补充数据 |t string"`
}

func (this ErrResponse) Error() string {
	return this.Msg
}

// @doc | hd.Response
type Response struct {
	Data interface{} `json:"data" doc:"|t any |c 响应数据 参考Data定义或说明"`
}
