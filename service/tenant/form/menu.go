/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:46
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:46
 */

package form

import (
	"time"

	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
)

type MenuCreateReq struct {

	//租户id
	TenantID string `bson:"tenantID" json:"tenantID" binding:"required"`

	//名称
	Name string `bson:"name" json:"name" binding:"required"`

	//icon
	Icon string `bson:"icon" json:"icon" `

	//路径
	Path string `bson:"path" json:"path" binding:"required"`

	//权限
	Access string `bson:"access" json:"access" binding:"required"`

	//状态
	Status enum.Status `bson:"status" json:"status" `

	//父菜单
	ParentKeys []string `bson:"parentKeys" json:"parentKeys" `

	//重定向
	Redirect string `bson:"redirect" json:"redirect" `

	//Layout
	Layout bool `bson:"layout" json:"layout" `

	//组件
	Component string `bson:"component" json:"component" `

	//创建时间
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

func (e *MenuCreateReq) SetCreatedAt() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
}

type MenuUpdateReq struct {
	curd.OneID

	//租户id
	TenantID string `bson:"tenantID" json:"tenantID" binding:"required"`

	//名称
	Name string `bson:"name" json:"name" binding:"required"`

	//icon
	Icon string `bson:"icon" json:"icon" `

	//路径
	Path string `bson:"path" json:"path" binding:"required"`

	//权限
	Access string `bson:"access" json:"access" binding:"required"`

	//状态
	Status enum.Status `bson:"status" json:"status" `

	//父菜单
	ParentKeys []string `bson:"parentKeys" json:"parentKeys" `

	//重定向
	Redirect string `bson:"redirect" json:"redirect" `

	//Layout
	Layout bool `bson:"layout" json:"layout" `

	//组件
	Component string `bson:"component" json:"component" `

	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

func (e *MenuUpdateReq) SetUpdatedAt() {
	e.UpdatedAt = time.Now()
}

type MenuGetReq struct {
	curd.OneID
}

type MenuGetResp struct {

	//id
	Id string `bson:"id" json:"id"`

	//租户id
	TenantID string `bson:"tenantID" json:"tenantID"`

	//名称
	Name string `bson:"name" json:"name"`

	//icon
	Icon string `bson:"icon" json:"icon"`

	//路径
	Path string `bson:"path" json:"path"`

	//权限
	Access string `bson:"access" json:"access"`

	//状态
	Status enum.Status `bson:"status" json:"status"`

	//<no value>
	ParentKeys []string `bson:"parentKeys" json:"parentKeys"`

	//重定向
	Redirect string `bson:"redirect" json:"redirect"`

	//Layout
	Layout bool `bson:"layout" json:"layout"`

	//组件
	Component string `bson:"component" json:"component"`

	//创建时间
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type MenuDeleteReq struct {
	curd.OneID
}

type MenuListReq struct {
	curd.Pagination

	//租户id
	TenantID string `bson:"tenantID" form:"tenantID" query:"tenantID" search:"type:exact;column:tenantID"`

	//名称
	Name string `bson:"name" form:"name" query:"name" search:"type:contains;column:name"`

	//路径
	Path string `bson:"path" form:"path" query:"path" search:"type:contains;column:path"`

	//权限
	Access string `bson:"access" form:"access" query:"access" search:"type:contains;column:access"`

	//状态
	Status enum.Status `bson:"status" form:"status" query:"status" search:"type:contains;column:status"`

	//<no value>
	ParentKeys []string `bson:"parentKeys" form:"parentKeys" query:"parentKeys" search:"type:<no value>;column:parentKeys"`

	//重定向
	Redirect string `bson:"redirect" form:"redirect" query:"redirect" search:"type:contains;column:redirect"`

	//Layout
	Layout bool `bson:"layout" form:"layout" query:"layout" search:"type:<no value>;column:layout"`

	//组件
	Component string `bson:"component" form:"component" query:"component" search:"type:contains;column:component"`
}

type MenuListItem struct {

	//id
	Id string `bson:"id" json:"id"`

	//租户id
	TenantID string `bson:"tenantID" json:"tenantID"`

	//名称
	Name string `bson:"name" json:"name"`

	//icon
	Icon string `bson:"icon" json:"icon"`

	//路径
	Path string `bson:"path" json:"path"`

	//权限
	Access string `bson:"access" json:"access"`

	//状态
	Status enum.Status `bson:"status" json:"status"`

	//<no value>
	ParentKeys []string `bson:"parentKeys" json:"parentKeys"`

	//重定向
	Redirect string `bson:"redirect" json:"redirect"`

	//Layout
	Layout bool `bson:"layout" json:"layout"`

	//组件
	Component string `bson:"component" json:"component"`
}
