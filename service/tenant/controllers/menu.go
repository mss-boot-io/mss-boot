/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:43
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:43
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"github.com/mss-boot-io/mss-boot/service/tenant/form"
)

func init() {
	e := new(Menu)
	e.TableName = "menu"
	e.Auth = true
	e.CreateReq = &form.MenuCreateReq{}
	e.UpdateReq = &form.MenuUpdateReq{}
	e.GetReq = &form.MenuGetReq{}
	e.GetResp = &form.MenuGetResp{}
	e.DeleteReq = &form.MenuDeleteReq{}
	e.ListReq = &form.MenuListReq{}
	e.ListResp = make([]form.MenuListItem, 0)
	response.AppendController(e)
}

type Menu struct {
	curd.DefaultController
}

// Create 创建
// @Summary 创建menu
// @Description 创建menu
// @Tags menu
// @Accept  application/json
// @Product application/json
// @Param data body form.menuCreateReq true "data"
// @Success 200 {object} response.Response
// @Router /tenant/api/v1/menu [post]
// @Security Bearer

// Update 更新
// @Summary 更新menu
// @Description 更新menu
// @Tags menu
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body form.menuUpdateReq true "data"
// @Success 200 {object} response.Response
// @Router /tenant/api/v1/menu/{id} [put]
// @Security Bearer

// Delete 删除
// @Summary 删除menu
// @Description 删除menu
// @Tags menu
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /tenant/api/v1/menu/{id} [delete]
// @Security Bearer

// Get 获取
// @Summary 获取menu
// @Description 获取menu
// @Tags menu
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response{data=form.menuGetResp}
// @Router /tenant/api/v1/menu/{id} [get]
// @Security Bearer

// List 列表
// @Summary 列表menu
// @Description 列表menu
// @Tags menu
// @Accept  application/json
// @Product application/json
// @Param name query string false "租户名称"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页容量"
// @Success 200 {object} response.Page{data=[]form.menuListItem}
// @Router /tenant/api/v1/menu [get]
// @Security Bearer

func (e *Menu) Other(r *gin.RouterGroup) {
}
