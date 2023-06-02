package zap

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mss-boot-io/mss-boot/core/logger"
)

// Options is zap logger options
type Options struct {
	logger.Options
}

type callerSkipKey struct{}

// WithCallerSkip pass caller skip to logger
func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type encoderKey struct{}

// WithEncoder pass zapcore.Encoder to logger
func WithEncoder(e zapcore.Encoder) logger.Option {
	return logger.SetOption(encoderKey{}, e)
}

type namespaceKey struct{}

// WithNamespace set namespace
func WithNamespace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}

type writerKey struct{}

// WithOutput set output
func WithOutput(out io.Writer) logger.Option {
	return logger.SetOption(writerKey{}, out)
}
