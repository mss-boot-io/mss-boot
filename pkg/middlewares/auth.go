/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/13 22:44
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/13 22:44
 */

package middlewares

import (
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := getTokenFromHeader(c)
		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorCode":    http.StatusUnauthorized,
				"errorMessage": "Unauthorized",
			})
			return
		}
		claims, err := auth.ParseJwtToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorCode":    http.StatusUnauthorized,
				"errorMessage": err,
			})
			return
		}
		claims.AccessToken = accessToken
		c.Set("claims", claims)
		c.Next()
	}
}

// getTokenFromHeader 获取token
func getTokenFromHeader(c *gin.Context) string {
	return strings.ReplaceAll(
		c.GetHeader("Authorization"),
		"Bearer ",
		"")
}
