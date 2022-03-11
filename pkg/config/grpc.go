package config

import (
	"time"

	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
)

// GRPC grpc服务公共配置(选用)
type GRPC struct {
	Network string // default: tcp
	Addr    string // default:  :9090
	Timeout int    // default: 30
	Name    string // default:
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
		grpc.WithTimeout(time.Duration(e.Timeout)*time.Second),
		grpc.WithID(e.Name),
	)
	s := grpc.New("grpc", opts...)
	s.Register(register)
	return s
}
