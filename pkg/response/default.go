/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/15 23:31
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/15 23:31
 */

package response

import (
	"errors"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/search/mgos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/middlewares"
)

type DefaultController struct {
	Api
	TableName string
	Auth      bool
	CreateReq interface{}
	UpdateReq interface{}
	DeleteReq interface{}
	GetReq    interface{}
	GetResp   interface{}
	ListReq   interface{}
	ListResp  interface{}
}

func (e *DefaultController) Path() string {
	return strings.ReplaceAll(strings.ToLower(e.TableName), "_", "-")
}

func (e *DefaultController) Handlers() []gin.HandlerFunc {
	ms := make([]gin.HandlerFunc, 0)
	if e.Auth {
		ms = append(ms, middlewares.AuthMiddleware())
	}
	return ms
}

func (e DefaultController) Create(c *gin.Context) {
	req := e.CreateReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	if _, err = mongodb.DB.Collection(e.TableName).InsertOne(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (e DefaultController) Update(c *gin.Context) {
	req := e.UpdateReq
	err := e.Make(c).Bind(req).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	if _, err = mongodb.DB.Collection(e.TableName).UpdateOne(c, bson.M{"_id": objID}, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (e DefaultController) Delete(c *gin.Context) {
	e.Make(c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	if _, err = mongodb.DB.Collection(e.TableName).DeleteOne(c, bson.M{"_id": objID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (e DefaultController) Get(c *gin.Context) {
	e.Make(c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	resp := e.GetResp
	if err = mongodb.DB.Collection(e.TableName).FindOne(c, bson.M{"_id": objID}).Decode(resp); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"success":      false,
				"errorCode":    http.StatusNotFound,
				"errorMessage": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (e DefaultController) List(c *gin.Context) {
	req := e.ListReq
	err := e.Make(c).Bind(&req).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success":      false,
			"errorCode":    http.StatusUnprocessableEntity,
			"errorMessage": err.Error(),
		})
		return
	}
	filter, sort := mgos.MakeCondition(req)
	ops := options.Find()
	//todo: pagination
	ops.SetLimit(10)
	ops.SetSort(sort)
	ops.SetSkip(0)
	var count int64
	count, err = mongodb.DB.Collection(e.TableName).CountDocuments(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	result, err := mongodb.DB.Collection(e.TableName).Find(c, filter, ops)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	defer result.Close(c)
	resp := e.ListResp
	err = result.All(c, &resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}
	e.PageOK(resp, count, 1, 10)
}
