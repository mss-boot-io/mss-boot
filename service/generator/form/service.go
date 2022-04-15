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

type ServiceCreateReq struct {
	Name        string      `bson:"name" json:"name"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"-"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"-"`
}

type ServiceUpdateReq struct {
	curd.OneID
	Name        string      `json:"name" bson:"name"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
}

func (e *ServiceUpdateReq) GetData() any {
	return *e
}

type ServiceGetReq struct {
	curd.OneID
}

type ServiceGetResp struct {
	ID          string      `bson:"_id" json:"id"`
	Name        string      `bson:"name" json:"name"`
	Status      enum.Status `bson:"status" json:"status"`
	Metadata    interface{} `bson:"metadata" json:"metadata"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
}

type ServiceDeleteReq struct {
	curd.OneID
}

type ServiceListReq struct {
	curd.Pagination
	Name string `query:"name" form:"name" search:"type:contains;column:name"`
}

type ServiceListItem struct {
	ID          string      `bson:"_id" json:"id"`
	Name        string      `bson:"name" json:"name"`
	Status      enum.Status `bson:"status" json:"status"`
	Description string      `bson:"description" json:"description"`
	CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
}
