/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 9:38 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 9:38 上午
 */

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Controllers = append(Controllers, &Category{})
}

type Category struct {
}

func (Category) Handlers() []gin.HandlerFunc {
	return nil
}

func (Category) Path() string {
	return "/category"
}

// Create 创建
// @Summary 创建category
// @Description 获取JSON
// @Tags category
// @Accept  application/json
// @Product application/json
// @Param data body form.CategoryCreateReq true "data"
// @Success 200 {object} form.CategoryCreateResp
// @Router /category [post]
// @Security Bearer
func (e Category) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (e Category) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (e Category) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (e Category) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (e Category) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
