package actions

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:36:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:36:12
 */

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"

	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// NewDeleteGorm new delete action
func NewDeleteGorm(m schema.Tabler, key string,
	scope func(ctx *gin.Context, table schema.Tabler) func(*gorm.DB) *gorm.DB) *Delete {
	return &Delete{
		Base: Base{ModelGorm: m, Scope: scope},
		Key:  key,
	}
}

func (e *Delete) deleteGorm(c *gin.Context, ids ...string) {
	api := response.Make(c)
	if len(ids) == 0 {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query := gormdb.DB.WithContext(c).Where(fmt.Sprintf("%s IN ?", e.Key), ids)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, e.ModelGorm))
	}
	err := query.Delete(e.ModelGorm).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Delete error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}
