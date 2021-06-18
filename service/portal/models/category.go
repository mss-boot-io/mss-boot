/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 9:22 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 9:22 上午
 */

package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID          string         `gorm:"primaryKey;size:64;default:uuid();comment:ID" json:"id"`
	Name        string         `gorm:"size:255;comment:名称" json:"name"`
	Description string         `gorm:"type:text;comment:描述" json:"description"`
	CreatedAt   int64          `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt   int64          `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:软删除"`
}

func (Category) TableName() string {
	return "category"
}
