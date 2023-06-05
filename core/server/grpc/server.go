package grpc

/*
 * @Author: lwnmengjing
 * @Date: 2021/6/2 4:26 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/2 4:26 下午
 */

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// Server server
type Server struct {
	name    string
	srv     *grpc.Server
	mux     sync.Mutex
	started bool
	options Options
}

// New grpc server
func New(name string, options ...Option) *Server {
	s := &Server{name: name}
	s.Options(options...)
	s.NewServer()
	return s
}

// String string
func (e *Server) String() string {
	return e.name
}

// Options set options
func (e *Server) Options(options ...Option) {
	e.options = *defaultOptions()
	for _, o := range options {
		o(&e.options)
	}
}

// Server return server
func (e *Server) Server() *grpc.Server {
	return e.srv
}

// NewServer new a server
func (e *Server) NewServer() {
	grpc.EnableTracing = false
	e.srv = grpc.NewServer(e.initGrpcServerOptions()...)
	reflection.Register(e.srv)
}

// Register register
func (e *Server) Register(do func(server *Server)) {
	do(e)
	prometheus.Register(e.srv)
}

func (e *Server) initGrpcServerOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(e.options.timeout),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(e.options.unaryServerInterceptors...)),
		grpc.StreamInterceptor(middleware.ChainStreamServer(e.options.streamServerInterceptors...)),
		grpc.MaxConcurrentStreams(uint32(e.options.maxConcurrentStreams)),
		grpc.MaxRecvMsgSize(e.options.maxMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: e.options.keepAlive / defaultMiniKeepAliveTimeRate,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:                  e.options.keepAlive,
			Timeout:               e.options.timeout,
			MaxConnectionAge:      e.options.maxConnectionAge,
			MaxConnectionAgeGrace: e.options.maxConnectionAgeGrace,
		}),
	}
	if e.options.certFile != "" && e.options.keyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(e.options.certFile, e.options.keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	return opts
}

// Start run
func (e *Server) Start(ctx context.Context) error {
	e.mux.Lock()
	defer e.mux.Unlock()

	if e.started {
		return errors.New("gRPC Server was started more than once. " +
			"This is likely to be caused by being added to a manage multiple times")
	}
	// Set the internal context
	if e.options.ctx != nil {
		ctx = e.options.ctx
	}

	ts, err := net.Listen("tcp", e.options.addr)
	if err != nil {
		return fmt.Errorf("gRPC Server listening on %s failed: %w", e.options.addr, err)
	}
	log.Infof("gRPC Server listening on %s", ts.Addr().String())

	go func() {
		if err = e.srv.Serve(ts); err != nil {
			log.Errorf("gRPC Server start error: %s", err.Error())
		}
	}()
	e.started = true
	<-ctx.Done()
	return e.Shutdown(ctx)
}

// Shutdown shutdown
func (e *Server) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	log.Info("gRPC Server will be shutdown gracefully")
	e.srv.GracefulStop()
	return nil
}
