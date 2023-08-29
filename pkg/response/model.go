package response

import (
	"strconv"
	"strings"
)

// Response response
type Response struct {
	Success      bool   `json:"success,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ShowType     uint8  `json:"showType,omitempty"`
	TraceID      string `json:"traceId,omitempty"`
	Host         string `json:"host,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

// Page page
type Page struct {
	Count    int64 `json:"total"`
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
}

type page struct {
	Page
	response
}

// SetData set data
func (e *response) SetData(data interface{}) {
	e.Data = data
}

// Clone clone
func (e *response) Clone() Responses {
	clone := *e
	return &clone
}

// SetTraceID set trace id
func (e *response) SetTraceID(id string) {
	e.TraceID = id
}

// SetMsg set msg
func (e *response) SetMsg(s ...string) {
	e.ErrorMessage += strings.Join(s, ",")
}

// SetCode set code
func (e *response) SetCode(code int32) {
	e.ErrorCode = strconv.Itoa(int(code))
}

// SetSuccess set success
func (e *response) SetSuccess(success bool) {
	e.Success = success
}
