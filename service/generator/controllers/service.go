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
	response.AppendController(&Service{})
}

type Service struct {
	response.Api
}

func (Service) C() *mongo.Collection {
	return (&models.Service{}).C()
}

func (Service) Path() string {
	return "/service"
}

func (e Service) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Create 创建
// @Summary 创建service
// @Description 创建service
// @Tags service
// @Accept  application/json
// @Product application/json
// @Param data body form.ServiceCreateReq true "data"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/service [post]
// @Security Bearer
func (e Service) Create(c *gin.Context) {
	req := &form.ServiceCreateReq{}

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
// @Summary 更新service
// @Description 更新service
// @Tags service
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body form.ServiceUpdateReq true "data"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/service/{id} [put]
// @Security Bearer
func (e Service) Update(c *gin.Context) {
	req := &form.ServiceUpdateReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	_, err = e.C().UpdateByID(c, objID, bson.D{{"$set",
		bson.D{
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
// @Summary 删除service
// @Description 删除service
// @Tags service
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /generator/api/v1/service/{id} [delete]
// @Security Bearer
func (e Service) Delete(c *gin.Context) {
	req := &form.ServiceDeleteReq{}
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
// @Summary 获取service
// @Description 获取service
// @Tags service
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response{data=form.ServiceGetResp}
// @Router /generator/api/v1/service/{id} [get]
// @Security Bearer
func (e Service) Get(c *gin.Context) {
	req := &form.ServiceGetReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.GeneratorSvcParamsInvalid, err)
		return
	}
	resp := &form.ServiceGetResp{}
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
// @Summary 列表service
// @Description 列表service
// @Tags service
// @Accept  application/json
// @Product application/json
// @Param name query string false "名称"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页容量"
// @Success 200 {object} response.Page{data=[]form.ServiceListItem}
// @Router /generator/api/v1/service [get]
// @Security Bearer
func (e Service) List(c *gin.Context) {
	req := &form.ServiceListReq{}
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
	list := make([]form.ServiceListItem, 0)
	err = result.All(c, &list)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
	e.PageOK(list, count, req.Page, req.PageSize)
	return
}

func (e Service) Generate(c *gin.Context) {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcObjectIDInvalid, err)
		return
	}
	m := &models.Service{}
	err = e.C().FindOne(c, bson.M{"_id": objID}).Decode(m)
	if err != nil {
		e.Log.Error(err)
		e.Err(errors.GeneratorSvcOperateDBFailed, err)
		return
	}
}

func (e Service) Other(r *gin.RouterGroup) {
	r.POST("/generate/service/:id", e.Generate)
}
