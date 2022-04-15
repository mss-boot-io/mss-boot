package response

import (
	"strconv"
	"strings"
)

type Response struct {
	Success      bool   `json:"success,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ShowType     uint8  `json:"showType,omitempty"`
	TraceId      string `json:"traceId,omitempty"`
	Host         string `json:"host,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type Page struct {
	Count    int64 `json:"total"`
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
}

type page struct {
	Page
	response
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.TraceId = id
}

func (e *response) SetMsg(s ...string) {
	e.ErrorMessage += strings.Join(s, ",")
}

func (e *response) SetCode(code int32) {
	e.ErrorCode = strconv.Itoa(int(code))
}

func (e *response) SetSuccess(success bool) {
	e.Success = success
}
