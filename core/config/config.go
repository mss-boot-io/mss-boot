// Package cfg is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/mss-boot-io/mss-boot/core/config/loader"
	"github.com/mss-boot-io/mss-boot/core/config/reader"
	"github.com/mss-boot-io/mss-boot/core/config/source"
	"github.com/mss-boot-io/mss-boot/core/config/source/file"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Values provide the reader.Values interface
	reader.Values
	// Init the cfg
	Init(opts ...Option) error
	// Options in the cfg
	Options() Options
	// Close Stop the cfg loader/watcher
	Close() error
	// Load cfg sources
	Load(source ...source.Source) error
	// Sync Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}

// Watcher is the cfg watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

// Entity 配置实体
type Entity interface {
	OnChange()
}

// Options 配置的参数
type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context

	Entity Entity
}

// Option 调用类型
type Option func(o *Options)

var (
	// DefaultConfig Default Config Manager
	DefaultConfig Config
)

// NewConfig returns new cfg
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}

// Bytes Return cfg as raw json
func Bytes() []byte {
	return DefaultConfig.Bytes()
}

// Map Return cfg as a map
func Map() map[string]interface{} {
	return DefaultConfig.Map()
}

// Scan values to a go type
func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}

// Sync Force a source changeset sync
func Sync() error {
	return DefaultConfig.Sync()
}

// Get a value from the cfg
func Get(path ...string) reader.Value {
	return DefaultConfig.Get(path...)
}

// Load cfg sources
func Load(source ...source.Source) error {
	return DefaultConfig.Load(source...)
}

// Watch a value for changes
func Watch(path ...string) (Watcher, error) {
	return DefaultConfig.Watch(path...)
}

// LoadFile is short hand for creating a file source and loading it
func LoadFile(path string) error {
	return Load(file.NewSource(
		file.WithPath(path),
	))
}
