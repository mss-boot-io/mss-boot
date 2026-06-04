package cache

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func testRedisClient(t *testing.T) redis.UniversalClient {
	t.Helper()
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        []string{"127.0.0.1:6379"},
		DialTimeout:  200 * time.Millisecond,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		MaxRetries:   0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		t.Skipf("skip redis integration test: %v", err)
	}
	return client
}

func TestRedis_HashSet(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
	}
	type args struct {
		ctx    context.Context
		hk     string
		key    string
		val    interface{}
		expire time.Duration
	}
	client := testRedisClient(t)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "hash set",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:    context.Background(),
				hk:     "h-set-test",
				key:    "h-set-key",
				val:    "h-set-value",
				expire: time.Minute,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
			}
			if err := r.HSet(tt.args.ctx, tt.args.hk, tt.args.key, tt.args.val, tt.args.expire).Err(); (err != nil) != tt.wantErr {
				t.Errorf("HashSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_HashGet(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
	}
	type args struct {
		ctx context.Context
		hk  string
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "hash get",
			fields: fields{
				client: testRedisClient(t),
			},
			args: args{
				ctx: context.Background(),
				hk:  "h-set-test",
				key: "h-set-key",
			},
			want:    "h-set-value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
			}
			got, err := r.HGet(tt.args.ctx, tt.args.hk, tt.args.key).Result()
			if (err != nil) != tt.wantErr {
				t.Errorf("HashGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_HashDel(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
	}
	type args struct {
		ctx context.Context
		hk  string
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "hash del",
			fields: fields{
				client: testRedisClient(t),
			},
			args: args{
				ctx: context.Background(),
				hk:  "h-set-test",
				key: "h-set-key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
			}
			if err := r.HDel(tt.args.ctx, tt.args.hk, tt.args.key).Err(); (err != nil) != tt.wantErr {
				t.Errorf("HashDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Set(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
		opts   Options
	}
	type args struct {
		ctx    context.Context
		key    string
		val    interface{}
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set",
			fields: fields{
				client: testRedisClient(t),
				opts: Options{
					QueryCacheDuration: 10 * time.Second,
				},
			},
			args: args{
				ctx:    context.Background(),
				key:    "set-test",
				val:    "set-value",
				expire: 10 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "set with empty value",
			fields: fields{
				client: testRedisClient(t),
				opts: Options{
					QueryCacheDuration: 10 * time.Second,
				},
			},
			args: args{
				ctx:    context.Background(),
				key:    "set-test",
				val:    "",
				expire: 10 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "set with int value",
			fields: fields{
				client: testRedisClient(t),
				opts: Options{
					QueryCacheDuration: 10 * time.Second,
				},
			},
			args: args{
				ctx:    context.Background(),
				key:    "set-test-int",
				val:    123,
				expire: 10 * time.Second,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
				opts:            tt.fields.opts,
			}
			if err := r.Set(tt.args.ctx, tt.args.key, tt.args.val, tt.args.expire).Err(); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Get(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get",
			fields: fields{
				client: testRedisClient(t),
			},
			args: args{
				ctx: context.Background(),
				key: "set-test",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
			}
			got, err := r.Get(tt.args.ctx, tt.args.key).Result()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Del(t *testing.T) {
	type fields struct {
		client redis.UniversalClient
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "del",
			fields: fields{
				client: testRedisClient(t),
			},
			args: args{
				ctx: context.Background(),
				key: "set-test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				UniversalClient: tt.fields.client,
			}
			if err := r.Del(tt.args.ctx, tt.args.key).Err(); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
