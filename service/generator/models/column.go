/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 15:15
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 15:15
 */

package models

import (
	"context"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Column struct {
	ID      string `bson:"id" json:"id"`
	ModelID string `bson:"modelID" json:"modelID"`
	Name    string `bson:"name" json:"name"`
	Key     string `bson:"key" json:"key"`
	Type    string `bson:"type" json:"type"`
	Comment string `bson:"comment" json:"comment"`
}

func (Column) TableName() string {
	return "column"
}

func (e *Column) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}

func (e *Column) Make() {
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{"name", bsonx.Int32(1)},
				{"modelID", bsonx.Int32(1)},
			},
			Options: ops,
		})
	if err != nil {
		log.Fatal(err)
	}
}
