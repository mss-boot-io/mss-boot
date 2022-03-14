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
)

type TenantCreateReq struct {
	//名称
	Name string `json:"name" bson:"name" binding:"required"`
	//邮箱
	Email string `json:"email" bson:"email" binding:"email"`
	//联系方式
	Contact string `json:"contact" bson:"contact"`
	//描述
	Description string `json:"description" bson:"description"`
	//域名
	Domains []string `json:"domains" bson:"domains"`
	//系统管理
	System bool `json:"system" bson:"system"`
	//状态
	Status enum.Status `json:"status" bson:"status"`
	//有效期
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt" binding:"required"`
	//创建时间
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

type TenantUpdateReq struct {
	//id
	ID string `uri:"id" json:"-" bson:"_id"`
	//名称
	Name string `json:"name" bson:"name"`
	//邮箱
	Email string `json:"email" bson:"email"`
	//联系方式
	Contact string `json:"contact" bson:"contact"`
	//描述
	Description string `json:"description" bson:"description"`
	//域名
	Domains []string `json:"domains" bson:"domains"`
	//有效期
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt" binding:"required"`
	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

type TenantGetReq struct {
	ID string `uri:"id" json:"-" bson:"_id"`
}

type TenantGetResp struct {
	//id
	ID string `uri:"id" json:"-" bson:"_id"`
	//名称
	Name string `json:"name" bson:"name"`
	//联系方式
	Contact string `json:"contact" bson:"contact"`
	//系统管理
	System bool `json:"system" bson:"system"`
	//状态
	Status enum.Status `json:"status" bson:"status"`
	//描述
	Description string `json:"description" bson:"description"`
	//域名
	Domains []string `json:"domains" bson:"domains"`
	//有效期
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt" binding:"required"`
	//创建时间
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TenantDeleteReq struct {
	ID string `uri:"id" json:"-" bson:"_id"`
}

type TenantListReq struct {
	Pagination
	//名称
	Name string `query:"name" form:"name" search:"type:contains;column:name"`
	//联系方式
	Contact string `query:"contact" form:"contact" search:"type:contains;column:contact"`
	//系统管理
	System bool `query:"system" form:"system"`
	//状态
	Status enum.Status `query:"status" form:"status"`
	//系统管理排序
	SystemSort int8 `query:"systemSort" form:"systemSort" search:"type:order;column:system"`
}

type TenantListItem struct {
	//id
	ID string `json:"id" bson:"_id"`
	//名称
	Name string `json:"name" bson:"name"`
	//联系方式
	Contact string `json:"contact" bson:"contact"`
	//系统管理
	System bool `json:"system" bson:"system"`
	//状态
	Status enum.Status `json:"status" bson:"status"`
	//描述
	Description string `json:"description" bson:"description"`
	//域名
	Domains []string `json:"domains" bson:"domains"`
	//有效期
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt" binding:"required"`
	//创建时间
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
