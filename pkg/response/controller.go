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
	Handlers() gin.HandlersChain
	// GetAction get action
	GetAction(string) Action
	// Other handler
	Other(*gin.RouterGroup)
	// GetKey get key
	GetKey() string
}

// AppendController add controllers to Controllers
func AppendController(c Controller) {
	Controllers = append(Controllers, c)
}
