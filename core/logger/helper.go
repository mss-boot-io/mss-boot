package logger

import (
	"os"

	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

type Helper struct {
	Logger
	fields map[string]interface{}
}

func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Info) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Info, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Info) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Info, template, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Trace) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Trace, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Trace) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Trace, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Debug) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Debug, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Debug) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Debug, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Warn) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Warn, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Warn) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Warn, template, args...)
}

func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Error) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Error, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Error) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Error, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Fatal) {
		return
	}
	h.Logger.Fields(h.fields).Log(level.Fatal, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level.Fatal) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level.Fatal, template, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}
