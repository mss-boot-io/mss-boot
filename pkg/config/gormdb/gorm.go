package gormdb

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 17:02:39
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 17:02:39
 */

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
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
	// IAM aws rds iam auth
	IAM AWSRDSIAM `yaml:"iam"`
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

// AWSRDSIAM configures AWS RDS IAM authentication for database connections.
// When Enable is true, the database connection will use IAM authentication tokens
// instead of traditional username/password authentication.
type AWSRDSIAM struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Region string `yaml:"region" json:"region"`
	User   string `yaml:"user" json:"user"`
	Host   string `yaml:"host" json:"host"`
	Port   int    `yaml:"port" json:"port"`
	DBName string `yaml:"dbName" json:"dbName"`
	Params string `yaml:"params" json:"params"`
}

// Init init db
func (e *Database) Init() {
	var err error
	// parse env
	e.Source = pkg.ParseEnvTemplate(e.Source)
	e.IAM.Region = pkg.ParseEnvTemplate(e.IAM.Region)
	e.IAM.User = pkg.ParseEnvTemplate(e.IAM.User)
	e.IAM.Host = pkg.ParseEnvTemplate(e.IAM.Host)
	e.IAM.DBName = pkg.ParseEnvTemplate(e.IAM.DBName)
	e.IAM.Params = pkg.ParseEnvTemplate(e.IAM.Params)
	for i := range e.Registers {
		for j := range e.Registers[i].Sources {
			e.Registers[i].Sources[j] = pkg.ParseEnvTemplate(e.Registers[i].Sources[j])
		}
		for j := range e.Registers[i].Replicas {
			e.Registers[i].Replicas[j] = pkg.ParseEnvTemplate(e.Registers[i].Replicas[j])
		}
	}
	switch e.Driver {
	case gorms.Postgres:
		gorms.Driver = gorms.Postgres
	case gorms.Mysql:
		gorms.Driver = gorms.Mysql
	case gorms.Dm:
		gorms.Driver = gorms.Dm
	}

	if e.IAM.Enable {
		if e.IAM.Port == 0 {
			if e.Driver == gorms.Postgres {
				e.IAM.Port = 5432
			} else {
				e.IAM.Port = 3306
			}
		}
		cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
		if cfgErr != nil {
			slog.Error("failed to load AWS configuration for RDS IAM authentication", slog.Any("err", cfgErr))
			os.Exit(-1)
		}
		endpoint := fmt.Sprintf("%s:%d", e.IAM.Host, e.IAM.Port)
		token, tokErr := auth.BuildAuthToken(context.TODO(), endpoint, e.IAM.Region, e.IAM.User, cfg.Credentials)
		if tokErr != nil {
			slog.Error("failed to build RDS IAM authentication token",
				slog.String("endpoint", endpoint),
				slog.String("region", e.IAM.Region),
				slog.String("user", e.IAM.User),
				slog.Any("err", tokErr))
			os.Exit(-1)
		}
		if e.Driver == gorms.Postgres {
			dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", e.IAM.Host, e.IAM.Port, e.IAM.User, token, e.IAM.DBName)
			if e.IAM.Params != "" {
				dsn = dsn + " " + e.IAM.Params
			}
			// WARNING: The DSN below contains a temporary IAM authentication token.
			// Do NOT log or expose e.Source, as it contains sensitive credentials.
			e.Source = dsn
		} else {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true&allowCleartextPasswords=true", e.IAM.User, token, e.IAM.Host, e.IAM.Port, e.IAM.DBName)
			if e.IAM.Params != "" {
				dsn = dsn + "&" + e.IAM.Params
			}
			// WARNING: The DSN below contains a temporary IAM authentication token.
			// Do NOT log or expose e.Source, as it contains sensitive credentials.
			e.Source = dsn
		}
		if e.ConnMaxLifeTime == 0 || e.ConnMaxLifeTime > 840 {
			e.ConnMaxLifeTime = 840
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
		// set casbin adapter
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
