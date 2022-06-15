/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/spf13/viper"
)

// Init 初始化配置
func Init(cfg Entity) error {
	var err error
	v := viper.New()
	v.SetConfigFile("cfg/application.yml")
	err = v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(cfg)
	if err != nil {
		return err
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		err := v.Unmarshal(cfg)
		if err != nil {
			log.Fatal(err)
		}
		cfg.OnChange()
	})
	v.WatchConfig()
	return err
}

// InitEmbed init config from embed file system, not support OnConfigChange, but support stage config merge
func InitEmbed(cfg Entity, fs embed.FS, stage string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	rb, err := fs.ReadFile("application.yml")
	if err != nil {
		return err
	}
	err = v.ReadConfig(bytes.NewReader(rb))
	if err != nil {
		return err
	}
	if stage != "" {
		rb, err = fs.ReadFile(fmt.Sprintf("application-%s.yml", stage))
		if err != nil {
			return err
		}
	}
	return v.Unmarshal(cfg)
}
