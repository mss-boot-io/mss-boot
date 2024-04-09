//go:build sqlite3
// +build sqlite3

package gormdb

import (
	dm "github.com/nfjBill/gorm-driver-dm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Opens = map[string]func(string) gorm.Dialector{
	"mysql":    mysql.Open,
	"postgres": postgres.Open,
	"sqlite3":  sqlite.Open,
	"dm":       dm.Open,
}
