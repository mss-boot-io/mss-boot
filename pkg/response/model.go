package response

import (
	"strings"
)

// Response response
type Response struct {
	Success      bool   `json:"success,omitempty"`
	Status       string `json:"status,omitempty"`
	Code         int    `json:"code,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ShowType     uint8  `json:"showType,omitempty"`
	TraceID      string `json:"traceId,omitempty"`
	Host         string `json:"host,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data,omitempty"`
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

// SetList set data
func (e *response) SetList(data interface{}) {
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
func (e *response) SetCode(code int) {
	e.Code = code
}

func (e *response) SetStatus(status string) {
	e.Status = status
}
