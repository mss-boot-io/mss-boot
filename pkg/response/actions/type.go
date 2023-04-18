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

type ModelProvider string

const (
	ModelProviderMgm  ModelProvider = "mgm"
	ModelProviderGorm ModelProvider = "gorm"
)

type Model interface {
	mgm.Model
	schema.Tabler
}

type ModelGorm struct {
	ID        string         `gorm:"primarykey" json:"id" bson:"_id,omitempty" form:"id" query:"id"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" bson:"-" json:"-"`
}

func (e *ModelGorm) PrepareID(id any) (any, error) {
	if e.ID != "" {
		e.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return e.ID, nil
}

func (e *ModelGorm) GetID() any {
	return e.ID
}

func (e *ModelGorm) SetID(id any) {
	e.ID = cast.ToString(id)
}

func (e *ModelGorm) TableName() string {
	return "mss-boot-base"
}
