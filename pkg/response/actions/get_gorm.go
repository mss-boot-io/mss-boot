package actions

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

// NewGetGorm new get action
func NewGetGorm(b Base, key string) *Get {
	return &Get{
		Base: b,
		Key:  key,
	}
}

// getGorm get one record by id
func (e *Get) getGorm(c *gin.Context, key string) {
	api := response.Make(c)
	m := pkg.TablerDeepCopy(e.ModelGorm)
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
