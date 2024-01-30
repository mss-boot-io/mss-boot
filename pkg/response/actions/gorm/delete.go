package gorm

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:36:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:36:12
 */

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// Delete action
type Delete struct {
	Base
	Key string
}

// NewDelete new delete action
func NewDelete(b Base, key string) *Delete {
	return &Delete{
		Base: b,
		Key:  key,
	}
}

func (e *Delete) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		if e.Model == nil {
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
			return
		}
		ids := make([]string, 0)
		v := c.Param(e.Key)
		if v == "batch" {
			api := response.Make(c).Bind(&ids)
			if api.Error != nil || len(ids) == 0 {
				api.Err(http.StatusUnprocessableEntity)
				return
			}
			e.delete(c, ids...)
			return
		}
		e.delete(c, v)
	}
	if e.Handlers != nil {
		return append(e.Handlers, h)
	}
	return gin.HandlersChain{h}
}

// String action name
func (*Delete) String() string {
	return "delete"
}

func (e *Delete) delete(c *gin.Context, ids ...string) {
	api := response.Make(c)
	if len(ids) == 0 {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query := gormdb.DB.WithContext(c).Where(fmt.Sprintf("%s IN ?", e.Key), ids)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, e.Model))
	}
	err := query.Delete(e.Model).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Delete error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}
