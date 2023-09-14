package model

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm/schema"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/10 16:11:55
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/10 16:11:55
 */

type Field struct {
	Name                   string          `json:"name" yaml:"name" binding:"required"`
	JsonTag                string          `json:"jsonTag" yaml:"jsonTag"`
	DataType               schema.DataType `json:"type" yaml:"type" binding:"required"`
	PrimaryKey             bool            `json:"primaryKey" yaml:"primaryKey"`
	AutoIncrement          bool            `json:"autoIncrement" yaml:"autoIncrement"`
	AutoIncrementIncrement int64           `json:"autoIncrementIncrement" yaml:"autoIncrementIncrement"`
	Creatable              bool            `json:"creatable" yaml:"creatable"`
	Updatable              bool            `json:"updatable" yaml:"updatable"`
	Readable               bool            `json:"readable" yaml:"readable"`
	DefaultValue           string          `json:"defaultValue" yaml:"defaultValue"`
	DefaultValueFN         func() string   `json:"-" yaml:"-"`
	NotNull                bool            `json:"notNull" yaml:"notNull"`
	Unique                 bool            `json:"unique" yaml:"unique"`
	Index                  string          `json:"index" yaml:"index"`
	Comment                string          `json:"comment" yaml:"comment"`
	Size                   int             `json:"size" yaml:"size"`
	Precision              int             `json:"precision" yaml:"precision"`
	Scale                  int             `json:"scale" yaml:"scale"`
	Search                 string          `json:"search" yaml:"search"`
}

type DefaultFN string

const (
	UUID DefaultFN = "uuid"
	Now  DefaultFN = "now"
)

var UUIDFN = func() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

var NowFN = func() string {
	return time.Now().String()
}

func (f *Field) Init() {
	if f.JsonTag == "" {
		f.JsonTag = f.Name
	}
	if f.PrimaryKey {
		f.DefaultValueFN = UUIDFN
	}
	if f.DataType == schema.Time && f.NotNull {
		f.DefaultValueFN = NowFN
	}
}

func (f *Field) GetName() string {
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

func (f *Field) MakeField() reflect.StructField {
	field := reflect.StructField{
		Name: f.GetName(),
		Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s" gorm:"column:%s"`, f.JsonTag, f.Name)),
	}
	switch f.DataType {
	case schema.Bool:
		field.Type = reflect.TypeOf(false)
	case schema.Float:
		field.Type = reflect.TypeOf(float64(0))
	case schema.Int:
		field.Type = reflect.TypeOf(int(0))
	case schema.Uint:
		field.Type = reflect.TypeOf(uint(0))
	case schema.Time:
		field.Type = reflect.TypeOf(time.Time{})
	default:
		field.Type = reflect.TypeOf("")
	}
	return field
}