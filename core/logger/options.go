// Package logger nolint
package logger

import (
	"context"
	"io"

	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

// Option set Options
type Option func(*Options)

// Options options
type Options struct {
	// The logging Level the logger should log at. default is `InfoLevel`
	Level level.Level
	// Fields to always be logged
	Fields map[string]interface{}
	// It's common to set this to a file, or leave it default which is `os.Stderr`
	Out io.Writer
	// Caller skip frame count for file:line info
	CallerSkipCount int
	// Alternative Options
	Context context.Context
	// Name logger Name
	Name string
}

// WithFields set default Fields for the logger
func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

// WithLevel set default Level for the logger
func WithLevel(level level.Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithOutput set default output writer for the logger
func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithCallerSkipCount set frame count to skip
func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

// WithName set Name for logger
func WithName(name string) Option {
	return func(args *Options) {
		args.Name = name
	}
}

// SetOption set option
func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
