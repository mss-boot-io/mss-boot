/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:43
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:43
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/errors"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
	"github.com/sanity-io/litter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"tenant/form"
	"tenant/models"
	"time"
)

func init() {
	response.AppendController(&Tenant{})
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)

	mongodb.CreateIndex("tenant", mongo.IndexModel{
		Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}},
		Options: ops,
	})
}

type Tenant struct {
	response.Api
}

func (Tenant) C() *mongo.Collection {
	return mongodb.DB.Collection("tenant")
}

func (Tenant) Path() string {
	return "/tenant"
}

func (e Tenant) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Create 创建
// @Summary 创建tenant
// @Description 创建tenant
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Param data body form.TenantCreateReq true "data"
// @Success 200 {object} response.Response
// @Router /tenant [post]
// @Security Bearer
func (e Tenant) Create(c *gin.Context) {
	req := &form.TenantCreateReq{}

	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.TenantSvcParamsInvalid, err)
		return
	}

	req.CreatedAt = time.Now()
	req.UpdatedAt = req.CreatedAt
	if req.ExpiredAt.Sub(req.CreatedAt) < 0 {
		e.Err(errors.TenantSvcParamsInvalid,
			errors.New("expiredAt must be after now"))
		return
	}

	if _, err = e.C().InsertOne(c, req); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			e.Err(errors.TenantSvcRecordIsExist, err)
			return
		}
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}

	e.OK(nil)
	return
}

// Update 更新
// @Summary 更新tenant
// @Description 更新tenant
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Param id path string false "id"
// @Param data body form.TenantUpdateReq true "data"
// @Success 200 {object} response.Response
// @Router /tenant/{id} [put]
// @Security Bearer
func (e Tenant) Update(c *gin.Context) {
	req := &form.TenantUpdateReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.TenantSvcParamsInvalid, err)
		return
	}
	m := &models.Tenant{}
	pkg.Translate(req, m)
}

// Delete 删除
// @Summary 删除tenant
// @Description 删除tenant
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Param id path string false "id"
// @Success 200 {object} response.Response
// @Router /tenant/{id} [delete]
// @Security Bearer
func (e Tenant) Delete(c *gin.Context) {
	req := &form.TenantDeleteReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.TenantSvcParamsInvalid, err)
		return
	}

	println(req.ID)
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcObjectIDInvalid, err)
		return
	}
	_, err = e.C().DeleteOne(c, bson.M{"_id": objID})
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		e.Err(errors.TenantSvcRecordNotFound, nil)
		return
	}

	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}
	e.OK(nil)
	return
}

// Get 获取
// @Summary 获取tenant
// @Description 获取tenant
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Param id path string false "id"
// @Success 200 {object} response.Response
// @Router /tenant/{id} [get]
// @Security Bearer
func (e Tenant) Get(c *gin.Context) {
	req := &form.TenantGetReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.TenantSvcParamsInvalid, err)
		return
	}
	resp := &form.TenantGetResp{}
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcObjectIDInvalid, err)
		return
	}
	err = e.C().FindOne(c, bson.M{"_id": objID}).Decode(&resp)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}
	e.OK(resp)
	return
}

// List 列表
// @Summary 列表tenant
// @Description 列表tenant
// @Tags tenant
// @Accept  application/json
// @Product application/json
// @Param name query string false "租户名称"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页容量"
// @Success 200 {object} response.Page{list=[]form.TenantListItem}
// @Router /tenant [get]
// @Security Bearer
func (e Tenant) List(c *gin.Context) {
	req := &form.TenantListReq{}
	err := e.Make(c).Bind(req).Error
	if err != nil {
		e.Err(errors.TenantSvcParamsInvalid, err)
		return
	}
	litter.Dump(req)
	filter := mgos.MakeCondition(*req)

	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	ops := options.Find()
	ops.SetLimit(int64(req.PageSize))
	ops.SetSkip(int64(req.PageSize * (req.Page - 1)))
	var count int64
	count, err = e.C().CountDocuments(c, filter)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}
	result, err := e.C().Find(c, filter, ops)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}
	defer result.Close(c)
	list := make([]form.TenantListItem, 0)
	//for result.Next(c) {
	//	result.Decode()
	//}
	err = result.All(c, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Err(errors.TenantSvcOperateDBFailed, err)
		return
	}
	e.PageOK(list, count, req.Page, req.PageSize)
	return
}

func (e Tenant) Other(_ *gin.RouterGroup) {

}
