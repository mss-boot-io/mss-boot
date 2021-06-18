/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 9:42 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 9:42 上午
 */

package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"portal/controller"
	_ "portal/docs"
)

func Init(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	for i := range controller.Controllers {
		e := v1.Group(controller.Controllers[i].Path()).Use(controller.Controllers[i].Handlers()...)
		e.GET("/:id", controller.Controllers[i].Get)
		e.POST("", controller.Controllers[i].Create)
		e.DELETE("/:id", controller.Controllers[i].Delete)
		e.PUT("/:id", controller.Controllers[i].Update)
		e.GET("", controller.Controllers[i].List)
	}
}
