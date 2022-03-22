/*
 * @Author: lwnmengjing
 * @Date: 2022/3/21 14:51
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/21 14:51
 */

package cfg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"github.com/robinjoseph08/redisqueue/v2"
	"store/pkg/storage"
	"store/pkg/storage/cache"
	"store/pkg/storage/locker"
	"store/pkg/storage/queue"
	"time"
)

type Redis struct {
	Network    string                      `yaml:"network" json:"network"`
	Addr       string                      `yaml:"addr" json:"addr"`
	Username   string                      `yaml:"username" json:"username"`
	Password   string                      `yaml:"password" json:"password"`
	DB         int                         `yaml:"db" json:"db"`
	PoolSize   int                         `yaml:"pool_size" json:"pool_size"`
	Tls        *config.Tls                 `yaml:"tls" json:"tls"`
	MaxRetries int                         `yaml:"max_retries" json:"max_retries"`
	Producer   *redisqueue.ProducerOptions `yaml:"producer" json:"producer"`
	Consumer   *redisqueue.ConsumerOptions `yaml:"consumer" json:"consumer"`
}

func (e Redis) GetRedisOptions() (*redis.Options, error) {
	r := &redis.Options{
		Network:    e.Network,
		Addr:       e.Addr,
		Username:   e.Username,
		Password:   e.Password,
		DB:         e.DB,
		MaxRetries: e.MaxRetries,
		PoolSize:   e.PoolSize,
	}
	var err error
	r.TLSConfig, err = e.Tls.GetTLS()
	return r, err
}

func (Redis) String() string {
	return "redis"
}

func (e *Redis) getClient() (*cache.Redis, error) {
	opts, err := e.GetRedisOptions()
	if err != nil {
		return nil, err
	}
	cr, ok := storage.Cache.(*cache.Redis)
	var r *cache.Redis
	if ok && cr != nil {
		r, err = cache.NewRedis(context.TODO(), cr.GetClient(), opts)
	} else {
		r, err = cache.NewRedis(context.TODO(), nil, opts)
	}
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (e *Redis) GetCache() (storage.AdapterCache, error) {
	return e.getClient()
}

func (e *Redis) GetQueue() (storage.AdapterQueue, error) {
	e.Consumer.ReclaimInterval = e.Consumer.ReclaimInterval * time.Second
	e.Consumer.BlockingTimeout = e.Consumer.BlockingTimeout * time.Second
	e.Consumer.VisibilityTimeout = e.Consumer.VisibilityTimeout * time.Second
	client, err := e.getClient()
	if err != nil {
		return nil, err
	}
	e.Producer.RedisClient = client.GetClient()
	e.Consumer.RedisClient = client.GetClient()
	return queue.NewRedis(e.Producer, e.Consumer)
}

func (e *Redis) GetLocker() (storage.AdapterLocker, error) {
	client, err := e.getClient()
	if err != nil {
		return nil, err
	}
	return locker.NewRedis(client.GetClient()), nil
}
