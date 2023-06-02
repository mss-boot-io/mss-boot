/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/4/19 01:00:40
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/4/19 01:00:40
 */

package actions

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ModelProvider model provider
type ModelProvider string

const (
	// ModelProviderMgm mgm model provider
	ModelProviderMgm ModelProvider = "mgm"
	// ModelProviderGorm gorm model provider
	ModelProviderGorm ModelProvider = "gorm"
)

// Model gorm and mgm model
type Model interface {
	mgm.Model
	schema.Tabler
}

// ModelGorm model gorm
type ModelGorm struct {
	// ID primary key
	ID string `gorm:"primarykey" json:"id" bson:"_id,omitempty" form:"id" query:"id"`
	// CreatedAt create time
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	// UpdatedAt update time
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	// DeletedAt delete time soft delete
	DeletedAt gorm.DeletedAt `gorm:"index" bson:"-" json:"-"`
}

// PrepareID prepare id
func (e *ModelGorm) PrepareID(_ any) (any, error) {
	if e.ID == "" {
		e.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return e.ID, nil
}

// GetID get id
func (e *ModelGorm) GetID() any {
	return e.ID
}

// SetID set id
func (e *ModelGorm) SetID(id any) {
	e.ID = cast.ToString(id)
}

// TableName return table name
func (e *ModelGorm) TableName() string {
	return "mss-boot-base"
}
