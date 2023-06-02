package antd

import (
	"fmt"
	"strings"

	resp "github.com/mss-boot-io/mss-boot/pkg/response"
)

const (
	// Silent success
	Silent = "0"
	// MessageWarn message warn
	MessageWarn = "1"
	// MessageError message error
	MessageError = "2"
	// Notification notification
	Notification = "4"
	// Page page
	Page = "9"
)

// Response response
type Response struct {
	Success      bool   `json:"success,omitempty"`      // if request is success
	ErrorCode    string `json:"errorCode,omitempty"`    // code for errorType
	ErrorMessage string `json:"errorMessage,omitempty"` // message display to user
	ShowType     string `json:"showType,omitempty"`     // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	TraceID      string `json:"traceId,omitempty"`      // Convenient for back-end Troubleshooting: unique request ID
	Host         string `json:"host,omitempty"`         // onvenient for backend Troubleshooting: host of current access server
}
type response struct {
	Response
	Data interface{} `json:"data,omitempty"` // response data
}

// Pages 分页数据
type Pages struct {
	Total    int `json:"total,omitempty"`
	Current  int `json:"current,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type pages struct {
	Pages
	Data interface{} `json:"data,omitempty"`
}

// SetCode 设置错误码
func (e *response) SetCode(code int32) {
	switch code {
	case 200, 0:
	default:
		e.ErrorCode = fmt.Sprintf("C%d", code)
	}
}

// SetTraceID 设置请求ID
func (e *response) SetTraceID(id string) {
	e.TraceID = id
}

// SetMsg 设置错误信息
func (e *response) SetMsg(msg ...string) {
	e.ErrorMessage = strings.Join(msg, ",")
}

// SetData 设置返回数据
func (e *response) SetData(data interface{}) {
	e.Data = data
}

// SetSuccess 设置是否成功
func (e *response) SetSuccess(success bool) {
	e.Success = success
}

// Clone 复制当前对象
func (e *response) Clone() resp.Responses {
	clone := *e
	return &clone
}
