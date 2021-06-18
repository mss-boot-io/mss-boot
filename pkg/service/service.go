package service

import (
	"github.com/lwnmengjing/core-go/errcode"
	"github.com/lwnmengjing/core-go/logger"
	"gorm.io/gorm"
)

type Service struct {
	Orm *gorm.DB
	Log *logger.Helper
}

// Make make
func (e *Service) Make(orm *gorm.DB, log *logger.Helper) {
	e.Orm = orm
	e.Log = log
}

func (e *Service) MakeError(id, domain string, code errcode.ErrorCoder) error {
	return errcode.New(id, domain, code)
}
