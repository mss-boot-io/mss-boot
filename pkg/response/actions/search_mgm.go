package actions

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 17:14:19
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 17:14:19
 */

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	mgm "github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetPage get page
func (e *Pagination) GetPage() int64 {
	if e.Page <= 0 {
		return 1
	}
	return e.Page
}

// GetPageSize get page size
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

// NewSearchMgm new search action
func NewSearchMgm(m mgm.Model, search response.Searcher) *Search {
	return &Search{
		Base:   Base{ModelMgm: m},
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
		if e.ModelMgm != nil {
			e.searchMgm(c)
			return
		}
		if e.ModelGorm != nil {
			e.searchGorm(c)
			return
		}
		response.Error(c,
			http.StatusNotImplemented,
			fmt.Errorf("not implemented"))
	}
}

func (e *Search) searchMgm(c *gin.Context) {
	req := pkg.DeepCopy(e.Search).(response.Searcher)
	api := response.Make(c).Bind(req)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	filter, sort := mgos.MakeCondition(req)

	count, err := mgm.Coll(e.ModelMgm).CountDocuments(c, filter)
	if err != nil {
		api.Log.Errorf("count items error, %s", err.Error())
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	linkConfigs := getLinkTag(e.ModelMgm)
	if len(linkConfigs) == 0 {
		ops := options.Find()
		ops.SetLimit(req.GetPageSize())
		if len(sort) > 0 {
			ops.SetSort(sort)
		}
		ops.SetSkip(req.GetPageSize() * (req.GetPage() - 1))

		result, err := mgm.Coll(e.ModelMgm).Find(c, filter, ops)
		if err != nil {
			api.AddError(err).Log.Errorf("find items error, %s", err.Error())
			api.Err(http.StatusInternalServerError)
			return
		}
		defer result.Close(c)
		items := make([]any, 0, req.GetPageSize())
		for result.Next(c) {
			m := pkg.ModelDeepCopy(e.ModelMgm)
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
	//use Aggregate
	//https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/
	pipeline := bson.A{
		//builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
	}
	for i := range linkConfigs {
		pipeline = append(pipeline,
			builder.S(
				builder.Lookup(
					linkConfigs[i].CollectionName,
					linkConfigs[i].LocalField,
					field.ID, linkConfigs[i].ForeignField)))
	}
	//limit skip sort
	pipeline = append(pipeline, bson.D{
		{Key: "$limit", Value: req.GetPageSize()},
	}, bson.D{
		{Key: "$skip", Value: req.GetPageSize() * (req.GetPage() - 1)},
	}, bson.D{
		{Key: "$sort", Value: sort},
	})
	result, err := mgm.Coll(e.ModelMgm).Aggregate(c, pipeline)
	if err != nil {
		api.Log.Errorf("find items error, %s", err.Error())
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	defer result.Close(c)
	items := make([]any, 0, req.GetPageSize())
	for result.Next(c) {
		m := pkg.ModelDeepCopy(e.ModelMgm)
		var bm bson.M
		err = result.Decode(bm)
		if err != nil {
			api.AddError(err)
			api.Err(http.StatusInternalServerError)
			return
		}
		if bm == nil {
			continue
		}
		//todo bson.M to model
		err = BsonMTransferModel(bm, m)
		if err != nil {
			api.AddError(err).Log.Errorf("transfer bson.M to model error, %s", err.Error())
			return
		}
		items = append(items, m)
	}
	api.PageOK(items, count, req.GetPage(), req.GetPageSize())
}

// LinkConfig link config
type LinkConfig struct {
	// FieldName field name
	FieldName string
	// CollectionName collection name
	CollectionName string
	// LocalField local field
	LocalField string
	// ForeignField foreign field
	ForeignField string
}

// BsonMTransferModel bson.M to model
func BsonMTransferModel(bm bson.M, model any) error {
	typeOf := reflect.TypeOf(model).Elem()
	valueOf := reflect.ValueOf(model).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		f := typeOf.Field(i)
		tagBson := f.Tag.Get("bson")
		if tagBson == "-" {
			continue
		}
		if tagBson == "" {
			tagBson = f.Name
		}
		if f.Type.Kind() == reflect.Struct {
			if f.Type.Name() == "mgm.DefaultModel" {
				dm := valueOf.Field(i).Interface().(mgm.DefaultModel)
				dm.SetID(bm["_id"].(primitive.ObjectID))
				dm.UpdatedAt = bm["updated_at"].(time.Time)
				dm.CreatedAt = bm["created_at"].(time.Time)
				continue
			}
			if strings.Contains(tagBson, "inline") {
				err := BsonMTransferModel(bm, valueOf.Field(i).Interface())
				if err != nil {
					return err
				}
				continue
			}

			continue
		}
		if f.Type.Kind() == reflect.Array {
			// transfer bson.M to array model
			switch ms := bm[tagBson].(type) {
			case []bson.M:
				bsonBytes, _ := bson.Marshal(ms)
				err := bson.Unmarshal(bsonBytes, valueOf.Field(i).Interface())
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("type %s not is array", reflect.TypeOf(ms).String())
			}
			continue
		}
		valueOf.Field(i).Set(reflect.ValueOf(bm[tagBson]))
	}
	return nil
}

// getLinkTag get link tag from object
func getLinkTag(model any) []LinkConfig {
	configs := make([]LinkConfig, 0)
	typeOf := reflect.TypeOf(model).Elem()
	valueOf := reflect.ValueOf(model).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		f := typeOf.Field(i)
		if f.Type.Kind() == reflect.Struct {
			if f.Type.String() == "mgm.DefaultModel" {
				continue
			}
			vm, ok := valueOf.Field(i).Interface().(mgm.Model)
			if !ok {
				continue
			}
			m := pkg.ModelDeepCopy(vm)
			configs = append(configs, LinkConfig{
				FieldName:      f.Tag.Get("bson"),
				CollectionName: mgm.Coll(m).Name(),
				LocalField:     f.Name + field.ID,
				ForeignField:   field.ID,
			})
		}
	}
	return configs
}
