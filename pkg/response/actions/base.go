/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 20:07:22
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 20:07:22
 */

package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

type Base struct {
	Model mgm.Model
}

func (*Base) String() string {
	return "base"
}

// Handler action handler
func (*Base) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}
