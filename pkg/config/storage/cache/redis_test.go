package cache

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

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
	client := redis.NewUniversalClient(&redis.UniversalOptions{})
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
				client: tt.fields.client,
			}
			if err := r.HashSet(tt.args.ctx, tt.args.hk, tt.args.key, tt.args.val, tt.args.expire); (err != nil) != tt.wantErr {
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: tt.fields.client,
			}
			got, err := r.HashGet(tt.args.ctx, tt.args.hk, tt.args.key)
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: tt.fields.client,
			}
			if err := r.HashDel(tt.args.ctx, tt.args.hk, tt.args.key); (err != nil) != tt.wantErr {
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: tt.fields.client,
				opts:   tt.fields.opts,
			}
			if err := r.Set(tt.args.ctx, tt.args.key, tt.args.val, tt.args.expire); (err != nil) != tt.wantErr {
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: tt.fields.client,
			}
			got, err := r.Get(tt.args.ctx, tt.args.key)
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
				client: redis.NewUniversalClient(&redis.UniversalOptions{}),
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
				client: tt.fields.client,
			}
			if err := r.Del(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
