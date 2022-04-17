package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// NewRedis redis模式
func NewRedis(ctx context.Context, client *redis.Client, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = redis.NewClient(options)
	}
	r := &Redis{
		client: client,
	}
	err := r.connect(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *redis.Client
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect(ctx context.Context) error {
	var err error
	_, err = r.client.Ping(ctx).Result()
	return err
}

// Get from key
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(ctx context.Context, key string, val interface{}, expire int) error {
	return r.client.Set(ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// HashGet from key
func (r *Redis) HashGet(ctx context.Context, hk, key string) (string, error) {
	return r.client.HGet(ctx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(ctx context.Context, hk, key string) error {
	return r.client.HDel(ctx, hk, key).Err()
}

// Increase +1
func (r *Redis) Increase(ctx context.Context, key string) error {
	return r.client.Incr(ctx, key).Err()
}

// Decrease -1
func (r *Redis) Decrease(ctx context.Context, key string) error {
	return r.client.Decr(ctx, key).Err()
}

// Expire Set ttl
func (r *Redis) Expire(ctx context.Context, key string, dur time.Duration) error {
	return r.client.Expire(ctx, key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}
