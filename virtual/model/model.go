package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/10 15:29:38
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/10 15:29:38
 */

type Model struct {
	Table          string          `json:"tableName" yaml:"tableName" binding:"required"`
	AutoCreateTime schema.TimeType `json:"autoCreateTime" yaml:"autoCreateTime"`
	AutoUpdateTime schema.TimeType `json:"autoUpdateTime" yaml:"autoUpdateTime"`
	HardDeleted    bool            `json:"hardDeleted" yaml:"hardDeleted"`
	Fields         []*Field        `json:"fields" yaml:"fields" binding:"required"`
}

// TableName get table name
func (m *Model) TableName() string {
	return m.Table
}

// PrimaryKeys get primary keys
func (m *Model) PrimaryKeys() []string {
	var keys []string
	for i := range m.Fields {
		if m.Fields[i].PrimaryKey {
			keys = append(keys, m.Fields[i].Name)
		}
	}
	return keys
}

func (m *Model) Init() {
	if m.AutoCreateTime == 0 {
		m.AutoCreateTime = schema.UnixSecond
	}
	if m.AutoUpdateTime == 0 {
		m.AutoUpdateTime = schema.UnixSecond
	}
	for i := range m.Fields {
		m.Fields[i].Init()
	}
}

func (m *Model) Default(data any) {
	for i := range m.Fields {
		df := m.Fields[i].DefaultValue
		if m.Fields[i].DefaultValueFN != nil {
			df = m.Fields[i].DefaultValueFN()
		}
		if df == "" {
			continue
		}
		reflect.ValueOf(data).Elem().FieldByName(m.Fields[i].GetName()).Set(reflect.ValueOf(df))
	}
}

// MakeModel make virtual model
func (m *Model) MakeModel() any {
	fieldTypes := make([]reflect.StructField, 0)
	for i := range m.Fields {
		fieldTypes = append(fieldTypes, m.Fields[i].MakeField())
	}
	return reflect.New(reflect.StructOf(fieldTypes)).Interface()
}

func (m *Model) MakeList() any {
	fieldTypes := make([]reflect.StructField, 0)
	for i := range m.Fields {
		fieldTypes = append(fieldTypes, m.Fields[i].MakeField())
	}
	return reflect.New(reflect.SliceOf(reflect.StructOf(fieldTypes))).Interface()
}

func (m *Model) TableScope(db *gorm.DB) *gorm.DB {
	return db.Table(m.TableName())
}

func (m *Model) URI(ctx *gin.Context) (f func(*gorm.DB) *gorm.DB) {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Table(m.TableName())
		for _, key := range m.PrimaryKeys() {
			db = db.Where(fmt.Sprintf("%s in (?)", key), strings.Split(ctx.Param(key), ","))
		}
		return db
	}
}

func (m *Model) Pagination(ctx *gin.Context, p PaginationImp) (f func(*gorm.DB) *gorm.DB) {
	err := ctx.ShouldBindQuery(p)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			_ = db.AddError(err)
			return db
		}
		offset := (p.GetCurrent() - 1) * p.GetPageSize()
		return db.Offset(offset).Limit(p.GetPageSize())
	}
}

func (m *Model) Search(ctx *gin.Context) (f func(*gorm.DB) *gorm.DB) {
	return func(db *gorm.DB) *gorm.DB {
		for i := range m.Fields {
			v, ok := ctx.GetQuery(m.Fields[i].JsonTag)
			if !ok {
				continue
			}
			switch m.Fields[i].Search {
			case "exact", "iexact":
				db = db.Where(fmt.Sprintf("`%s`.`%s` = ?", m.Table, m.Fields[i].Name), v)
			case "contains", "icontains":
				db = db.Where(fmt.Sprintf("`%s`.`%s` like ?", m.Table, m.Fields[i].Name), "%"+v+"%")
			case "gt":
				db = db.Where(fmt.Sprintf("`%s`.`%s` > ?", m.Table, m.Fields[i].Name), v)
			case "gte":
				db = db.Where(fmt.Sprintf("`%s`.`%s` >= ?", m.Table, m.Fields[i].Name), v)
			case "lt":
				db = db.Where(fmt.Sprintf("`%s`.`%s` < ?", m.Table, m.Fields[i].Name), v)
			case "lte":
				db = db.Where(fmt.Sprintf("`%s`.`%s` <= ?", m.Table, m.Fields[i].Name), v)
			case "startWith", "istartWith":
				db = db.Where(fmt.Sprintf("`%s`.`%s` like ?", m.Table, m.Fields[i].Name), v+"%")
			case "endWith", "iendWith":
				db = db.Where(fmt.Sprintf("`%s`.`%s` like ?", m.Table, m.Fields[i].Name), "%"+v)
			case "in":
				arr, ok := ctx.GetQueryArray(m.Fields[i].JsonTag)
				if !ok {
					continue
				}
				db = db.Where(fmt.Sprintf("`%s`.`%s` in (?)", m.Table, m.Fields[i].JsonTag), arr)
			case "isnull":
				db = db.Where(fmt.Sprintf("`%s`.`%s` isnull", m.Table, m.Fields[i].JsonTag))
			case "order":
				switch v {
				case "desc":
					db = db.Order(fmt.Sprintf("`%s`.`%s` desc", m.Table, m.Fields[i].JsonTag))
				case "asc":
					db = db.Order(fmt.Sprintf("`%s`.`%s` asc", m.Table, m.Fields[i].JsonTag))
				}
			}
		}
		return db
	}
}
