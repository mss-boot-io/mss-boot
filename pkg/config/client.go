/*
 * @Author: lwnmengjing
 * @Date: 2021/6/21 2:34 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/21 2:34 下午
 */

package config

import (
	"strings"
	"time"

	"github.com/mss-boot-io/mss-boot/client"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
)

// Client client
type Client struct {
	Addr    string `yaml:"addr" json:"addr"`
	Timeout int    `yaml:"timeout" json:"timeout"`
}

// Clients clients
type Clients map[string]Client

// Init init
func (e Clients) Init() error {

	opts := make([]client.Option, 0)
	for k := range e {
		s := grpc.Service{}
		err := s.Dial(e[k].Addr,
			time.Duration(e[k].Timeout)*time.Second)
		if err != nil {
			return err
		}
		switch strings.ToLower(k) {
		//case "tenant", "tenants":
		//	opts = append(opts, client.WithTenant(client.NewTenant(s)))
		case "store", "stores":
			opts = append(opts, client.WithStore(client.NewStore(s)))
		}
	}
	client.Default.Init(opts...)
	return nil
}
