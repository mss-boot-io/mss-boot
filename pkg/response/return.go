package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg"
)

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg ...string) {
	res := Default.Clone()
	if msg == nil {
		msg = make([]string, 0)
	}
	if err != nil {
		msg = append(msg, err.Error())
	}
	res.SetMsg(msg...)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(int32(code))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(code, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg ...string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	res.SetMsg(msg...)
	res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	switch c.Request.Method {
	case http.MethodDelete:
		res.SetCode(http.StatusNoContent)
		c.AbortWithStatusJSON(http.StatusNoContent, res)
		return
	case http.MethodPost:
		res.SetCode(http.StatusCreated)
		c.AbortWithStatusJSON(http.StatusCreated, res)
		return
	default:
		res.SetCode(http.StatusOK)
		c.AbortWithStatusJSON(http.StatusOK, res)
	}
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int64, pageIndex int64, pageSize int64, msg ...string) {
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
