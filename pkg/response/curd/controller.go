/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/15 23:31
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/15 23:31
 */

package curd

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
)

type DefaultController struct {
	response.Api
	Auth   bool
	Model  mgm.Model
	Search Searcher
}

func (e DefaultController) Path() string {
	return strings.ReplaceAll(strings.ToLower(mgm.CollName(e.Model)), "_", "-")
}

func (e DefaultController) Handlers() []gin.HandlerFunc {
	ms := make([]gin.HandlerFunc, 0)
	if e.Auth {
		ms = append(ms, response.AuthHandler)
	}
	return ms
}

func (e DefaultController) Create(c *gin.Context) {
	m := e.NewModel()
	err := e.Make(c).Bind(m).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	err = mgm.Coll(e.Model).CreateWithCtx(c, m)
	if err != nil {
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Update(c *gin.Context) {
	m := e.NewModel()
	err := e.Make(c).Bind(m).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	id, err := e.Model.PrepareID(c.Param("id"))
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	m.SetID(id)
	err = mgm.Coll(m).UpdateWithCtx(c, m)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Err(http.StatusNotFound, err)
			return
		}
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Delete(c *gin.Context) {
	e.Make(c)
	id := c.Param("id")
	if id == "batch" {
		//batch delete
		e.BatchDelete(c)
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	if _, err = mgm.Coll(e.Model).DeleteOne(c, bson.M{"_id": objID}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Err(http.StatusNotFound, err)
			return
		}
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

// BatchDelete todo
func (e DefaultController) BatchDelete(c *gin.Context) {
	req := make([]string, 0)
	err := e.Make(c).Bind(&req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	_, err = mgm.Coll(e.Model).DeleteMany(c, bson.M{"_id": bson.M{"$in": req}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.OK(nil)
			return
		}
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Get(c *gin.Context) {
	e.Make(c)
	m := e.NewModel()
	err := mgm.Coll(e.Model).FindByIDWithCtx(c, c.Param("id"), m)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Err(http.StatusNotFound, err)
			return
		}
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(m)
}

func (e DefaultController) List(c *gin.Context) {
	req := e.Search
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	filter, sort := mgos.MakeCondition(req)

	ops := options.Find()
	if sort != nil {
		ops.SetSort(sort)
	}
	ops.SetLimit(req.GetPageSize())
	ops.SetSkip(req.GetPageSize() * (req.GetPage() - 1))
	var count int64
	count, err = mgm.Coll(e.Model).CountDocuments(c, filter)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	result, err := mgm.Coll(e.Model).Find(c, filter, ops)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	defer result.Close(c)
	items := make([]any, 0)
	for {
		m := e.NewModel()
		if !result.Next(c) {
			break
		}
		err = result.Decode(m)
		if err != nil {
			e.Err(http.StatusInternalServerError, err)
			return
		}
		items = append(items, m)
	}
	e.PageOK(items, count, req.GetPage(), req.GetPageSize())
}

func (e DefaultController) NewModel() mgm.Model {
	return reflect.New(reflect.TypeOf(e.Model).Elem()).Interface().(mgm.Model)
}

func (e DefaultController) getUpdateD(data any) bson.D {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var tag string
	var ok bool
	update := make(bson.D, 0)
	for i := 0; i < t.NumField(); i++ {
		tag, ok = t.Field(i).Tag.Lookup("bson")
		if !ok {
			continue
		}
		update = append(update, bson.E{Key: tag, Value: v.Field(i).Interface()})
	}
	return update
}
