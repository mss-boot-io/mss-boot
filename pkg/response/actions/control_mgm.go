package actions

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 17:13:38
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 17:13:38
 */

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgm "github.com/kamva/mgm/v3"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Control action
type Control struct {
	Base
	Key string
}

// NewControlMgm new control action
func NewControlMgm(m mgm.Model, key string) *Control {
	return &Control{
		Base: Base{ModelMgm: m},
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
			if e.ModelGorm != nil {
				e.createGorm(c)
				break
			}
			if e.ModelMgm != nil {
				e.createMongo(c)
				break
			}
		case http.MethodPut:
			//update
			if e.ModelGorm != nil {
				e.updateGorm(c)
				break
			}
			if e.ModelMgm != nil {
				e.updateMongo(c)
				break
			}
		default:
			response.Error(c,
				http.StatusNotImplemented,
				fmt.Errorf("not implemented"))
		}
	}
}

func (e *Control) createMongo(c *gin.Context) {
	m := pkg.ModelDeepCopy(e.ModelMgm)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err := mgm.Coll(e.ModelMgm).CreateWithCtx(c, m)
	if err != nil {
		api.Log.Error(err)
		api.AddError(err)
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}

func (e *Control) updateMongo(c *gin.Context) {
	m := pkg.ModelDeepCopy(e.ModelMgm)
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
	err = mgm.Coll(e.ModelMgm).UpdateWithCtx(c, m)
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
