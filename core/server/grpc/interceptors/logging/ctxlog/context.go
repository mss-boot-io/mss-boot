package ctxlog

import (
	"context"
	"log/slog"

	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

type ctxMarker struct{}

type ctxLogger struct {
	logger *slog.Logger
	fields []any
}

var (
	ctxMarkerKey = &ctxMarker{}
)

// AddFields adds logger fields to the logger.
func AddFields(ctx context.Context, fields ...any) {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return
	}
	if l.fields == nil {
		l.fields = make([]any, 0)
	}
	l.fields = append(l.fields, fields...)
}

// Extract takes the call-scoped Log from grpc_logger middleware.
// It always returns a Log that has all the grpc_ctxtags updated.
func Extract(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return slog.Default()
	}
	// Add grpc_ctxtags tags metadata until now.
	fields := TagsToFields(ctx)
	// Add logger fields added until now.
	return l.logger.With(fields...)
}

// TagsToFields transforms the Tags on the supplied context into logger fields.
func TagsToFields(ctx context.Context) []any {
	args := make([]any, 0)
	for k, v := range grpcCtxTags.Extract(ctx).Values() {
		args = append(args, k, v)
	}
	return args
}

// ToContext adds the *slog.Logger to the context for extraction later.
// Returning the new context that has been created.
func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	l := &ctxLogger{
		logger: logger,
	}
	return context.WithValue(ctx, ctxMarkerKey, l)
}

// Debug is equivalent to calling Debug on the *slog.Logger in the context.
// It is a no-op if the context does not contain a *slog.Logger.
func Debug(ctx context.Context, msg string, fields ...any) {
	Extract(ctx).With(fields...).Log(ctx, slog.LevelDebug, msg)
}

// Info is equivalent to calling Info on the *slog.Logger in the context.
// It is a no-op if the context does not contain a *slog.Logger.
func Info(ctx context.Context, msg string, fields ...any) {
	Extract(ctx).With(fields...).Log(ctx, slog.LevelInfo, msg)
}

// Warn is equivalent to calling Warn on the *slog.Logger in the context.
// It is a no-op if the context does not contain a *slog.Logger.
func Warn(ctx context.Context, msg string, fields ...any) {
	Extract(ctx).With(fields...).Log(ctx, slog.LevelWarn, msg)
}

// Error is equivalent to calling Error on the *slog.Logger in the context.
// It is a no-op if the context does not contain a *slog.Logger.
func Error(ctx context.Context, msg string, fields ...any) {
	Extract(ctx).With(fields...).Log(ctx, slog.LevelError, msg)
}
