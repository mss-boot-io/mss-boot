/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:20
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:20
 */

package models

import (
	"context"
	"time"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func init() {
	mongodb.AppendTable(&Role{})
}

type Role struct {
	ID        string      `bson:"_id" json:"id"`
	TenantID  string      `bson:"tenantID" json:"tenantID"`
	Name      string      `bson:"name" json:"name"`
	Status    enum.Status `bson:"status" json:"status"`
	Metadata  interface{} `bson:"metadata" json:"metadata"`
	CreatedAt time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time   `bson:"updatedAt" json:"updatedAt"`
}

func (Role) TableName() string {
	return "role"
}

func (e *Role) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}

func (e *Role) Make() {
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{"name", bsonx.Int32(1)},
				{"tenantID", bsonx.Int32(1)},
			},
			Options: ops,
		})
	if err != nil {
		log.Fatal(err)
	}
}
