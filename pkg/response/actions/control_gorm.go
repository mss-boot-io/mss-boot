/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:30:34
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:30:34
 */

package actions

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"gorm.io/gorm/schema"
)

// NewControlGorm new control action
func NewControlGorm(m schema.Tabler, key string) *Control {
	return &Control{
		Base: Base{ModelGorm: m},
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
	err := gormdb.DB.Create(m).Error
	if err != nil {
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
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
	//find object
	err := gormdb.DB.Where(e.Key, id).First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.AddError(fmt.Errorf("%s(%s) record not found", e.Key, id))
			api.Err(http.StatusNotFound)
			return
		}
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}

	api = api.Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err = gormdb.DB.Save(m).Error
	if err != nil {
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}
