/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:37
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:37
 */

package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
)

type Tenant struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Contact     string             `bson:"contact"`
	Description string             `bson:"description"`
	Domains     []string           `bson:"domains"`
	Status      uint8              `bson:"status"`
	Expire      time.Time          `bson:"expire"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

func (e *Tenant) TableName() string {
	return "tenant"
}

func (e *Tenant) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}

//
//func (e *Tenant) Create() error {
//	e.ID = bson.NewObjectId().Hex()
//	e.CreatedAt = time.Now()
//	e.UpdatedAt = time.Now()
//	return e.C().InsertOne(context.TODO(), e)
//}

//func (e *Tenant) Update() error {
//	return e.C().UpdateId(e.ID, e)
//}

func (e *Tenant) Exist(ctx context.Context) (bool, error) {
	sr := e.C().FindOne(ctx, bson.M{"name": e.Name})
	if sr.Err() != nil {
		if errors.Is(mongo.ErrNoDocuments, sr.Err()) {
			return false, nil
		}
		return false, sr.Err()
	}
	_, err := sr.DecodeBytes()
	if err != nil {
		return false, nil
	}
	return true, nil
}
