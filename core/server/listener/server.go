package listener

/*
 * @Author: lwnmengjing
 * @Date: 2021/6/8 2:04 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/8 2:04 下午
 */

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/pprof"
	"strings"

	ginPprof "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server server manage
type Server struct {
	ctx     context.Context
	srv     *http.Server
	options Options
	started bool
}

// New 实例化
func New(opts ...Option) server.Runnable {
	s := &Server{}

	s.Options(opts...)
	if s.options.handler == nil {
		return nil
	}
	switch h := s.options.handler.(type) {
	case *http.ServeMux:
		if s.options.pprof && h != http.DefaultServeMux {
			h.HandleFunc("/debug/pprof/", pprof.Index)
			h.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			h.HandleFunc("/debug/pprof/profile", pprof.Profile)
			h.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			h.HandleFunc("/debug/pprof/trace", pprof.Trace)
		}
		if s.options.metrics {
			h.Handle("/metrics", promhttp.Handler())
		}
		if s.options.healthz {
			h.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
		}
		if s.options.readyz {
			h.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
		}
		s.options.handler = h
	case *gin.Engine:
		if s.options.pprof {
			ginPprof.Register(h)
		}
		if s.options.metrics {
			h.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}
		if s.options.healthz {
			h.GET("/healthz", func(c *gin.Context) {
				c.AbortWithStatus(http.StatusOK)
			})
		}
		if s.options.readyz {
			h.GET("/readyz", func(c *gin.Context) {
				c.AbortWithStatus(http.StatusOK)
			})
		}
	}
	return s
}

// Options 设置参数
func (e *Server) Options(options ...Option) {
	e.options = *defaultOptions()
	for _, o := range options {
		o(&e.options)
	}
}

// String string
func (e *Server) String() string {
	return e.options.name
}

// Start server
func (e *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", e.options.addr)
	if err != nil {
		return err
	}
	e.ctx = ctx
	e.started = true
	e.srv = &http.Server{Handler: e.options.handler}
	if e.options.endHook != nil {
		e.srv.RegisterOnShutdown(e.options.endHook)
	}
	e.srv.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}
	go func() {
		if e.options.keyFile == "" || e.options.certFile == "" {
			if err = e.srv.Serve(l); err != nil {
				slog.ErrorContext(ctx, e.options.name+" Server start error", slog.Any("err", err.Error()))
			}
		} else {
			if err = e.srv.ServeTLS(l, e.options.certFile, e.options.keyFile); err != nil {
				slog.ErrorContext(ctx, e.options.name+" Server start error", slog.Any("err", err.Error()))
			}
		}
		<-ctx.Done()
		err = e.Shutdown(ctx)
		if err != nil {
			slog.ErrorContext(ctx, e.options.name+" Server shutdown error", slog.Any("err", err.Error()))
		}
	}()
	if e.options.startedHook != nil {
		e.options.startedHook()
	}
	e.PrintRunningInfo()
	return nil
}

// Shutdown 停止
func (e *Server) Shutdown(ctx context.Context) error {
	return e.srv.Shutdown(ctx)
}

func (e *Server) PrintRunningInfo() {
	var port string
	//todo 优雅输出
	fmt.Println("        \033[90m╔════════════════════════════════════════════════════╗\033[0m")
	fmt.Println("        \033[90m║\033[0m App listening at:                                  \033[90m║\033[0m")
	arr := strings.Split(e.options.addr, ":")
	if len(arr) <= 1 {
		port = "80"
	} else {
		port = arr[len(arr)-1]
	}
	hosts := []string{"localhost", "127.0.0.1"}
	if strings.Contains(e.options.addr, "0.0.0.0") || !strings.Contains(e.options.addr, ".") {
		interfaces, _ := net.Interfaces()
		for i := range interfaces {
			if interfaces[i].Flags&net.FlagUp == 0 {
				continue
			}
			addresses, _ := interfaces[i].Addrs()
			for _, addr := range addresses {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					hosts = append(hosts, ipNet.IP.String())
				}
			}
		}
	}
	prefix := "        "
	for i := range hosts {
		if i == 1 {
			prefix = "\033[32mready\033[0m - "
		}
		s := fmt.Sprintf("%s\033[90m║  >\033[0m   http://%s:%s", prefix, hosts[i], port)
		fmt.Println(s + strings.Repeat(
			" ",
			46-len(fmt.Sprintf("http://%s:%s", hosts[i], port)),
		) + "\033[90m║\033[0m")
		prefix = "        "
	}
	fmt.Println("        \033[90m║\033[0m                                                    \033[90m║\033[0m")
	fmt.Println("        \033[90m║\033[0m \033[1;97mNow you can open browser with the above addresses↑\033[0m \033[90m║\033[0m")
	fmt.Println("        \033[90m╚════════════════════════════════════════════════════╝\033[0m")
}
