/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 17:14:19
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 17:14:19
 */

package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
)

// Pagination pagination params
type Pagination struct {
	Page     int64 `form:"page" query:"page"`
	PageSize int64 `form:"pageSize" query:"pageSize"`
}

func (e *Pagination) GetPage() int64 {
	if e.Page <= 0 {
		return 1
	}
	return e.Page
}

func (e *Pagination) GetPageSize() int64 {
	if e.PageSize <= 0 {
		return 10
	}
	return e.PageSize
}

// Search action
type Search struct {
	Base
	Search response.Searcher
}

// NewSearch new search action
func NewSearch(m mgm.Model, search response.Searcher) *Search {
	return &Search{
		Base:   Base{Model: m},
		Search: search,
	}
}

// String action name
func (*Search) String() string {
	return "search"
}

// Handler action handler
func (e *Search) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := pkg.DeepCopy(e.Search).(response.Searcher)
		api := response.Make(c).Bind(req)
		if api.Error != nil {
			api.Err(http.StatusUnprocessableEntity)
			return
		}
		filter, sort := mgos.MakeCondition(req)
		ops := options.Find()
		ops.SetLimit(req.GetPageSize())
		if len(sort) > 0 {
			ops.SetSort(sort)
		}
		ops.SetSkip(req.GetPageSize() * (req.GetPage() - 1))
		count, err := mgm.Coll(e.Model).CountDocuments(c, filter)
		if err != nil {
			api.Log.Errorf("count items error, %s", err.Error())
			api.AddError(err)
			api.Err(http.StatusInternalServerError)
			return
		}
		result, err := mgm.Coll(e.Model).Find(c, filter, ops)
		if err != nil {
			api.Log.Errorf("find items error, %s", err.Error())
			api.AddError(err)
			api.Err(http.StatusInternalServerError)
			return
		}
		defer result.Close(c)
		items := make([]any, 0, req.GetPageSize())
		for result.Next(c) {
			m := pkg.ModelDeepCopy(e.Model)
			err = result.Decode(m)
			if err != nil {
				api.AddError(err)
				api.Err(http.StatusInternalServerError)
				return
			}
			items = append(items, m)
		}
		api.PageOK(items, count, req.GetPage(), req.GetPageSize())
	}
}
