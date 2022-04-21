/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 13:34
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 13:34
 */

package controllers

import (
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

func init() {
	e := &User{}
	e.Auth = true
	response.AppendController(e)
}

type User struct {
	curd.DefaultController
}

func (e User) Other(r *gin.RouterGroup) {
	r.Use(response.AuthHandler)
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
	v, ok := c.Get("user")
	if !ok {
		e.Err(http.StatusUnauthorized, nil, "claims not found")
		return
	}
	ok = false
	user, ok := v.(*middlewares.User)
	if !ok {
		e.Err(http.StatusUnauthorized, nil, "claims type error")
		return
	}
	//写入用户信息
	//go models.CreateOrUpdateUser(user)
	e.OK(user)
}
