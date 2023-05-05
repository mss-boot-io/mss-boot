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
	ctx     context.Context
	srv     *http.Server
	opts    options
	started bool
}

// New 实例化
func New(opts ...Option) server.Runnable {
	s := &Server{
		opts: setDefaultOption(),
	}

	s.opts.handler = http.DefaultServeMux
	s.Options(opts...)

	if s.opts.metrics {
		http.Handle("/metrics", promhttp.Handler())
	}
	if s.opts.healthz {
		http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	}
	if s.opts.readyz {
		http.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	}
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
	return e.opts.name
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
	log.Infof("%s Server listening on %s", e.opts.name, l.Addr().String())
	go func() {
		if e.opts.keyFile == "" || e.opts.certFile == "" {
			if err = e.srv.Serve(l); err != nil {
				log.Errorf("%s Server start error: %s", e.opts.name, err.Error())
			}
		} else {
			if err = e.srv.ServeTLS(l, e.opts.certFile, e.opts.keyFile); err != nil {
				log.Errorf("%s Server start error: %s", e.opts.name, err.Error())
			}
		}
		<-ctx.Done()
		err = e.Shutdown(ctx)
		if err != nil {
			log.Errorf("%S Server shutdown error: %s", e.opts.name, err.Error())
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
