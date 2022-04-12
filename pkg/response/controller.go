/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 10:44 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 10:44 上午
 */

package response

import "github.com/gin-gonic/gin"

// Controllers controllers
var Controllers = make([]Controller, 0)

// Controller controllers
type Controller interface {
	// Path http path
	Path() string
	// Handlers middlewares
	Handlers() []gin.HandlerFunc
	// Create create
	Create(*gin.Context)
	// Update update
	Update(*gin.Context)
	// Delete delete
	Delete(*gin.Context)
	// Get get
	Get(*gin.Context)
	// List list
	List(*gin.Context)
	// Other other
	Other(*gin.RouterGroup)
}

// AppendController add controllers to Controllers
func AppendController(c Controller) {
	Controllers = append(Controllers, c)
}
