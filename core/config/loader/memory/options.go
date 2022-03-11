package memory

import (
	"github.com/mss-boot-io/mss-boot/core/config/loader"
	"github.com/mss-boot-io/mss-boot/core/config/reader"
	"github.com/mss-boot-io/mss-boot/core/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the cfg reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
