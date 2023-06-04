package logger

import (
	"os"

	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

// Helper is a logger helper
type Helper struct {
	// Logger is the logger
	Logger
	fields map[string]interface{}
}

// NewHelper new a logger helper
func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

// Info info level
func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Info) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Info, args...)
}

// Infof info level
func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Info) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Info, template, args...)
}

// Trace trace level
func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Trace) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Trace, args...)
}

// Tracef trace level
func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Trace) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Trace, template, args...)
}

// Debug debug level
func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Debug) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Debug, args...)
}

// Debugf debug level
func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Debug) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Debug, template, args...)
}

// Warn warn level
func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Warn) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Warn, args...)
}

// Warnf warn level
func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Warn) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Warn, template, args...)
}

// Error error level
func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Error) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Error, args...)
}

// Errorf error level
func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Error) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Error, template, args...)
}

// Fatal fatal level
func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Fatal) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Fatal, args...)
	os.Exit(1)
}

// Fatalf fatal level
func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Fatal) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Fatal, template, args...)
	os.Exit(1)
}

// WithError with error
func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

// WithFields with fields
func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}
