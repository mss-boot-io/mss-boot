/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 15:29
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 15:29
 */

package controllers

import (
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"github.com/mss-boot-io/mss-boot/service/generator/form"
)

func init() {
	e := &Model{}
	e.TableName = "model"
	e.Auth = true
	e.CreateReq = &form.ModelCreateReq{}
	e.UpdateReq = &form.ModelUpdateReq{}
	e.DeleteReq = &form.ModelDeleteReq{}
	e.GetReq = &form.ModelGetReq{}
	e.GetResp = &form.ModelGetResp{}
	e.ListReq = &form.ModelListReq{}
	e.ListResp = &form.ModelGetResp{}
	response.AppendController(e)
}

type Model struct {
	curd.DefaultController
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
