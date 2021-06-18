package config

import (
	"time"

	log "github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/core-go/tools/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Database database config
type Database struct {
	Driver          string
	Source          string
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
	Registers       []DBResolverConfig
}

// DBResolverConfig resolver config
type DBResolverConfig struct {
	Sources  []string
	Replicas []string
	Policy   string
	Tables   []string
}

// Init 初始化
func (e *Database) Init() (*gorm.DB, error) {
	registers := make([]database.ResolverConfigure, len(e.Registers))
	for i := range e.Registers {
		registers[i] = database.NewResolverConfigure(
			e.Registers[i].Sources,
			e.Registers[i].Replicas,
			e.Registers[i].Policy,
			e.Registers[i].Tables)
	}
	resolverConfig := database.NewConfigure(
		e.Source,
		e.MaxIdleConns,
		e.MaxOpenConns,
		e.ConnMaxIdleTime,
		e.ConnMaxLifeTime,
		registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: database.NewLogger(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.Level().ToGorm()),
			},
		),
	}, opens[e.Driver])
	return db, err
}
