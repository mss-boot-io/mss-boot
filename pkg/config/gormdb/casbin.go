/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 00:24:48
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 00:24:48
 */

package gormdb

// CasbinRule casbin rule
type CasbinRule struct {
	PType string `json:"pType" gorm:"size:100;"`
	V0    string `json:"v0" gorm:"size:100;"`
	V1    string `json:"v1" gorm:"size:100;"`
	V2    string `json:"v2" gorm:"size:100;"`
	V3    string `json:"v3" gorm:"size:100;"`
	V4    string `json:"v4" gorm:"size:100;"`
	V5    string `json:"v5" gorm:"size:100;"`
}

// TableName table name
func (*CasbinRule) TableName() string {
	return "casbin_rule"
}
