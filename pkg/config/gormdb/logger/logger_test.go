package logger

import (
	"context"
	"testing"
	"time"

	loggerCore "github.com/mss-boot-io/mss-boot/core/logger"
	"gorm.io/gorm/logger"
)

func TestNew(t *testing.T) {
	l := New(logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel: logger.LogLevel(
			loggerCore.DefaultLogger.Options().Level.ToGorm()),
	})
	l.Info(context.TODO(), "test")
}
