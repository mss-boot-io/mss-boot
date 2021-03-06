/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 4:25 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 4:25 下午
 */

package config

import (
	"github.com/mss-boot-io/mss-boot/core/server/listener"
)

type Listen struct {
	Addr     string `yaml:"addr" json:"addr"`
	CertFile string `yaml:"certFile" json:"certFile"`
	KeyFile  string `yaml:"keyFile" json:"keyFile"`
	Timeout  int    `yaml:"timeout" json:"timeout"` // default: 10s
}

func (e *Listen) Init(opts ...listener.Option) []listener.Option {
	if e == nil {
		return nil
	}
	if opts == nil {
		opts = make([]listener.Option, 0)
	}
	if e.Timeout == 0 {
		e.Timeout = 10
	}
	opts = append(opts,
		listener.WithAddr(e.Addr),
		listener.WithCert(e.CertFile),
		listener.WithKey(e.KeyFile),
		listener.WithTimeout(e.Timeout),
	)
	return opts
}
