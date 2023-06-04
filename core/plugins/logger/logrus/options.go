package logrus

import (
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/mss-boot-io/mss-boot/core/logger"
)

// Options Options
type Options struct {
	logger.Options
	Formatter logrus.Formatter
	Hooks     logrus.LevelHooks
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Exit Function to call when FatalLevel log
	ExitFunc        func(int)
	CallerSkipCount int
}

type formatterKey struct{}

// WithTextTextFormatter withTextTextFormatter
func WithTextTextFormatter(formatter *logrus.TextFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

// WithJSONFormatter withJSONFormatter
func WithJSONFormatter(formatter *logrus.JSONFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

type hooksKey struct{}

// WithLevelHooks withLevelHooks
func WithLevelHooks(hooks logrus.LevelHooks) logger.Option {
	return logger.SetOption(hooksKey{}, hooks)
}

type reportCallerKey struct{}

// ReportCaller warning to use this option. because logrus doest not open CallerDepth option
// this will only print this package
func ReportCaller() logger.Option {
	return logger.SetOption(reportCallerKey{}, true)
}

type exitKey struct{}

// WithExitFunc withExitFunc
func WithExitFunc(exit func(int)) logger.Option {
	return logger.SetOption(exitKey{}, exit)
}

type logrusLoggerKey struct{}

// WithLogger withLogger
func WithLogger(l logrus.StdLogger) logger.Option {
	return logger.SetOption(logrusLoggerKey{}, l)
}

type skipKey struct{}

// WithSkip set skip
func WithSkip(skip int) logger.Option {
	return logger.SetOption(skipKey{}, skip)
}

func callerPrettyfier(skip int) func(f *runtime.Frame) (string, string) {
	return func(f *runtime.Frame) (string, string) {
		pc := make([]uintptr, 25)
		runtime.Callers(skip, pc)
		frames := runtime.CallersFrames(pc)
		f1, _ := frames.Next()
		f = &f1
		funcName := f.Func.Name()
		fileName := f.File + ":" + strconv.Itoa(f.Line)
		return funcName, fileName
	}
}
