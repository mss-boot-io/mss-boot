/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 9:33 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 9:33 上午
 */

package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/portal?charset=utf8mb4&parseTime=True&loc=Local&timeout=10000ms"))
	if err != nil {
		panic(err)
	}
	err = DB.Migrator().AutoMigrate(&Category{})
	if err != nil {
		panic(err)
	}
}
