/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:38:45
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:38:45
 */

package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// getGorm get one record by id
func (e *Get) getGorm(c *gin.Context, key string) {
	api := response.Make(c)
	m := pkg.TablerDeepCopy(e.ModelGorm)
	err := gormdb.DB.First(m, c.Param(key)).Error
	if err != nil {
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
