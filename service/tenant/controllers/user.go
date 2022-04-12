/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 13:34
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 13:34
 */

package controllers

import (
	"strings"

	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/errors"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

func init() {
	response.AppendController(&User{})
}

type User struct {
	response.Api
}

func (User) Path() string {
	return "/user"
}

func (e User) Other(r *gin.RouterGroup) {
	r.GET("/current-user", e.GetCurrentUser)
}

// GetCurrentUser 获取当前用户
// @Summary 获取当前用户
// @Description 获取当前用户
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Success 200 {object} auth.Claims
// @Router /tenant/api/v1/current-user [get]
// @Security Bearer
func (e User) GetCurrentUser(c *gin.Context) {
	e.Make(c)
	accessToken := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
	if accessToken == "" {
		e.OK(nil)
	}
	claims, err := auth.ParseJwtToken(accessToken)
	if err != nil {
		e.Err(errors.TenantSvcAccessTokenParseFailed, err)
		return
	}
	claims.AccessToken = accessToken
	e.OK(claims)
}
