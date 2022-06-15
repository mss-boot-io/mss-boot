/*
 * @Author: lwnmengjing
 * @Date: 2021/6/23 3:33 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/23 3:33 下午
 */

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

func Init(r *gin.Engine) {
	// init middleware
	response.AuthHandler = (&middlewares.AuthMiddleware{}).AuthMiddleware()
	r.Use(middlewares.Cors())
	r.NoRoute(middlewares.Options())
}
