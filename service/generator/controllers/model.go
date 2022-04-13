/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 15:29
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 15:29
 */

package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/errors"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"generator/form"
	"generator/models"
)

func init() {
	response.AppendController(&Model{})
}

type Model struct {
	response.Api
}

func (Model) C() *mongo.Collection {
	return (&models.Model{}).C()
}

func (Model) Path() string {
	return "/model"
}

func (e Model) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Create 创建
// @Summary 创建model
// @Description 创建model
// @Tags model
// @Accept  application/json
// @Product application/json
// @Param data body form.ModelCreateReq true "data"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/model [post]
// @Security Bearer
func (e Model) Create(c *gin.Context) {
	req := &form.ModelCreateReq{}

	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}

	req.CreatedAt = time.Now()
	req.UpdatedAt = req.CreatedAt

	if _, err = e.C().InsertOne(c, req); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			e.Err(errors.GeneratorSvcRecordIsExist, err)
			return
		}
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}

	e.OK(nil)
	return
}

// Update 更新
// @Summary 更新model
// @Description 更新model
// @Tags model
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body form.ModelUpdateReq true "data"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/model/{id} [put]
// @Security Bearer
func (e Model) Update(c *gin.Context) {
	req := &form.ModelUpdateReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	_, err = e.C().UpdateByID(c, objID, bson.D{{"$set",
		bson.D{
			{"service", req.Service},
			{"name", req.Name},
			{"metadata", req.Metadata},
			{"description", req.Description},
			{"updatedAt", time.Now()},
		}}})
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	e.OK(nil)
	return
}

// Delete 删除
// @Summary 删除model
// @Description 删除model
// @Tags model
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/model/{id} [delete]
// @Security Bearer
func (e Model) Delete(c *gin.Context) {
	req := &form.ModelDeleteReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}

	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcObjectIDInvalid, err)
		return
	}
	_, err = e.C().DeleteOne(c, bson.M{"_id": objID})
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		e.Err(errors.GeneratorSvcRecordNotFound, nil)
		return
	}

	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	e.OK(nil)
	return
}

// Get 获取
// @Summary 获取model
// @Description 获取model
// @Tags model
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response{data=form.ModelGetResp}
// @Router /generator/api/v1/model/{id} [get]
// @Security Bearer
func (e Model) Get(c *gin.Context) {
	req := &form.ModelGetReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}
	resp := &form.ModelGetResp{}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcObjectIDInvalid, err)
		return
	}
	err = e.C().FindOne(c, bson.M{"_id": objID}).Decode(&resp)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	e.OK(resp)
	return
}

// List 列表
// @Summary 列表model
// @Description 列表model
// @Tags model
// @Accept  application/json
// @Product application/json
// @Param name query string false "名称"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页容量"
// @Success 200 {object} response.Page{data=[]form.ModelListItem}
// @Router /generator/api/v1/model [get]
// @Security Bearer
func (e Model) List(c *gin.Context) {
	req := &form.ModelListReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}
	filter, sort := mgos.MakeCondition(*req)

	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	ops := options.Find()
	if sort != nil {
		ops.SetSort(sort)
	}
	ops.SetLimit(int64(req.PageSize))
	ops.SetSkip(int64(req.PageSize * (req.Page - 1)))
	var count int64
	count, err = e.C().CountDocuments(c, filter)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	result, err := e.C().Find(c, filter, ops)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	defer result.Close(c)
	list := make([]form.ModelListItem, 0)
	err = result.All(c, &list)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	e.PageOK(list, count, req.Page, req.PageSize)
	return
}
