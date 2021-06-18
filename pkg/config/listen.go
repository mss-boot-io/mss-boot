/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 4:25 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 4:25 下午
 */

package config

import (
	"github.com/lwnmengjing/core-go/server/listener"
)

type Listen struct {
	Addr     string `json:"addr"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
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
