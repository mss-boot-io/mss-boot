// Package logger provides a log interface nolint
package logger

import (
	"os"

	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

var (
	// DefaultLogger logger
	DefaultLogger Logger
)

// Logger is a generic logging interface
type Logger interface {
	// Init initialises Options
	Init(options ...Option) error
	// Options The Logger Options
	Options() Options
	// Fields set Fields to always be logged
	Fields(fields map[string]interface{}) Logger
	// Log writes a log entry
	Log(level level.Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level level.Level, format string, v ...interface{})
	// String returns the Name of logger
	String() string
}

// Init init default logger
func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

// Info Log writes a log entry
func Info(args ...interface{}) {
	DefaultLogger.Log(level.Info, args...)
}

// Infof Log writes a log entry
func Infof(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Info, template, args...)
}

// Trace Log writes a log entry
func Trace(args ...interface{}) {
	DefaultLogger.Log(level.Trace, args...)
}

// Tracef Log writes a log entry
func Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Trace, template, args...)
}

// Debug Log writes a log entry
func Debug(args ...interface{}) {
	DefaultLogger.Log(level.Debug, args...)
}

// Debugf Log writes a log entry
func Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Debug, template, args...)
}

// Warn Log writes a log entry
func Warn(args ...interface{}) {
	DefaultLogger.Log(level.Warn, args...)
}

// Warnf Log writes a log entry
func Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Warn, template, args...)
}

// Error Log writes a log entry
func Error(args ...interface{}) {
	DefaultLogger.Log(level.Error, args...)
}

// Errorf Log writes a log entry
func Errorf(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Error, template, args...)
}

// Fatal Log writes a log entry
func Fatal(args ...interface{}) {
	DefaultLogger.Log(level.Fatal, args...)
	os.Exit(1)
}

// Fatalf Log writes a log entry
func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Logf(level.Fatal, template, args...)
	os.Exit(1)
}

// Log writes a log entry
func Log(l level.Level, v ...interface{}) {
	DefaultLogger.Log(l, v...)
}

// Logf writes a log entry
func Logf(l level.Level, format string, v ...interface{}) {
	DefaultLogger.Logf(l, format, v...)
}

// V Returns true if the given Level is at or lower the current logger Level
func V(lvl level.Level, log Logger) bool {
	l := DefaultLogger
	if log != nil {
		l = log
	}
	return l.Options().Level <= lvl
}

// Level return logger Level
func Level() level.Level {
	return DefaultLogger.Options().Level
}
