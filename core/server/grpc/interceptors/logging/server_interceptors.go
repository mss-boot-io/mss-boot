package logging

/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 9:13 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 9:13 下午
 */

import (
	"context"
	"path"
	"time"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mss-boot-io/mss-boot/core/server/grpc/interceptors/logging/ctxlog"
	"github.com/mss-boot-io/mss-boot/core/tools/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds zap.Logger to the context.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateServerOpt(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		newCtx := newLoggerForCall(ctx, info.FullMethod, startTime, o.timestampFormat)

		resp, err := handler(newCtx, req)
		if !o.shouldLog(info.FullMethod, err) {
			return resp, err
		}
		code := o.codeFunc(err)
		level := o.levelFunc(code)
		duration := o.durationFunc(time.Since(startTime))

		o.messageFunc(newCtx, "finished unary call with code "+code.String(), level, code, err, &duration)
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateServerOpt(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()
		newCtx := newLoggerForCall(stream.Context(), info.FullMethod, startTime, o.timestampFormat)
		wrapped := middleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx

		err := handler(srv, wrapped)
		if !o.shouldLog(info.FullMethod, err) {
			return err
		}
		code := o.codeFunc(err)
		level := o.levelFunc(code)
		duration := o.durationFunc(time.Since(startTime))

		o.messageFunc(newCtx, "finished streaming call with code "+code.String(), level, code, err, &duration)
		return err
	}
}

func serverCallFields(fullMethodString string) *ctxlog.Fields {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	f := ctxlog.NewFields("system", "grpc")
	f.Merge(ctxlog.NewFields("span.kind", "server"))
	f.Set("grpc.service", service)
	f.Set("grpc.method", method)
	return f
}

func newLoggerForCall(ctx context.Context, fullMethodString string, start time.Time, timestampFormat string) context.Context {
	f := serverCallFields(fullMethodString)
	f.Set("grpc.start_time", start.Format(timestampFormat))
	if d, ok := ctx.Deadline(); ok {
		f.Set("grpc.request.deadline", d.Format(timestampFormat))
	}
	requestID := utils.GetRequestID(ctx)
	f.Set(utils.RequestIDKey, requestID)
	callLog := ctxlog.Extract(ctx).Fields(f.Values())
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		utils.RequestIDKey, requestID,
	)
	return ctxlog.ToContext(ctx, callLog)
}
