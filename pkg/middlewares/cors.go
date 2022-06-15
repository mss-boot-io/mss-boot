/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/22 10:34
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/22 10:34
 */

package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"},
		MaxAge:          6 * time.Hour,
	}

	return cors.New(c)
}
