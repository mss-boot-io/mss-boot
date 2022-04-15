/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 15:33
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 15:33
 */

package form

import (
	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"time"
)

type ModelCreateReq struct {
	Service     string      `bson:"service" json:"service"`
	Name        string      `bson:"name" json:"name"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"-"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"-"`
}

type ModelUpdateReq struct {
	curd.OneID
	Name        string      `json:"name" bson:"name"`
	Service     string      `bson:"service" json:"service"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
}

type ModelGetReq struct {
	curd.OneID
}

type ModelGetResp struct {
	ID          string      `bson:"_id" json:"id"`
	Service     string      `bson:"service" json:"service"`
	Name        string      `bson:"name" json:"name"`
	Status      enum.Status `bson:"status" json:"status"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
}

type ModelDeleteReq struct {
	curd.OneID
}

type ModelListReq struct {
	curd.Pagination
	Name    string `query:"name" form:"name" search:"type:contains;column:name"`
	Service string `query:"service" form:"service" search:"type:contains;column:service"`
}

type ModelListItem struct {
	ID          string      `bson:"_id" json:"id"`
	Service     string      `bson:"service" json:"service"`
	Name        string      `bson:"name" json:"name"`
	Status      enum.Status `bson:"status" json:"status"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
}
