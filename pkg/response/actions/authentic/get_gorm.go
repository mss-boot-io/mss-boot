package authentic

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
	"gorm.io/gorm/schema"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// NewGetGorm new get action
func NewGetGorm(m schema.Tabler, key string) *Get {
	return &Get{
		Base: Base{ModelGorm: m},
		Key:  key,
	}
}

// getGorm get one record by id
func (e *Get) getGorm(c *gin.Context, key string) {
	api := response.Make(c)
	m := pkg.TablerDeepCopy(e.ModelGorm)
	err := gormdb.DB.First(m, "id = ?", c.Param(key)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.Err(http.StatusNotFound)
			return
		}
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
