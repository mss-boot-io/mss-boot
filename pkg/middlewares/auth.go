/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/13 22:44
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/13 22:44
 */

package middlewares

import (
	"errors"
	"github.com/mss-boot-io/mss-boot/pkg/store"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/response"
)

type AuthMiddleware struct {
	response.Api
}

// AuthMiddleware 认证中间件
func (e AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		e.Make(c)
		accessToken := getTokenFromHeader(c)
		if accessToken == "" {
			e.Err(http.StatusUnauthorized, errors.New("token is empty"))
			return
		}
		client, err := store.DefaultOAuth2Store.
			GetClientByDomain(c.Request.Context(), c.Request.Host)
		if err != nil {
			e.Err(http.StatusUnauthorized, err)
			return
		}
		provider, err := oidc.NewProvider(c, client.GetIssuer())
		if err != nil {
			e.Err(http.StatusUnauthorized, err)
			return
		}
		idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: client.GetClientID()})
		idToken, err := idTokenVerifier.Verify(c, accessToken)
		if err != nil {
			e.Err(http.StatusUnauthorized, err)
			return
		}
		user := &User{}
		err = idToken.Claims(user)
		if err != nil {
			e.Err(http.StatusUnauthorized, err)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

// getTokenFromHeader 获取token
func getTokenFromHeader(c *gin.Context) string {
	return strings.ReplaceAll(strings.ReplaceAll(
		c.GetHeader("Authorization"),
		"Bearer ",
		""),
		"bearer",
		"")
}
