/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/15 14:57
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/15 14:57
 */

package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
)

// Menu 菜单
type Menu struct {
	ID        string         `bson:"_id" json:"id"`
	TenantID  string         `bson:"tenantID" json:"tenantID"`
	Name      string         `bson:"name" json:"name"`
	Icon      string         `bson:"icon" json:"icon"`
	Path      string         `bson:"path" json:"path"`
	Access    string         `bson:"access" json:"access"`
	Status    enum.Status    `bson:"status" json:"status"`
	Routes    []MenuChildren `bson:"routes" json:"routes"`
	Metadata  interface{}    `bson:"metadata" json:"metadata"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
}

// MenuChildren 菜单子节点
type MenuChildren struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	HideInMenu bool   `json:"hideInMenu"`
}

func (Menu) TableName() string {
	return "menu"
}

func (e *Menu) Make() {
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bsonx.Doc{
			{"tenantID", bsonx.Int32(1)},
			{"name", bsonx.Int32(1)},
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
