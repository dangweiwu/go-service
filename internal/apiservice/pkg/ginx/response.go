package ginx

type ErrKind string

const (
	CODE ErrKind = "code" //需要映射 code映射成相关中文名 多语言支持
	MSG  ErrKind = "msg"  //无需映射 直接显示
)

type ErrResponse struct {
	Kind ErrKind `json:"kind"` //kind=msg|code 类型
	Msg  string  `json:"msg"`  //说明语句或者异常代码
	Data string  `json:"data"` //data补充数据
}

func (this ErrResponse) Error() string {
	return this.Msg
}

type Response struct {
	Data interface{} `json:"data"` // 返回数据
}
