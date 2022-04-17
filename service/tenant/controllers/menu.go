/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/15 23:20
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/15 23:20
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

func init() {
	response.AppendController(&Menu{
		DefaultController: response.DefaultController{
			TableName: "Menu",
		},
	})
}

type Menu struct {
	response.DefaultController
}

func (Menu) Path() string {
	return "/menu"
}

func (e Menu) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middlewares.AuthMiddleware(),
	}
}
