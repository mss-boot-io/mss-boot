package actions

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 20:07:22
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 20:07:22
 */

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	mgm "github.com/kamva/mgm/v3"
	"gorm.io/gorm/schema"
)

// Base action
type Base struct {
	ModelMgm  mgm.Model
	ModelGorm schema.Tabler
	Scope     func(ctx *gin.Context, table schema.Tabler) func(db *gorm.DB) *gorm.DB
}

// String string
func (*Base) String() string {
	return "base"
}

// Handler action handler
func (*Base) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}
