package models

import (
	"gorm.io/gorm"
	"time"
)

type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

func (*Migration) TableName() string {
	return "mss_boot_migration"
}

func (e *Migration) SetVersion(version string) {
	e.Version = version
}

func (e *Migration) Done(tx *gorm.DB) (bool, error) {
	var count int64
	err := tx.Model(e).Where("version = ?", e.Version).Count(&count).Error
	return count > 0, err
}
