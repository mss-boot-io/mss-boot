/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

package config

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/spf13/viper"
)

// Init 初始化配置
func Init(filename string, cfg Entity) error {
	var err error
	viper.SetConfigFile(filename)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return err
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(cfg)
		if err != nil {
			log.Fatal(err)
		}
		cfg.OnChange()
	})
	viper.WatchConfig()
	return err
}
