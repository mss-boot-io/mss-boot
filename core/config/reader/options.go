package reader

import (
	"github.com/mss-boot-io/mss-boot/core/config/encoder"
	"github.com/mss-boot-io/mss-boot/core/config/encoder/json"
	"github.com/mss-boot-io/mss-boot/core/config/encoder/toml"
	"github.com/mss-boot-io/mss-boot/core/config/encoder/xml"
	"github.com/mss-boot-io/mss-boot/core/config/encoder/yaml"
)

// Options reader options
type Options struct {
	Encoding map[string]encoder.Encoder
}

// Option set options
type Option func(o *Options)

// NewOptions new options
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// WithEncoder set encoder
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
