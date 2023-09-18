package virtual

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/17 08:06:51
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/17 08:06:51
 */

// Create action
type Create struct {
	*Base
}

// NewCreate new create action 真实的
func NewCreate(b *Base) *Create {
	return &Create{
		Base: b,
	}
}

// String print action name
func (*Create) String() string {
	return "create"
}

// Handler create action handler
func (e *Create) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			api := response.Make(c)
			//create
			m := e.GetModel(c)
			if m != nil {
				// no set model
				api.Err(http.StatusNotFound)
				return
			}
			req := m.MakeModel()
			m.Default(req)
			if api.Bind(req).Error != nil {
				api.Err(http.StatusUnprocessableEntity)
				return
			}
			if err := gormdb.DB.Scopes(m.TableScope).Create(req).Error; err != nil {
				api.AddError(err)
				api.Log.Errorf("create %s error", c.Param(PathKey))
				api.Err(http.StatusInternalServerError)
				return
			}
			api.OK(nil)
			return
		default:
			c.AbortWithStatus(http.StatusMethodNotAllowed)
		}
	}
}
