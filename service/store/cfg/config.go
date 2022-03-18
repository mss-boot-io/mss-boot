package cfg

import (
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
	"github.com/mss-boot-io/mss-boot/core/server/listener"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
)

type Config struct {
	Logger   config.Logger    `yaml:"logger" json:"logger"`
	Server   config.GRPC      `yaml:"server" json:"server"`
	Health   *config.Listen   `yaml:"health" json:"health"`
	Metrics  *config.Listen   `yaml:"metrics" json:"metrics"`
	Clients  config.Clients   `yaml:"clients" json:"clients"`
	Database mongodb.Database `yaml:"database" json:"database"`
}

func (e *Config) Init(handler func(srv *grpc.Server)) {
	e.Logger.Init()
	e.Database.Init()

	runnable := []server.Runnable{
		e.Server.Init(handler, grpc.WithID("store")),
	}
	if e.Health != nil {
		runnable = append(runnable, listener.NewHealthz(e.Health.Init()...))
	}
	if e.Metrics != nil {
		runnable = append(runnable, listener.NewMetrics(e.Metrics.Init()...))
	}

	server.Manage.Add(runnable...)
}

func (e *Config) OnChange() {
	e.Logger.Init()
	e.Database.Init()
	log.Info("!!! cfg change and reload")
}
