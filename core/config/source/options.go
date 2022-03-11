package source

import (
	"context"

	"github.com/mss-boot-io/mss-boot/core/config/encoder"
	"github.com/mss-boot-io/mss-boot/core/config/encoder/json"
)

// Options params
type Options struct {
	// Encoder
	Encoder encoder.Encoder

	// for alternative data
	Context context.Context
}

// Option options set
type Option func(o *Options)

// NewOptions new default options
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: json.NewEncoder(),
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// WithEncoder sets the source encoder
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}
