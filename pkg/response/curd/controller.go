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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
)

type DefaultController struct {
	response.Api
	TableName string
	Auth      bool
	UpdateReq UpdateRequester
	DeleteReq DeleteRequester
	GetReq    GetRequester
	ListReq   ListRequester
	CreateReq CreateRequester
	GetResp   interface{}
	ListResp  interface{}
}

func (e *DefaultController) Path() string {
	return strings.ReplaceAll(strings.ToLower(e.TableName), "_", "-")
}

func (e *DefaultController) Handlers() []gin.HandlerFunc {
	ms := make([]gin.HandlerFunc, 0)
	if e.Auth {
		ms = append(ms, response.AuthHandler)
	}
	return ms
}

func (e DefaultController) Create(c *gin.Context) {
	req := e.CreateReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	if _, err = mongodb.DB.Collection(e.TableName).InsertOne(c, req); err != nil {
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Update(c *gin.Context) {
	req := e.UpdateReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.GetID())
	_, err = mongodb.DB.Collection(e.TableName).UpdateByID(c, objID, bson.D{{"$set",
		e.getUpdateD(req)}})
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Delete(c *gin.Context) {
	req := e.UpdateReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.GetID())
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	if _, err = mongodb.DB.Collection(e.TableName).DeleteOne(c, bson.M{"_id": objID}); err != nil {
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(nil)
}

func (e DefaultController) Get(c *gin.Context) {
	req := e.GetReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.GetID())
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	resp := e.GetResp
	if err = mongodb.DB.Collection(e.TableName).FindOne(c, bson.M{"_id": objID}).Decode(resp); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Err(http.StatusNotFound, err)
			return
		}
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.OK(resp)
}

func (e DefaultController) List(c *gin.Context) {
	req := e.ListReq
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
	count, err = mongodb.DB.Collection(e.TableName).CountDocuments(c, filter)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	result, err := mongodb.DB.Collection(e.TableName).Find(c, filter, ops)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	defer result.Close(c)
	list := e.ListResp
	err = result.All(c, &list)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.PageOK(list, count, req.GetPage(), req.GetPageSize())
}

func (e *DefaultController) getUpdateD(data interface{}) bson.D {
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
