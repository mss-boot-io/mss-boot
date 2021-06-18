/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 9:19 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 9:19 上午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"portal/router"
)

// @title portal API
// @version 2.0.0
// @description admin-go接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	router.Init(r)
	r.Run(":8080")
}
