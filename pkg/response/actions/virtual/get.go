package virtual

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/17 09:00:47
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/17 09:00:47
 */

// Get action
type Get struct {
	*Base
}

// NewGet new get action
func NewGet(b *Base) *Get {
	return &Get{
		Base: b,
	}
}

// String print action name
func (*Get) String() string {
	return "get"
}

// Handler get action handler
func (e *Get) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			api := response.Make(c)
			//get
			m := e.GetModel(c)
			if m == nil {
				// no set model
				api.Err(http.StatusNotFound)
				return
			}
			req := m.MakeModel()
			if api.Bind(req).Error != nil {
				api.Err(http.StatusUnprocessableEntity)
				return
			}
			if err := gormdb.DB.Scopes(m.TableScope, m.URI(c)).First(req).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					api.Err(http.StatusNotFound)
					return
				}
				api.AddError(err)
				api.Log.Errorf("get %s error", c.Param(PathKey))
				api.Err(http.StatusInternalServerError)
				return
			}
			api.OK(req)
			return
		default:
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
	}
}
