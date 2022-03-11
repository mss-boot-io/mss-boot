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
}

func (e *Listen) Init(opts ...listener.Option) []listener.Option {
	if e == nil {
		return nil
	}
	if opts == nil {
		opts = make([]listener.Option, 0)
	}
	opts = append(opts,
		listener.WithAddr(e.Addr),
		listener.WithCert(e.CertFile),
		listener.WithKey(e.KeyFile),
	)
	return opts
}
