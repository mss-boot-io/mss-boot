package gorm

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 20:07:22
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 20:07:22
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Base action
type Base struct {
	opts *Options
}

// String string
func (*Base) String() string {
	return "base"
}

// Handler action handler
func (*Base) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
	return gin.HandlersChain{h}
}
