/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 3:00 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 3:00 下午
 */

package cfg

import (
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/listener"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"net/http"
	"oauth2/common"
)

type Config struct {
	Logger   config.Logger  `yaml:"logger"`
	Server   config.Listen  `yaml:"server"`
	Health   *config.Listen `yaml:"health"`
	Metrics  *config.Listen `yaml:"metrics"`
	Clients  []Client       `yaml:"clients"`
	Database Database       `yaml:"database"`
}

type Database struct {
	URL     string `yaml:"url"`
	Name    string `yaml:"name"`
	Timeout int    `yaml:"timeout"`
}

type Client struct {
	ID       string `yaml:"id" bson:"ID"`
	Secret   string `yaml:"secret" bson:"Secret"`
	Domain   string `yaml:"domain" bson:"Domain"`
	Redirect string `yaml:"redirect" bson:"Redirect"`
}

func (e *Config) Init(mgr server.Manager, handler http.Handler) error {
	e.Logger.Init()

	err := common.MakeDB(e.Database.URL,
		e.Database.Name,
		e.Database.Timeout)
	if err != nil {
		return err
	}
	mgr.Add(
		listener.New("oauth2", e.Server.Init(listener.WithHandler(handler))...),
		listener.NewHealthz(e.Health.Init()...),
		listener.NewMetrics(e.Metrics.Init()...),
	)
	return nil
}

func (e *Config) OnChange() {
	e.Logger.Init()
	log.Info("!!! cfg change and reload")
}
