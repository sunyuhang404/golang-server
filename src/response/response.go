package response

import "encoding/json"

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息描述
	Data    interface{} `json:"data"`    // 返回数据
}

// 自定义响应数据
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

// 追加数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

// 返回 JSON 格式的错误数据
func (res *Response) ToString() string {
	err := &struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 错误码
var (
	OK          = response(200, "ok")   // 成功
	Error       = response(0, "error")  // 失败
	AuthFail    = response(401, "登录过期") // 登录过期
	ErrorParams = response(0, "error")  // 入参错误
)
