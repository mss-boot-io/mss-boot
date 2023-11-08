package action

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
 * @Date: 2023/9/17 08:58:04
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/17 08:58:04
 */

// Delete action
type Delete struct {
	*Base
}

// NewDelete new delete action
func NewDelete(b *Base) *Delete {
	return &Delete{
		Base: b,
	}
}

// String print action name
func (*Delete) String() string {
	return "delete"
}

// Handler delete action handler
func (e *Delete) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodDelete:
			api := response.Make(c)
			//delete
			m := e.GetModel(c)
			if m == nil {
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
			if err := gormdb.DB.Scopes(m.TableScope, m.URI(c)).Delete(req).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					api.Err(http.StatusNotFound)
					return
				}
				api.AddError(err).Log.ErrorContext(c, "delete error", PathKey, c.Param(PathKey))
				api.Err(http.StatusInternalServerError)
				return
			}
			api.OK(nil)
		default:
			c.AbortWithStatus(http.StatusMethodNotAllowed)
		}
	}
}

//func (e *Delete) GenOpenAPI(m *model.Model) *spec.PathItemProps {
//
//}
