package config

import (
	"github.com/mss-boot-io/mss-boot/core/config/loader"
	"github.com/mss-boot-io/mss-boot/core/config/reader"
	"github.com/mss-boot-io/mss-boot/core/config/source"
)

// WithLoader sets the loader for manage cfg
func WithLoader(l loader.Loader) Option {
	return func(o *Options) {
		o.Loader = l
	}
}

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the cfg reader
func WithReader(r reader.Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}

// WithEntity sets the cfg Entity
func WithEntity(e Entity) Option {
	return func(o *Options) {
		o.Entity = e
	}
}
