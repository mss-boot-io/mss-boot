package storage

import (
	"context"
	"time"

	"github.com/bsm/redislock"
)

const (
	PrefixKey = "__host"
)

var (
	Cache  AdapterCache
	Queue  AdapterQueue
	Locker AdapterLocker
)

type AdapterCache interface {
	String() string
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val interface{}, expire int) error
	Del(ctx context.Context, key string) error
	HashGet(ctx context.Context, hk, key string) (string, error)
	HashDel(ctx context.Context, hk, key string) error
	Increase(ctx context.Context, key string) error
	Decrease(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, dur time.Duration) error
}

type AdapterQueue interface {
	String() string
	Append(message Messager) error
	Register(name string, f ConsumerFunc)
	Run()
	Shutdown()
}

type Messager interface {
	SetID(string)
	SetStream(string)
	SetValues(map[string]interface{})
	GetID() string
	GetStream() string
	GetValues() map[string]interface{}
	GetPrefix() string
	SetPrefix(string)
}

type ConsumerFunc func(Messager) error

type AdapterLocker interface {
	String() string
	Lock(ctx context.Context, key string, ttl int64, options *redislock.Options) (*redislock.Lock, error)
}
