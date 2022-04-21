/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:24
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:24
 */

package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	mongodb.AppendTable(&Menu{})
}

// Menu <no value>
type Menu struct {
	Id string `bson:"id" json:"id"`

	TenantID string `bson:"tenantID" json:"tenantID"`

	Name string `bson:"name" json:"name"`

	Icon string `bson:"icon" json:"icon"`

	Path string `bson:"path" json:"path"`

	Access string `bson:"access" json:"access"`

	Status enum.Status `bson:"status" json:"status"`

	Routes []Menu `bson:"routes" json:"routes"`

	ParentKeys []string `bson:"parentKeys" json:"parentKeys"`

	Redirect string `bson:"redirect" json:"redirect"`

	Layout bool `bson:"layout" json:"layout"`

	Component string `bson:"component" json:"component"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (Menu) TableName() string {
	return "menu"
}

func (e *Menu) Make() {
	ops := options.Index()
	ops.SetName("path")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bsonx.Doc{
			{"tenantID", bsonx.Int32(1)},
			{"path", bsonx.Int32(1)},
		},
		Options: ops,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (e *Menu) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}
