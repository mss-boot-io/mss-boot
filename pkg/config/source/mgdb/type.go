/*
 * @Author: lwnmengjing
 * @Date: 2023/1/13 04:26:39
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/13 04:26:39
 */

package mgdb

import (
	"encoding/json"
	"fmt"
	"github.com/kamva/mgm/v3"
	"strings"

	"gopkg.in/yaml.v3"
)

type SystemConfig struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Ext              string `bson:"ext" json:"ext"`
	Tags             []Tag  `bson:"tags" json:"tags"`
	Description      string `bson:"description" json:"description"`
	Metadata         any    `bson:"metadata" json:"metadata"`
}

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
		data[i] = buildMap(keys, c.Tags[i].Value)
	}
	switch c.Ext {
	case "yml", "yaml":
		return yaml.Marshal(mergeMapsDepth(data...))
	case "json":
		return json.Marshal(mergeMapsDepth(data...))
	default:
		return nil, fmt.Errorf("not support %s", c.Ext)
	}
}

func buildMap(keys []string, value string) map[string]any {
	data := make(map[string]any)
	if len(keys) > 1 {
		data[keys[0]] = buildMap(keys[1:], value)
	} else {
		return map[string]any{keys[0]: value}
	}
	return data
}

// mergeMapsDepth deep merge multi map
func mergeMapsDepth(ms ...map[string]any) map[string]any {
	data := make(map[string]any)
	for i := range ms {
		data = mergeMapDepth(data, ms[i])
	}
	return data
}

// mergeMapDepth deep merge map
func mergeMapDepth(m1, m2 map[string]any) map[string]any {
	for k := range m2 {
		if v, ok := m1[k]; ok {
			if m, ok := v.(map[string]any); ok {
				m1[k] = mergeMapDepth(m, m2[k].(map[string]any))
			} else {
				m1[k] = m2[k]
			}
		} else {
			m1[k] = m2[k]
		}
	}
	return m1
}

// mergeMap merge map
func mergeMap(m1, m2 map[string]any) map[string]any {
	for k := range m2 {
		m1[k] = m2[k]
	}
	return m1
}
