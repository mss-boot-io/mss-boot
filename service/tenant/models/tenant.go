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

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func init() {
	mongodb.AppendTable(&Tenant{})
}

// Tenant 租户
type Tenant struct {
	ID          string      `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Contact     string      `json:"contact" bson:"contact"`
	System      bool        `json:"system" bson:"system"`
	Status      enum.Status `json:"status" bson:"status"`
	Description string      `json:"description" bson:"description"`
	Domains     []string    `json:"domains" bson:"domains"`
	Metadata    interface{} `json:"metadata" bson:"metadata"`
	ExpiredAt   time.Time   `json:"expiredAt" bson:"expiredAt" binding:"required"`
	CreatedAt   time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt" bson:"updatedAt"`
}

func (Tenant) TableName() string {
	return "tenant"
}

func (e *Tenant) Make() {
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}},
		Options: ops,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (e *Tenant) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}
