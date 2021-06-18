/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:26 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:26 下午
 */

package config

import (
	log "github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/core-go/server"
	"github.com/lwnmengjing/core-go/server/grpc"
	"github.com/lwnmengjing/core-go/server/listener"
	"github.com/lwnmengjing/mss-boot/pkg/config"

	"tenant/models"
)

// Config store config
type Config struct {
	Logger   config.Logger   `json:"logger"`
	Server   config.GRPC     `json:"grpc"`
	Database config.Database `json:"database"`
	Health   *config.Listen  `json:"health"`
	Metrics  *config.Listen  `json:"metrics"`
}

func (e *Config) Init(
	mgr server.Manager,
	register func(srv *grpc.Server)) {
	e.Logger.Init()
	var err error
	models.Orm, err = e.Database.Init()
	if err != nil {
		log.Errorf("config database init failed, %s", err.Error())
		panic(err)
	}
	mgr.Add(
		e.Server.Init(register),
		listener.NewHealthz(e.Health.Init()...),
		listener.NewMetrics(e.Metrics.Init()...),
	)
}

// OnChange receive config file changed event
func (e *Config) OnChange() {
	log.Info("!!! config change and reload")
}
