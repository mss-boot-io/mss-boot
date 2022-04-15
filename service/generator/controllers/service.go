/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 15:29
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 15:29
 */

package controllers

import (
	"generator/form"
	"generator/models"
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func init() {
	e := &Service{}
	e.TableName = "service"
	e.Auth = true
	e.CreateReq = &form.ServiceCreateReq{}
	e.UpdateReq = &form.ServiceUpdateReq{}
	e.DeleteReq = &form.ServiceDeleteReq{}
	e.GetReq = &form.ServiceGetReq{}
	e.GetResp = &form.ServiceGetResp{}
	e.ListReq = &form.ServiceListReq{}
	e.ListResp = make([]form.ServiceListItem, 0)
	response.AppendController(e)
}

type Service struct {
	curd.DefaultController
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
		e.Err(http.StatusUnprocessableEntity, err)
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
	list := make([]form.ServiceListItem, 0)
	err = result.All(c, &list)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
	e.PageOK(list, count, req.Page, req.PageSize)
	return
}

// Generate 生成
func (e Service) Generate(c *gin.Context) {
	req := &form.ServiceGetReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusUnprocessableEntity, err)
		return
	}
	m := &models.Service{}
	err = mongodb.DB.Collection(e.TableName).FindOne(c, bson.M{"_id": objID}).Decode(m)
	if err != nil {
		e.Log.Error(err)
		e.Err(http.StatusInternalServerError, err)
		return
	}
}

func (e Service) Other(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	//r.POST("/generate/service/:id", e.Generate)
}
