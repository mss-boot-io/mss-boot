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
	. "github.com/mss-boot-io/mss-boot/pkg/config/gormdb/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var Enforcer casbin.IEnforcer

type Database struct {
	Driver          string             `yaml:"driver"`
	Source          string             `yaml:"source"`
	ConnMaxIdleTime int                `yaml:"connMaxIdleTime"`
	ConnMaxLifeTime int                `yaml:"connMaxLifeTime"`
	MaxIdleConns    int                `yaml:"maxIdleConns"`
	MaxOpenConns    int                `yaml:"maxOpenConns"`
	Registers       []DBResolverConfig `yaml:"registers"`
	CasbinModel     string             `yaml:"casbinModel"`
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
