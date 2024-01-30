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
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Base action
type Base struct {
	Model     schema.Tabler
	Scope     func(ctx *gin.Context, table schema.Tabler) func(db *gorm.DB) *gorm.DB
	Handlers  gin.HandlersChain
	TreeField string
	Depth     int
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
