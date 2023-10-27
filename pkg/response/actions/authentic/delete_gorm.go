package authentic

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:36:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:36:12
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"

	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// NewDeleteGorm new delete action
func NewDeleteGorm(m schema.Tabler, key string) *Delete {
	return &Delete{
		Base: Base{ModelGorm: m},
		Key:  key,
	}
}

func (e *Delete) deleteGorm(c *gin.Context, ids ...string) {
	api := response.Make(c)
	if len(ids) == 0 {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err := gormdb.DB.Delete(e.ModelGorm, ids).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Delete error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}
