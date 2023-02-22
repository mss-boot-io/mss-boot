/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 17:02:39
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 17:02:39
 */

package gormdb

import (
	log "github.com/mss-boot-io/mss-boot/core/logger"
	. "github.com/mss-boot-io/mss-boot/pkg/config/gormdb/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Database struct {
	Driver          string             `yaml:"driver"`
	Source          string             `yaml:"source"`
	ConnMaxIdleTime int                `yaml:"connMaxIdleTime"`
	ConnMaxLifeTime int                `yaml:"connMaxLifeTime"`
	MaxIdleConns    int                `yaml:"maxIdleConns"`
	MaxOpenConns    int                `yaml:"maxOpenConns"`
	Registers       []DBResolverConfig `yaml:"registers"`
}

type DBResolverConfig struct {
	Sources  []string `yaml:"sources"`
	Replicas []string `yaml:"replicas"`
	Policy   string   `yaml:"policy"`
	Tables   []string `yaml:"tables"`
}

func (e *Database) Init() {
	var err error
	registers := make([]ResolverConfigure, len(e.Registers))
	for i := range e.Registers {
		registers[i] = NewResolverConfigure(
			e.Registers[i].Sources,
			e.Registers[i].Replicas,
			e.Registers[i].Policy,
			e.Registers[i].Tables)
	}
	resolverConfig := NewConfigure(
		e.Source,
		e.MaxIdleConns,
		e.MaxOpenConns,
		e.ConnMaxIdleTime,
		e.ConnMaxLifeTime,
		registers)
	DB, err = resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				//SlowThreshold: time.Second,
				Colorful: true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.ToGorm()),
			},
		),
	}, opens[e.Driver])
	if err != nil {
		log.Fatalf("%s connect error : %s", e.Driver, err.Error())
	}
}
