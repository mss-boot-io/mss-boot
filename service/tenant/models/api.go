/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/5/23 0:22
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/5/23 0:22
 */

package models

import (
	"context"
	"flag"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
)

var Routes []gin.RouteInfo

func init() {
	mongodb.AppendTable(&API{})
}

type API struct {
	ID          string    `bson:"_id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Method      string    `bson:"method" json:"method"`
	Path        string    `bson:"path" json:"path"`
	Handler     string    `bson:"handler" json:"handler"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}

func (API) TableName() string {
	return "routes"
}

func (e *API) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}

func (e *API) Make() {
	ops := options.Index()
	ops.SetName("path")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{"path", bsonx.Int32(1)},
				{"method", bsonx.Int32(1)},
			},
			Options: ops,
		})
	if err != nil {
		log.Fatal(err)
	}
	if cast.ToBool(flag.Lookup("api-refresh").Value.String()) {
		err = e.BatchAddPath(Routes...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (e *API) BatchAddPath(routes ...gin.RouteInfo) error {
	if len(routes) == 0 {
		return nil
	}
	apiList := make([]interface{}, len(routes))
	now := time.Now()
	for i := range routes {
		apiList[i] = API{
			ID:        primitive.NewObjectID().String(),
			Method:    routes[i].Method,
			Path:      routes[i].Path,
			Handler:   routes[i].Handler,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}
	//清空数据
	_, err := e.C().DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	_, err = e.C().InsertMany(context.TODO(), apiList)
	return err
}
