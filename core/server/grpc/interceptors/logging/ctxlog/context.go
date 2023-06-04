// Package ctxlog nolint
package ctxlog

import (
	"context"

	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

type ctxMarker struct{}

type ctxLogger struct {
	logger logger.Logger
	fields map[string]interface{}
}

var (
	ctxMarkerKey = &ctxMarker{}
)

// AddFields adds logger fields to the logger.
func AddFields(ctx context.Context, fields map[string]interface{}) {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return
	}
	for k, v := range fields {
		l.fields[k] = v
	}
}

// Extract takes the call-scoped Log from grpc_logger middleware.
// It always returns a Log that has all the grpc_ctxtags updated.
func Extract(ctx context.Context) logger.Logger {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return logger.DefaultLogger
	}
	// Add grpc_ctxtags tags metadata until now.
	fields := TagsToFields(ctx)
	// Add logger fields added until now.
	for k, v := range l.fields {
		fields[k] = v
	}
	return l.logger.Fields(fields)
}

// TagsToFields transforms the Tags on the supplied context into logger fields.
func TagsToFields(ctx context.Context) map[string]interface{} {
	return grpcCtxTags.Extract(ctx).Values()
}

// ToContext adds the logger.Logger to the context for extraction later.
// Returning the new context that has been created.
func ToContext(ctx context.Context, logger logger.Logger) context.Context {
	l := &ctxLogger{
		logger: logger,
	}
	return context.WithValue(ctx, ctxMarkerKey, l)
}

// Debug is equivalent to calling Debug on the logger.Logger in the context.
// It is a no-op if the context does not contain a logger.Logger.
func Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	Extract(ctx).Fields(fields).Log(level.Debug, msg)
}

// Info is equivalent to calling Info on the logger.Logger in the context.
// It is a no-op if the context does not contain a logger.Logger.
func Info(ctx context.Context, msg string, fields map[string]interface{}) {
	Extract(ctx).Fields(fields).Log(level.Info, msg)
}

// Warn is equivalent to calling Warn on the logger.Logger in the context.
// It is a no-op if the context does not contain a logger.Logger.
func Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	Extract(ctx).Fields(fields).Log(level.Warn, msg)
}

// Error is equivalent to calling Error on the logger.Logger in the context.
// It is a no-op if the context does not contain a logger.Logger.
func Error(ctx context.Context, msg string, fields map[string]interface{}) {
	Extract(ctx).Fields(fields).Log(level.Error, msg)
}
