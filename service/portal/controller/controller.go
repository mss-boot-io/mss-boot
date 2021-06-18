/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 10:44 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 10:44 上午
 */

package controller

import "github.com/gin-gonic/gin"

var Controllers = make([]Controller, 0)

type Controller interface {
	Path() string
	Handlers() []gin.HandlerFunc
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}
