package actions

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:30:34
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:30:34
 */

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// NewControlGorm new control action
func NewControlGorm(m schema.Tabler, key string,
	scope func(ctx *gin.Context, table schema.Tabler) func(*gorm.DB) *gorm.DB) *Control {
	return &Control{
		Base: Base{ModelGorm: m, Scope: scope},
		Key:  key,
	}
}

func (e *Control) createGorm(c *gin.Context) {
	m := pkg.TablerDeepCopy(e.ModelGorm)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err := gormdb.DB.WithContext(c).Create(m).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Create error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}

func (e *Control) updateGorm(c *gin.Context) {
	m := pkg.TablerDeepCopy(e.ModelGorm)
	id := c.Param(e.Key)
	api := response.Make(c)
	if id == "" {
		api.AddError(errors.New("id is empty"))
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query := gormdb.DB.Where(e.Key, id)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	//find object
	err := query.First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.AddError(fmt.Errorf("%s(%s) record not found", e.Key, id))
			api.Err(http.StatusNotFound)
			return
		}
		api.AddError(err).Log.ErrorContext(c, "Update error", "error", err.Error())
		api.Err(http.StatusInternalServerError)
		return
	}

	api = api.Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query = gormdb.DB.WithContext(c)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	err = query.Save(m).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Update error", "error", err.Error())
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
