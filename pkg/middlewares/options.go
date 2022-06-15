/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/22 11:02
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/22 11:02
 */

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Options 跨域请求
func Options() gin.HandlerFunc {
	return func(c *gin.Context) {

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//	return
		//}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errorMessage": "Not Found"})
	}
}
