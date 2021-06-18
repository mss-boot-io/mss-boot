/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 5:04 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 5:04 下午
 */

package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tenant 租户
type Tenant struct {
	ID          string         `gorm:"primaryKey;size:64;default:(UUID());comment:ID"`
	Name        string         `gorm:"size:255;comment:名称"`
	Status      uint8          `gorm:"index;comment:状态"`
	System      uint8          `gorm:"index;comment:管理端"`
	Contact     string         `gorm:"comment:联系方式"`
	Domains     string         `gorm:"size:255;comment:域名"`
	Description string         `gorm:"type:text;comment:描述"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:软删除"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Tenant) TableName() string {
	return "tenant"
}

func (e *Tenant) BeforeCreate(_ *gorm.DB) error {
	e.ID = strings.Join(strings.Split(uuid.New().String(), "-"), "")
	return nil
}

// SetDomains set domains
func (e *Tenant) SetDomains(domains ...string) {
	e.Domains = strings.Join(domains, ",")
}

// GetDomains get domains
func (e *Tenant) GetDomains() []string {
	return strings.Split(e.Domains, ",")
}
