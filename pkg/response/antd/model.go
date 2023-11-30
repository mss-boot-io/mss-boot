package antd

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
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
	Total    int64 `json:"total,omitempty"`
	Current  int64 `json:"current,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type pages struct {
	Pages
	response
}

// SetCode 设置错误码
func (e *response) SetCode(code int) {
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

// SetList 设置返回数据
func (e *response) SetList(data interface{}) {
	e.Data = data
}

// SetStatus 设置是否成功
func (e *response) SetStatus(status string) {
	switch strings.ToLower(status) {
	case "ok", "success", "1", "t", "true":
		e.Success = true
	}
}

// Clone 复制当前对象
func (e *response) Clone() resp.Responses {
	clone := *e
	return &clone
}

func checkContext(c *gin.Context) {
	if c == nil {
		slog.Error("context is nil, please check, e.g. e.Make(c) add your controller function")
		os.Exit(-1)
	}
}

// Error error
func (e *response) Error(c *gin.Context, code int, err error, msg ...string) {
	checkContext(c)
	res := e.Clone()
	if msg == nil {
		msg = make([]string, 0)
	}
	if err != nil {
		msg = append(msg, err.Error())
	}
	res.SetMsg(msg...)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(code)
	res.SetStatus("error")
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(code, res)
}

// OK ok
func (e *response) OK(c *gin.Context, data interface{}) {
	checkContext(c)
	res := e.Clone()
	res.SetList(data)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	switch c.Request.Method {
	case http.MethodDelete:
		res.SetCode(http.StatusNoContent)
		c.AbortWithStatusJSON(http.StatusNoContent, data)
		return
	case http.MethodPost:
		res.SetCode(http.StatusCreated)
		c.AbortWithStatusJSON(http.StatusCreated, data)
		return
	default:
		res.SetCode(http.StatusOK)
		c.AbortWithStatusJSON(http.StatusOK, data)
	}
}

// PageOK page ok
func (e *response) PageOK(c *gin.Context, result interface{}, count, pageIndex, pageSize int64) {
	checkContext(c)
	var res pages
	res.Total = count
	res.Current = pageIndex
	res.PageSize = pageSize
	res.response.SetList(result)
	res.response.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
