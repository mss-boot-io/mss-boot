/*
 * @Author: lwnmengjing
 * @Date: 2021/6/2 4:30 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/2 4:30 下午
 */

package grpc

import (
	"context"
	"crypto/tls"
	"math"
	"time"

	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/mss-boot-io/mss-boot/core/errcode"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server/grpc/interceptors/logging"
	requesttag "github.com/mss-boot-io/mss-boot/core/server/grpc/interceptors/request_tag"
	"google.golang.org/grpc"
)

const (
	infinity                           = time.Duration(math.MaxInt64)
	defaultMaxMsgSize                  = 4 << 20
	defaultMaxConcurrentStreams        = 100000
	defaultKeepAliveTime               = 30 * time.Second
	defaultConnectionIdleTime          = 10 * time.Second
	defaultMaxServerConnectionAgeGrace = 10 * time.Second
	defaultMiniKeepAliveTimeRate       = 2
)

// Option set options
type Option func(*Options)

// Options options
type Options struct {
	id                       string
	domain                   string
	addr                     string
	tls                      *tls.Config
	keepAlive                time.Duration
	timeout                  time.Duration
	maxConnectionAge         time.Duration
	maxConnectionAgeGrace    time.Duration
	maxConcurrentStreams     int
	maxMsgSize               int
	unaryServerInterceptors  []grpc.UnaryServerInterceptor
	streamServerInterceptors []grpc.StreamServerInterceptor
	ctx                      context.Context
}

// WithContext set ContextOption
func WithContext(c context.Context) Option {
	return func(o *Options) {
		o.ctx = c
	}
}

// WithID set IDOption
func WithID(s string) Option {
	return func(o *Options) {
		o.id = s
	}
}

// WithDomain set DomainOption
func WithDomain(s string) Option {
	return func(o *Options) {
		o.domain = s
	}
}

// WithAddr set AddrOption
func WithAddr(s string) Option {
	return func(o *Options) {
		o.addr = s
	}
}

// WithTLS set TlsOption
func WithTLS(tls *tls.Config) Option {
	return func(o *Options) {
		o.tls = tls
	}
}

// WithKeepAlive set KeepAliveOption
func WithKeepAlive(t time.Duration) Option {
	return func(o *Options) {
		o.keepAlive = t
	}
}

// WithTimeout set TimeoutOption
func WithTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.keepAlive = t
	}
}

// WithMaxConnectionAge set MaxConnectionAgeOption
func WithMaxConnectionAge(t time.Duration) Option {
	return func(o *Options) {
		o.maxConnectionAge = t
	}
}

// WithMaxConnectionAgeGrace set MaxConnectionAgeGraceOption
func WithMaxConnectionAgeGrace(t time.Duration) Option {
	return func(o *Options) {
		o.maxConnectionAgeGrace = t
	}
}

// WithMaxConcurrentStreamsOption set MaxConcurrentStreamsOption
func WithMaxConcurrentStreamsOption(i int) Option {
	return func(o *Options) {
		o.maxConcurrentStreams = i
	}
}

// WithMaxMsgSizeOption set MaxMsgSizeOption
func WithMaxMsgSizeOption(i int) Option {
	return func(o *Options) {
		o.maxMsgSize = i
	}
}

// WithUnaryServerInterceptors set UnaryServerInterceptorsOption
func WithUnaryServerInterceptors(u ...grpc.UnaryServerInterceptor) Option {
	return func(o *Options) {
		if o.unaryServerInterceptors == nil {
			o.unaryServerInterceptors = make([]grpc.UnaryServerInterceptor, 0)
		}
		o.unaryServerInterceptors = append(o.unaryServerInterceptors, u...)
	}
}

// WithStreamServerInterceptors set StreamServerInterceptorsOption
func WithStreamServerInterceptors(u ...grpc.StreamServerInterceptor) Option {
	return func(o *Options) {
		if o.streamServerInterceptors == nil {
			o.streamServerInterceptors = make([]grpc.StreamServerInterceptor, 0)
		}
		o.streamServerInterceptors = append(o.streamServerInterceptors, u...)
	}
}

func defaultOptions() *Options {
	return &Options{
		addr:                  ":0",
		keepAlive:             defaultKeepAliveTime,
		timeout:               defaultConnectionIdleTime,
		maxConnectionAge:      infinity,
		maxConnectionAgeGrace: defaultMaxServerConnectionAgeGrace,
		maxConcurrentStreams:  defaultMaxConcurrentStreams,
		maxMsgSize:            defaultMaxMsgSize,
		unaryServerInterceptors: []grpc.UnaryServerInterceptor{
			requesttag.UnaryServerInterceptor(),
			ctxtags.UnaryServerInterceptor(),
			opentracing.UnaryServerInterceptor(),
			logging.UnaryServerInterceptor(),
			prometheus.UnaryServerInterceptor,
			//recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(customRecovery("", ""))),
		},
		streamServerInterceptors: []grpc.StreamServerInterceptor{
			requesttag.StreamServerInterceptor(),
			ctxtags.StreamServerInterceptor(),
			opentracing.StreamServerInterceptor(),
			logging.StreamServerInterceptor(),
			prometheus.StreamServerInterceptor,
			//recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(customRecovery("", ""))),
		},
	}
}

// customRecovery customRecovery
func customRecovery(id, domain string) recovery.RecoveryHandlerFunc {
	return func(p interface{}) (err error) {
		log.Errorf("panic triggered: %v", p)
		return errcode.New(id, domain, errcode.GRPCInternalServerError)
	}
}
