/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 14:23
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 14:23
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/service/tenant/models"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	_ "github.com/mss-boot-io/mss-boot/service/tenant/controllers"
	_ "github.com/mss-boot-io/mss-boot/service/tenant/docs"
	"github.com/mss-boot-io/mss-boot/service/tenant/middleware"
	_ "github.com/mss-boot-io/mss-boot/service/tenant/models"
)

func Init(path string) *gin.Engine {
	engine := gin.Default()
	middleware.Init(engine)
	r := engine.Group(path)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	var e *gin.RouterGroup
	for i := range response.Controllers {
		response.Controllers[i].Other(v1)
		e = v1.Group(response.Controllers[i].Path(), response.Controllers[i].Handlers()...)
		e.Use(middlewares.Cors())
		e.GET("/:id", response.Controllers[i].Get)
		e.POST("", response.Controllers[i].Create)
		e.DELETE("/:id", response.Controllers[i].Delete)
		e.PUT("/:id", response.Controllers[i].Update)
		e.GET("", response.Controllers[i].List)
	}
	models.Routes = engine.Routes()
	return engine
}
