package config

import (
	"time"

	"google.golang.org/grpc"

	"github.com/mss-boot-io/mss-boot/core/server"
	serverGRPC "github.com/mss-boot-io/mss-boot/core/server/grpc"
)

// GRPC grpc服务公共配置(选用)
type GRPC struct {
	Addr     string                `yaml:"addr" json:"addr"` // default:  :9090
	CertFile string                `yaml:"certFile" json:"certFile"`
	KeyFile  string                `yaml:"keyFile" json:"keyFile"`
	Timeout  int                   `yaml:"timeout" json:"timeout"` // default: 10
	Client   map[string]GRPCClient `yaml:"client" json:"client"`
}

// Init init
func (e *GRPC) Init(
	register func(srv *serverGRPC.Server),
	opts ...serverGRPC.Option) server.Runnable {
	if opts == nil {
		opts = make([]serverGRPC.Option, 0)
	}
	opts = append(opts,
		serverGRPC.WithAddr(e.Addr),
		serverGRPC.WithKey(e.KeyFile),
		serverGRPC.WithCert(e.CertFile),
		serverGRPC.WithTimeout(time.Duration(e.Timeout)*time.Second),
	)
	s := serverGRPC.New("grpc", opts...)
	s.Register(register)
	return s
}

type GRPCClient struct {
	conn    *grpc.ClientConn
	Address string        `yaml:"address" json:"address"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

func (g *GRPCClient) Init() error {
	//cc, err := grpc.Dial(g.Address,
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithChainUnaryInterceptor(
	//		timeout.UnaryClientInterceptor(g.Timeout),
	//
	//	))
	return nil
}

type Clients map[string]GRPCClient

func (cs Clients) Init() error {
	return nil
}
