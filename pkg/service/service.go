package service

import (
	"github.com/mss-boot-io/mss-boot/core/errcode"
	"github.com/mss-boot-io/mss-boot/core/logger"
)

type Service struct {
	Log *logger.Helper
}

// Make make
func (e *Service) Make(log *logger.Helper) {
	e.Log = log
}

func (e *Service) MakeError(id, domain string, code errcode.ErrorCoder) error {
	return errcode.New(id, domain, code)
}
