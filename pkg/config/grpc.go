package config

import (
	"time"

	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
)

// GRPC grpc服务公共配置(选用)
type GRPC struct {
	Addr     string `yaml:"addr" json:"addr"` // default:  :9090
	CertFile string `yaml:"certFile" json:"certFile"`
	KeyFile  string `yaml:"keyFile" json:"keyFile"`
	Timeout  int    `yaml:"timeout" json:"timeout"` // default: 10
}

// Init init
func (e *GRPC) Init(
	register func(srv *grpc.Server),
	opts ...grpc.Option) server.Runnable {
	if opts == nil {
		opts = make([]grpc.Option, 0)
	}
	opts = append(opts,
		grpc.WithAddr(e.Addr),
		grpc.WithKey(e.KeyFile),
		grpc.WithCert(e.CertFile),
		grpc.WithTimeout(time.Duration(e.Timeout)*time.Second),
	)
	s := grpc.New("grpc", opts...)
	s.Register(register)
	return s
}
