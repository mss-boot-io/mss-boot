/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/5/18 23:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/5/18 23:12
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/response"
)

type DocumentController struct {
	response.Api
}

func (e *DocumentController) Path() string {
	return "/document"
}

func (e DocumentController) All(c *gin.Context) {
	data := map[string]interface{}{
		"menu": map[string]interface{}{
			"list": "",
		},
	}
	e.Make(c)
	e.OK(data)
}
