package gorm

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:38:45
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:38:45
 */

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// Get action
type Get struct {
	Base
	Key string
}

// String action name
func (*Get) String() string {
	return "get"
}

// NewGet new get action
func NewGet(b Base, key string) *Get {
	return &Get{
		Base: b,
		Key:  key,
	}
}

func (e *Get) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		if e.Model == nil {
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
			return
		}
		e.get(c, e.Key)
	}
	if e.Handlers != nil {
		return append(e.Handlers, h)
	}
	return gin.HandlersChain{h}
}

// get one record by id
func (e *Get) get(c *gin.Context, key string) {
	api := response.Make(c)
	m := pkg.TablerDeepCopy(e.Model)
	preloads := c.QueryArray("preloads[]")
	query := gormdb.DB.Model(m).Where("id = ?", c.Param(key))

	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	err := query.First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.Err(http.StatusNotFound)
			return
		}
		api.AddError(err).Log.ErrorContext(c, "Get error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
