package queue

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mss-boot-io/redisqueue/v2"
	"github.com/redis/go-redis/v9"

	"github.com/mss-boot-io/mss-boot/pkg/config/storage"
)

func TestRedis_Append(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue.ConsumerOptions
		ProducerOptions *redisqueue.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue.Consumer
		producer        *redisqueue.Producer
	}
	type args struct {
		name    string
		message storage.Messager
	}
	client := redis.NewUniversalClient(&redis.UniversalOptions{})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: false,
					RedisClient:          client,
				},
			},
			args{
				name: "test",
				message: &Message{Message: redisqueue.Message{
					ID:     "",
					Stream: "test",
					Values: map[string]interface{}{
						"key": "value",
					},
				}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				if err := r.Append(storage.WithMessage(tt.args.message)); (err != nil) != tt.wantErr {
					t.Errorf("SetQueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}

	t.Log("ok")
}

func TestRedis_Register(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue.ConsumerOptions
		ProducerOptions *redisqueue.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue.Consumer
		producer        *redisqueue.Producer
	}
	type args struct {
		name string
		f    storage.ConsumerFunc
	}
	client := redis.NewClient(&redis.Options{})
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: true,
					RedisClient:          client,
				},
			},
			args{
				name: "test",
				f: func(message storage.Messager) error {
					fmt.Println("ok")
					fmt.Println(message.GetValues())
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				r.Register(storage.WithTopic(tt.args.name), storage.WithConsumerFunc(tt.args.f))
				go func() {
					time.Sleep(time.Second)
					r.Shutdown()
				}()
				r.Run(context.Background())
			}
		})
	}
}
