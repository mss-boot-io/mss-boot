package response

import (
	"github.com/mss-boot-io/mss-boot/core/errcode"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg"
)

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code errcode.ErrorCoder, err error, msg ...string) {
	res := Default.Clone()
	if msg == nil {
		msg = make([]string, 0)
	}
	msg = append([]string{code.String()}, msg...)
	if err != nil {
		msg = append(msg, err.Error())
	}
	res.SetMsg(msg...)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(code.Code())
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg ...string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	res.SetMsg(msg...)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	switch c.Request.Method {
	case http.MethodDelete:
		c.AbortWithStatusJSON(http.StatusNoContent, res)
		return
	case http.MethodPost:
		c.AbortWithStatusJSON(http.StatusCreated, res)
		return
	default:
		c.AbortWithStatusJSON(http.StatusOK, res)
	}
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int64, pageIndex int, pageSize int, msg ...string) {
	var res page
	res.Count = count
	res.Current = pageIndex
	res.PageSize = pageSize
	res.response.SetData(result)
	res.response.SetMsg(msg...)
	res.response.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.response.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
