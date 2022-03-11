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
	c.AbortWithStatusJSON(http.StatusOK, res)
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
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int64, pageIndex int, pageSize int, msg ...string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg...)
}
