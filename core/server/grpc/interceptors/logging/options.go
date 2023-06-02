/*
 * @Author: lwnmengjing
 * @Date: 2021/5/19 11:14 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/19 11:14 上午
 */

package logging

import (
	"context"
	"time"

	grpcLogging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"github.com/mss-boot-io/mss-boot/core/logger/level"
	"github.com/mss-boot-io/mss-boot/core/server/grpc/interceptors/logging/ctxlog"
	"google.golang.org/grpc/codes"
)

var (
	defaultOptions = &Options{
		levelFunc:       DefaultCodeToLevel,
		shouldLog:       grpcLogging.DefaultDeciderMethod,
		codeFunc:        grpcLogging.DefaultErrorToCode,
		durationFunc:    DefaultDurationToField,
		messageFunc:     DefaultMessageProducer,
		timestampFormat: time.RFC3339,
	}
)

// Options Options
type Options struct {
	levelFunc       CodeToLevel
	shouldLog       grpcLogging.Decider
	codeFunc        grpcLogging.ErrorToCode
	durationFunc    DurationToField
	messageFunc     MessageProducer
	timestampFormat string
}

// Option set Options
type Option func(*Options)

// CodeToLevel function defines the mapping between gRPC return codes and interceptor log level.
type CodeToLevel func(code codes.Code) level.Level

// DurationToField function defines how to produce duration fields for logging
type DurationToField func(duration time.Duration) ctxlog.Fields

// WithDecider customizes the function for deciding if the gRPC interceptor logs should log.
func WithDecider(f grpcLogging.Decider) Option {
	return func(o *Options) {
		o.shouldLog = f
	}
}

// WithLevels customizes the function for mapping gRPC return codes and interceptor log level statements.
func WithLevels(f CodeToLevel) Option {
	return func(o *Options) {
		o.levelFunc = f
	}
}

// WithCodes customizes the function for mapping errors to error codes.
func WithCodes(f grpcLogging.ErrorToCode) Option {
	return func(o *Options) {
		o.codeFunc = f
	}
}

// WithDurationField customizes the function for mapping request durations to Zap fields.
func WithDurationField(f DurationToField) Option {
	return func(o *Options) {
		o.durationFunc = f
	}
}

// WithMessageProducer customizes the function for message formation.
func WithMessageProducer(f MessageProducer) Option {
	return func(o *Options) {
		o.messageFunc = f
	}
}

// WithTimestampFormat customizes the timestamps emitted in the log fields.
func WithTimestampFormat(format string) Option {
	return func(o *Options) {
		o.timestampFormat = format
	}
}

// MessageProducer produces a user defined log message
type MessageProducer func(ctx context.Context, msg string, level level.Level, code codes.Code, err error, duration *ctxlog.Fields)

func evaluateServerOpt(opts []Option) *Options {
	optCopy := &Options{}
	*optCopy = *defaultOptions
	optCopy.levelFunc = DefaultCodeToLevel
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

func evaluateClientOpt(opts []Option) *Options {
	optCopy := &Options{}
	*optCopy = *defaultOptions
	optCopy.levelFunc = DefaultClientCodeToLevel
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

// DefaultCodeToLevel is the default implementation of gRPC return codes and interceptor log level for server side.
func DefaultCodeToLevel(code codes.Code) level.Level {
	switch code {
	case codes.OK:
		return level.Info
	case codes.Canceled:
		return level.Info
	case codes.Unknown:
		return level.Error
	case codes.InvalidArgument:
		return level.Info
	case codes.DeadlineExceeded:
		return level.Warn
	case codes.NotFound:
		return level.Info
	case codes.AlreadyExists:
		return level.Info
	case codes.PermissionDenied:
		return level.Warn
	case codes.Unauthenticated:
		return level.Info // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return level.Warn
	case codes.FailedPrecondition:
		return level.Warn
	case codes.Aborted:
		return level.Warn
	case codes.OutOfRange:
		return level.Warn
	case codes.Unimplemented:
		return level.Error
	case codes.Internal:
		return level.Error
	case codes.Unavailable:
		return level.Warn
	case codes.DataLoss:
		return level.Error
	default:
		return level.Error
	}
}

// DefaultClientCodeToLevel is the default implementation of gRPC return codes to log levels for client side.
func DefaultClientCodeToLevel(code codes.Code) level.Level {
	switch code {
	case codes.OK:
		return level.Debug
	case codes.Canceled:
		return level.Debug
	case codes.Unknown:
		return level.Info
	case codes.InvalidArgument:
		return level.Debug
	case codes.DeadlineExceeded:
		return level.Info
	case codes.NotFound:
		return level.Debug
	case codes.AlreadyExists:
		return level.Debug
	case codes.PermissionDenied:
		return level.Info
	case codes.Unauthenticated:
		return level.Info // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return level.Debug
	case codes.FailedPrecondition:
		return level.Debug
	case codes.Aborted:
		return level.Debug
	case codes.OutOfRange:
		return level.Debug
	case codes.Unimplemented:
		return level.Warn
	case codes.Internal:
		return level.Warn
	case codes.Unavailable:
		return level.Warn
	case codes.DataLoss:
		return level.Warn
	default:
		return level.Info
	}
}

// DefaultDurationToField is the default implementation of converting request duration to a Zap field.
var DefaultDurationToField = DurationToTimeMillisField

// DurationToTimeMillisField converts the duration to milliseconds and uses the key `grpc.time_ms`.
func DurationToTimeMillisField(duration time.Duration) ctxlog.Fields {
	return *ctxlog.NewFields("grpc.time_ms", durationToMilliseconds(duration))
}

// DurationToDurationField uses a Duration field to log the request duration
// and leaves it up to Zap's encoder settings to determine how that is output.
func DurationToDurationField(duration time.Duration) map[string]interface{} {
	return map[string]interface{}{"grpc.duration": duration}
}

func durationToMilliseconds(duration time.Duration) float32 {
	return float32(duration.Nanoseconds()/1000) / 1000
}

// DefaultMessageProducer writes the default message
func DefaultMessageProducer(ctx context.Context, msg string, level level.Level, code codes.Code, err error, duration *ctxlog.Fields) {
	// re-extract logger from newCtx, as it may have extra fields that changed in the holder.
	fields := duration
	fields.Set("grpc.code", code.String())
	if err != nil {
		fields.Set("grpc.error", err)
	}
	ctxlog.Extract(ctx).Fields(fields.Values()).Log(level, msg)
}
