package actions

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/26 00:42:12
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/26 00:42:12
 */

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// Get action
type Get struct {
	Base
	Key string
}

// NewGetMgm new get action
func NewGetMgm(m mgm.Model, key string) *Get {
	return &Get{
		Base: Base{ModelMgm: m},
		Key:  key,
	}
}

// String action name
func (*Get) String() string {
	return "get"
}

// Handler action handler
func (e *Get) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if e.ModelMgm != nil {
			e.getMgm(c, e.Key)
			return
		}
		if e.ModelMgm != nil {
			e.getMgm(c, e.Key)
			return
		}
		response.Error(c,
			http.StatusNotImplemented,
			fmt.Errorf("not implemented"))
	}
}

func (e *Get) getMgm(c *gin.Context, key string) {
	api := response.Make(c)
	m := pkg.ModelDeepCopy(e.ModelMgm)
	id, err := primitive.ObjectIDFromHex(c.Param(key))
	if err != nil {
		api.AddError(err)
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err = mgm.Coll(e.ModelMgm).
		FindOne(c, bson.M{"_id": id}).
		Decode(m)
	if err != nil {
		api.AddError(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			api.Err(http.StatusNotFound)
			return
		}
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
