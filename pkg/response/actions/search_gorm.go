package actions

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/9 00:57:48
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/9 00:57:48
 */

import (
	"errors"
	"gorm.io/gorm/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/gorms"
	"gorm.io/gorm"
)

// NewSearchGorm new search action
func NewSearchGorm(m Model, search response.Searcher,
	scope func(ctx *gin.Context, table schema.Tabler) func(*gorm.DB) *gorm.DB) *Search {
	return &Search{
		Base:   Base{ModelGorm: m, Scope: scope},
		Search: search,
	}
}

func (e *Search) searchGorm(c *gin.Context) {
	req := pkg.DeepCopy(e.Search).(response.Searcher)
	api := response.Make(c).Bind(req)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	db := gormdb.DB
	m := pkg.TablerDeepCopy(e.ModelGorm)

	var count int64

	query := db.WithContext(c).Model(m).
		Scopes(
			gorms.MakeCondition(req),
			gorms.Paginate(int(req.GetPageSize()), int(req.GetPage())),
		)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	if err := query.Limit(-1).Offset(-1).Count(&count).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		api.AddError(err).Log.ErrorContext(c, "Search error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}

	rows, err := query.Rows()
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Search error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	items := make([]any, 0, req.GetPageSize())
	for rows.Next() {
		m = pkg.TablerDeepCopy(e.ModelGorm)
		err = db.ScanRows(rows, m)
		if err != nil {
			api.AddError(err).Log.ErrorContext(c, "search error", "error", err)
			api.Err(http.StatusInternalServerError)
			return
		}
		items = append(items, m)
	}
	err = query.Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		api.AddError(err).Log.ErrorContext(c, "search error", "error", err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.PageOK(items, count, req.GetPage(), req.GetPageSize(), "search success")
	c.Next()
}
