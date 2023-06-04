/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 17:02:39
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 17:02:39
 */

package gormdb

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	gormLogger "github.com/mss-boot-io/mss-boot/pkg/config/gormdb/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB gorm db
var DB *gorm.DB

// Enforcer casbin enforcer
var Enforcer casbin.IEnforcer

// Database database config
type Database struct {
	// Driver db type
	Driver string `yaml:"driver"`
	// Source db source
	Source string `yaml:"source"`
	// ConnMaxIdleTime conn max idle time
	ConnMaxIdleTime int `yaml:"connMaxIdleTime"`
	// ConnMaxLifeTime conn max lifetime
	ConnMaxLifeTime int `yaml:"connMaxLifeTime"`
	// MaxIdleConns max idle conns
	MaxIdleConns int `yaml:"maxIdleConns"`
	// MaxOpenConns max open conns
	MaxOpenConns int `yaml:"maxOpenConns"`
	// Registers db registers
	Registers []DBResolverConfig `yaml:"registers"`
	// CasbinModel casbin model
	CasbinModel string `yaml:"casbinModel"`
}

// DBResolverConfig db resolver config
type DBResolverConfig struct {
	// Sources db sources
	Sources []string `yaml:"sources"`
	// Replicas db replicas
	Replicas []string `yaml:"replicas"`
	// Policy db policy
	Policy string `yaml:"policy"`
	// Tables db tables
	Tables []string `yaml:"tables"`
}

// Init init db
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
		Logger: gormLogger.New(
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
	// casbin
	if e.CasbinModel != "" {
		//set casbin adapter
		var a persist.Adapter
		a, err = gormadapter.NewAdapterByDBWithCustomTable(DB, &CasbinRule{})
		if err != nil {
			log.Fatalf("gormadapter.NewAdapterByDB error : %s", err.Error())
		}
		var m model.Model
		m, err = model.NewModelFromString(e.CasbinModel)
		if err != nil {
			log.Fatalf("model.NewModelFromString error : %s", err.Error())
		}
		Enforcer, err = casbin.NewEnforcer(m, a)
		if err != nil {
			log.Fatalf("casbin.NewEnforcer error : %s", err.Error())
		}
		err = Enforcer.LoadPolicy()
		if err != nil {
			log.Fatalf("Enforcer.LoadPolicy error : %s", err.Error())
		}
		//Enforcer.EnableAutoSave(true)
		//Enforcer.EnableAutoBuildRoleLinks(true)
		//Enforcer.EnableLog(true)
	}
}
