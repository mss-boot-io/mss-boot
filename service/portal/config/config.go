/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 3:00 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 3:00 下午
 */

package config

import (
	log "github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/core-go/server"
	"github.com/lwnmengjing/core-go/server/listener"
	"github.com/lwnmengjing/mss-boot/pkg/config"
)

type Config struct {
	Logger config.Logger `yaml:"logger"`
	Server config.Listen `yaml:"server"`
	Health   *config.Listen  `json:"health"`
	Metrics  *config.Listen  `yaml:"metrics"`
}

func (e *Config) Init(mgr server.Manager) {
	e.Logger.Init()
	mgr.Add(
		listener.New("portal", e.Server.Init()...),
		listener.NewHealthz(e.Health.Init()...),
		listener.NewMetrics(e.Metrics.Init()...),
	)
}

func (e *Config) OnChange() {
	e.Logger.Init()
	log.Info("!!! config change and reload")
}
