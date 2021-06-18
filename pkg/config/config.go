/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

package config

import (
	"github.com/lwnmengjing/core-go/config"
	"github.com/lwnmengjing/core-go/config/source/file"
)

// Init 初始化配置
func Init(filename string, cfg config.Entity) error {
	var err error
	config.DefaultConfig, err = config.NewConfig(
		config.WithEntity(cfg),
		config.WithSource(
			file.NewSource(
				file.WithPath(filename))))
	return err
}
