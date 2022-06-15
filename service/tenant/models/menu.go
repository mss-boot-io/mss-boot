/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:24
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:24
 */

package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID          string      `bson:"_id" json:"id"`
	TenantID    string      `bson:"tenantID" json:"tenantID"`
	Name        string      `bson:"name" json:"name"`
	Icon        string      `bson:"icon" json:"icon"`
	Path        string      `bson:"path" json:"path"`
	Locale      string      `bson:"locale" json:"locale"`
	Access      string      `bson:"access" json:"access"`
	Description string      `bson:"description" json:"description"`
	HideInMenu  bool        `bson:"hideInMenu" json:"hideInMenu"`
	Status      enum.Status `bson:"status" json:"status"`
	Routes      []Menu      `bson:"routes" json:"routes"`
	ParentKeys  []string    `bson:"parentKeys" json:"parentKeys"`
	Redirect    string      `bson:"redirect" json:"redirect"`
	Layout      bool        `bson:"layout" json:"layout"`
	Component   string      `bson:"component" json:"component"`
	CreatedAt   time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt" bson:"updatedAt"`
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
	//初始化管理员租户数据
	tenant := &Tenant{}
	err = tenant.C().FindOne(context.TODO(), bson.M{
		"system": true,
		"status": enum.Enabled,
	}).Decode(tenant)
	if err != nil {
		log.Fatal(err)
	}
	count, err := e.C().CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		now := time.Now()
		e.C().InsertMany(context.TODO(), []interface{}{
			Menu{
				ID:          primitive.NewObjectID().String(),
				TenantID:    tenant.ID,
				Name:        "菜单",
				Locale:      "menu.menu",
				Description: "菜单管理",
				Path:        "/dynamic/menu",
				Layout:      true,
				CreatedAt:   now,
				UpdatedAt:   now,
				Routes: []Menu{
					{Path: "/dynamic/menu", Redirect: "/dynamic/menu"},
					{Name: "列表", Locale: "menu.list", Path: "/dynamic/menu/list", Description: "菜单列表"},
					{Name: "新建", Locale: "menu.create", Path: "/dynamic/menu/control/0"},
					{Name: "详情", HideInMenu: true, Path: "/dynamic/menu/detail/:id"},
					{Name: "修改", Path: "/dynamic/menu/control/:id", HideInMenu: true},
				},
			},
			Menu{
				ID:   primitive.NewObjectID().String(),
				Path: "/user",
				Routes: []Menu{
					{Name: "user", Routes: []Menu{
						{Name: "login", Path: "/user/login"},
					}},
				},
			},
		})
	}
}

func (e *Menu) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}
