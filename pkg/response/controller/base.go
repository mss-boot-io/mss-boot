/*
 * @Author: lwnmengjing
 * @Date: 2023/1/26 01:15:07
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/26 01:15:07
 */

package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/response"
)

var DefaultKey = "id"

type Base struct {
}

func (e *Base) Path() string {
	return ""
}

func (e *Base) Handlers() gin.HandlersChain {
	return nil
}

func (e *Base) GetAction(_ string) response.Action {
	return nil
}

func (e *Base) Other(_ *gin.RouterGroup) {}

// GetKey get key
func (e *Base) GetKey() string {
	return DefaultKey
}
