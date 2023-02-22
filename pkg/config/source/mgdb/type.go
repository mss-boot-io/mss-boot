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
	"strings"

	"github.com/kamva/mgm/v3"
	"gopkg.in/yaml.v3"

	"github.com/mss-boot-io/mss-boot/pkg"
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
		data[i] = pkg.BuildMap(keys, c.Tags[i].Value)
	}
	switch c.Ext {
	case "yml", "yaml":
		return yaml.Marshal(pkg.MergeMapsDepth(data...))
	case "json":
		return json.Marshal(pkg.MergeMapsDepth(data...))
	default:
		return nil, fmt.Errorf("not support %s", c.Ext)
	}
}
