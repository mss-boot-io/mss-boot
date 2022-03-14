/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 14:23
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 14:23
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "tenant/controllers"
	_ "tenant/docs"
	"tenant/middleware"
)

func Init(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	middleware.Init(v1)
	var e *gin.RouterGroup
	for i := range response.Controllers {
		response.Controllers[i].Other(v1)
		e = v1.Group(response.Controllers[i].Path(), response.Controllers[i].Handlers()...)
		e.GET("/:id", response.Controllers[i].Get)
		e.POST("", response.Controllers[i].Create)
		e.DELETE("/:id", response.Controllers[i].Delete)
		e.PUT("/:id", response.Controllers[i].Update)
		e.GET("", response.Controllers[i].List)
	}
}
