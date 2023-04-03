/*
 * @Author: lwnmengjing
 * @Date: 2021/6/8 2:04 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/8 2:04 下午
 */

package listener

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server server manage
type Server struct {
	name    string
	ctx     context.Context
	srv     *http.Server
	opts    options
	started bool
}

// New 实例化
func New(name string, opts ...Option) server.Runnable {
	s := &Server{
		name: name,
		opts: setDefaultOption(),
	}
	s.Options(opts...)
	return s
}

// NewMetrics 新建默认监控服务
func NewMetrics(opts ...Option) server.Runnable {
	s := &Server{
		name: "metrics",
		opts: setDefaultOption(),
	}
	s.opts.addr = ":3000"
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())
	s.opts.handler = h
	s.Options(opts...)
	return s
}

// NewHealthz default heealthz server
func NewHealthz(opts ...Option) server.Runnable {
	s := &Server{
		name: "healthz",
		opts: setDefaultOption(),
	}
	s.opts.addr = ":4000"
	h := http.NewServeMux()
	h.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	s.opts.handler = h
	s.Options(opts...)
	return s
}

// NewReadyz default readyz server
func NewReadyz(opts ...Option) server.Runnable {
	s := &Server{
		name: "readyz",
		opts: setDefaultOption(),
	}
	s.opts.addr = ":2000"
	h := http.NewServeMux()
	h.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	s.opts.handler = h
	s.Options(opts...)
	return s
}

func NewPprof(opts ...Option) server.Runnable {
	s := &Server{
		name: "pprof",
		opts: setDefaultOption(),
	}
	s.opts.addr = ":5000"
	s.opts.handler = http.DefaultServeMux
	s.Options(opts...)
	return s
}

// Options 设置参数
func (e *Server) Options(opts ...Option) {
	for _, o := range opts {
		o(&(e.opts))
	}
}

// String string
func (e *Server) String() string {
	return e.name
}

// Start server
func (e *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", e.opts.addr)
	if err != nil {
		return err
	}
	e.ctx = ctx
	e.started = true
	e.srv = &http.Server{Handler: e.opts.handler}
	if e.opts.endHook != nil {
		e.srv.RegisterOnShutdown(e.opts.endHook)
	}
	e.srv.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}
	log.Infof("%s Server listening on %s", e.name, l.Addr().String())
	go func() {
		if e.opts.keyFile == "" || e.opts.certFile == "" {
			if err = e.srv.Serve(l); err != nil {
				log.Errorf("%s Server start error: %s", e.name, err.Error())
			}
		} else {
			if err = e.srv.ServeTLS(l, e.opts.certFile, e.opts.keyFile); err != nil {
				log.Errorf("%s Server start error: %s", e.name, err.Error())
			}
		}
		<-ctx.Done()
		err = e.Shutdown(ctx)
		if err != nil {
			log.Errorf("%S Server shutdown error: %s", e.name, err.Error())
		}
	}()
	if e.opts.startedHook != nil {
		e.opts.startedHook()
	}
	return nil
}

// Shutdown 停止
func (e *Server) Shutdown(ctx context.Context) error {
	return e.srv.Shutdown(ctx)
}
