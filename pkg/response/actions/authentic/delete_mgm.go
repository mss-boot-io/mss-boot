package authentic

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 17:13:59
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 17:13:59
 */

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mss-boot-io/mss-boot/pkg/response"
)

// Delete action
type Delete struct {
	Base
	Key string
}

// NewDeleteMgm new deleteMgm action
func NewDeleteMgm(m mgm.Model, key string) *Delete {
	return &Delete{
		Base: Base{ModelMgm: m},
		Key:  key,
	}
}

// String action name
func (*Delete) String() string {
	return "deleteMgm"
}

// Handler action handler
func (e *Delete) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ids := make([]string, 0)
		v := c.Param(e.Key)
		if v == "batch" {
			api := response.Make(c).Bind(&ids)
			if api.Error != nil {
				api.Err(http.StatusUnprocessableEntity)
				return
			}
		}
		if e.ModelMgm != nil {
			e.deleteMgm(c, v)
			return
		}
		if e.ModelGorm != nil {
			e.deleteGorm(c, v)
			return
		}
		response.Error(c,
			http.StatusNotImplemented,
			fmt.Errorf("not implemented"))
	}
}

// deleteMgm  batch and single deleteMgm
func (e *Delete) deleteMgm(c *gin.Context, ids ...string) {
	api := response.Make(c)
	if len(ids) == 0 {
		api.OK(nil)
		return
	}
	_, err := mgm.Coll(e.ModelMgm).DeleteMany(c, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			api.OK(nil)
			return
		}
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(nil)
}
