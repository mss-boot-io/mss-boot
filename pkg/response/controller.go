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

// Controller controller
type Controller interface {
	// Path http path
	Path() string
	// Handlers middlewares
	Handlers() []gin.HandlerFunc
	// Create create
	Create(c *gin.Context)
	// Update update
	Update(c *gin.Context)
	// Delete delete
	Delete(c *gin.Context)
	// Get get
	Get(c *gin.Context)
	// List list
	List(c *gin.Context)
	// Other other
	Other(*gin.RouterGroup)
}

// AppendController add controller to Controllers
func AppendController(c Controller) {
	Controllers = append(Controllers, c)
}
