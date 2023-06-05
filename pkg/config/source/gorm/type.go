package gorm

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/11 01:22:20
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/11 01:22:20
 */

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

// SystemConfig system config
type SystemConfig struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Ext         string `gorm:"column:ext" json:"ext"`
	Description string `gorm:"column:description" json:"description"`
	Tags        []Tag  `gorm:"-" json:"tags"`
	Metadata    []byte `gorm:"column:metadata" json:"metadata"`
}

// Tag tag
type Tag struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	DataType string `json:"dataType"`
}

// GenerateBytes generate bytes
func (c *SystemConfig) GenerateBytes() ([]byte, error) {
	data := make([]map[string]interface{}, len(c.Tags))
	for i := range c.Tags {
		keys := strings.Split(c.Tags[i].Name, ".")
		data[i] = pkg.BuildMap(keys, c.Tags[i].Value)
	}
	switch c.Ext {
	case source.SchemeYaml.String(), source.SchemeYml.String():
		return yaml.Marshal(pkg.MergeMapsDepth(data...))
	case source.SchemeJSOM.String():
		return json.Marshal(pkg.MergeMapsDepth(data...))
	default:
		return nil, fmt.Errorf("not support %s", c.Ext)
	}
}
