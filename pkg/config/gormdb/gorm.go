package gormdb

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 17:02:39
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 17:02:39
 */

import (
	"log/slog"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/search/gorms"
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
	// Config gorm config
	Config GORMConfig `yaml:"config"`
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

type GORMConfig struct {
	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool `yaml:"skipDefaultTransaction" json:"skipDefaultTransaction"`
	// FullSaveAssociations full save associations
	FullSaveAssociations bool `yaml:"fullSaveAssociations" json:"fullSaveAssociations"`
	// DryRun generate sql without execute
	DryRun bool `yaml:"dryRun" json:"dryRun"`
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool `yaml:"prepareStmt" json:"prepareStmt"`
	// DisableAutomaticPing
	DisableAutomaticPing bool `yaml:"disableAutomaticPing" json:"disableAutomaticPing"`
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool `yaml:"disableForeignKeyConstraintWhenMigrating" json:"disableForeignKeyConstraintWhenMigrating"`
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool `yaml:"ignoreRelationshipsWhenMigrating" json:"ignoreRelationshipsWhenMigrating"`
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool `yaml:"disableNestedTransaction" json:"disableNestedTransaction"`
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool `yaml:"allowGlobalUpdate" json:"allowGlobalUpdate"`
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool `yaml:"queryFields" json:"queryFields"`
	// CreateBatchSize default create batch size
	CreateBatchSize int `yaml:"createBatchSize" json:"createBatchSize"`
	// TranslateError enabling error translation
	TranslateError bool `yaml:"translateError" json:"translateError"`
}

// Init init db
func (e *Database) Init() {
	var err error
	//parse env
	e.Source = pkg.ParseEnvTemplate(e.Source)
	for i := range e.Registers {
		for j := range e.Registers[i].Sources {
			e.Registers[i].Sources[j] = pkg.ParseEnvTemplate(e.Registers[i].Sources[j])
		}
		for j := range e.Registers[i].Replicas {
			e.Registers[i].Replicas[j] = pkg.ParseEnvTemplate(e.Registers[i].Replicas[j])
		}
	}

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
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default,
	}
	config.SkipDefaultTransaction = e.Config.SkipDefaultTransaction
	config.FullSaveAssociations = e.Config.FullSaveAssociations
	config.DryRun = e.Config.DryRun
	config.PrepareStmt = e.Config.PrepareStmt
	config.DisableAutomaticPing = e.Config.DisableAutomaticPing
	config.DisableForeignKeyConstraintWhenMigrating = e.Config.DisableForeignKeyConstraintWhenMigrating
	config.IgnoreRelationshipsWhenMigrating = e.Config.IgnoreRelationshipsWhenMigrating
	config.DisableNestedTransaction = e.Config.DisableNestedTransaction
	config.AllowGlobalUpdate = e.Config.AllowGlobalUpdate
	config.QueryFields = e.Config.QueryFields
	config.CreateBatchSize = e.Config.CreateBatchSize
	config.TranslateError = e.Config.TranslateError
	// gorm
	DB, err = resolverConfig.Init(config, Opens[e.Driver])
	if err != nil {
		slog.Error(e.Driver+" connect failed", slog.Any("err", err))
		os.Exit(-1)
	}
	// casbin
	if e.CasbinModel != "" {
		//set casbin adapter
		var a persist.Adapter
		a, err = gormadapter.NewAdapterByDBUseTableName(DB, "mss_boot", "casbin_rule")
		if err != nil {
			slog.Error("gormadapter.NewAdapterByDB error", slog.Any("err", err))
			os.Exit(-1)
		}
		var m model.Model
		m, err = model.NewModelFromString(e.CasbinModel)
		if err != nil {
			slog.Error("model.NewModelFromString error", slog.Any("err", err))
			os.Exit(-1)
		}
		Enforcer, err = casbin.NewEnforcer(m, a)
		if err != nil {
			slog.Error("casbin.NewEnforcer error", slog.Any("err", err))
			os.Exit(-1)
		}
		err = Enforcer.LoadPolicy()
		if err != nil {
			slog.Error("Enforcer.LoadPolicy error", slog.Any("err", err))
			os.Exit(-1)
		}
		Enforcer.EnableAutoSave(true)
		Enforcer.EnableAutoBuildRoleLinks(true)
		Enforcer.EnableLog(true)
		gorms.Driver = e.Driver
	}
}
