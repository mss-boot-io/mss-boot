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

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// NewDeleteGorm new delete action
func NewDeleteGorm(b Base, key string) *Delete {
	return &Delete{
		Base: b,
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
