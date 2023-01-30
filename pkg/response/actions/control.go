/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 17:13:38
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 17:13:38
 */

package actions

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// Control action
type Control struct {
	Base
	Key string
}

// NewControl new control action
func NewControl(m mgm.Model, key string) *Control {
	return &Control{
		Base: Base{Model: m},
		Key:  key,
	}
}

// String action name
func (*Control) String() string {
	return "control"
}

// Handler action handler
func (e *Control) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			//create
			e.create(c)
		case http.MethodPut:
			//update
			e.update(c)
		default:
			response.Error(c,
				http.StatusMethodNotAllowed,
				fmt.Errorf("method %s not support", c.Request.Method))
		}
	}
}

func (e *Control) create(c *gin.Context) {
	m := pkg.ModelDeepCopy(e.Model)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err := mgm.Coll(e.Model).CreateWithCtx(c, m)
	if err != nil {
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}

func (e *Control) update(c *gin.Context) {
	m := pkg.ModelDeepCopy(e.Model)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	id, err := primitive.ObjectIDFromHex(c.Param(e.Key))
	if err != nil {
		api.AddError(err)
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	m.SetID(id)
	err = mgm.Coll(e.Model).UpdateWithCtx(c, m)
	if err != nil {
		api.AddError(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			api.Err(http.StatusNotFound)
			return
		}
		api.Log.Error(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
