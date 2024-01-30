package gorm

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/9 00:57:48
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/9 00:57:48
 */

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/gorms"
	"gorm.io/gorm"
)

// Search action
type Search struct {
	Base
	Search response.Searcher
}

// String action name
func (*Search) String() string {
	return "search"
}

// NewSearch new search action
func NewSearch(b Base, search response.Searcher) *Search {
	return &Search{
		Base:   b,
		Search: search,
	}
}

// Handler action handler
func (e *Search) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		if e.Model == nil {
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
			return
		}
		e.search(c)
	}
	if e.Handlers != nil {
		return append(e.Handlers, h)
	}
	return gin.HandlersChain{h}
}

func (e *Search) search(c *gin.Context) {
	req := pkg.DeepCopy(e.Search).(response.Searcher)
	api := response.Make(c).Bind(req)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	db := gormdb.DB
	m := pkg.TablerDeepCopy(e.Model)

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

	if e.Base.TreeField != "" && e.Base.Depth > 0 {
		treeFields := make([]string, 0, e.Base.Depth)
		for i := 0; i < e.Base.Depth; i++ {
			treeFields[i] = e.Base.TreeField
		}
		query = query.Preload(strings.Join(treeFields, "."))
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
		m = pkg.TablerDeepCopy(e.Model)
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
