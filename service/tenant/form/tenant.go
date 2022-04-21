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

func (e *TenantCreateReq) SetCreatedAt() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
}

type TenantUpdateReq struct {
	curd.OneID
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

func (e *TenantUpdateReq) SetUpdatedAt() {
	e.UpdatedAt = time.Now()
}

type TenantGetReq struct {
	curd.OneID
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
	curd.OneID
}

type TenantListReq struct {
	curd.Pagination
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

type TenantAuthURLReq struct {
	ExtraScopes []string `query:"extraScopes" form:"extraScopes"`
	CrossClient []string `query:"crossClient" form:"crossClient"`
	ConnectorID string   `query:"connectorID" form:"connectorID"`
	State       string   `query:"state" form:"state"`
}

type TenantAuthURLResp struct {
	URL string `json:"url"`
}

type TenantClientResp struct {
	ServerURL        string `json:"serverUrl"`
	ClientID         string `json:"clientId"`
	AppName          string `json:"appName"`
	OrganizationName string `json:"organizationName"`
	AuthCodeURL      string `json:"authCodeURL"`
}

type TenantCallbackReq struct {
	Code             string `query:"code" form:"code"`
	RefreshToken     string `query:"refreshToken" form:"refreshToken"`
	State            string `query:"state" form:"state"`
	Error            string `query:"error" form:"error"`
	ErrorDescription string `query:"error_description" form:"error_description"`
}

type TenantCallbackResp struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"accessToken"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"tokenType,omitempty"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refreshToken,omitempty"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry time.Time `json:"expiry,omitempty"`
}

type TenantRefreshTokenReq struct {
	RefreshToken string `query:"refreshToken" form:"refreshToken"`
}
