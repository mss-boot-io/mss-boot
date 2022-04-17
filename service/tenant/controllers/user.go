/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 13:34
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 13:34
 */

package controllers

import (
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/service/tenant/models"
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

func (e User) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middlewares.AuthMiddleware(),
	}
}

func (e User) Other(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("/current-user", e.GetCurrentUser)
}

// GetCurrentUser 获取当前用户
// @Summary 获取当前用户
// @Description 获取当前用户
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response{data=auth.Claims}
// @Router /tenant/api/v1/current-user [get]
// @Security Bearer
func (e User) GetCurrentUser(c *gin.Context) {
	e.Make(c)
	v, ok := c.Get("claims")
	if !ok {
		e.Err(http.StatusUnauthorized, nil, "claims not found")
		return
	}
	ok = false
	claims, ok := v.(*auth.Claims)
	if !ok {
		e.Err(http.StatusUnauthorized, nil, "claims type error")
		return
	}
	//写入用户信息
	go models.CreateOrUpdateUser(claims)
	e.OK(claims)
}
