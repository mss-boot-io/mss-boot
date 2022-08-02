/*
 * @Author: lwnmengjing
 * @Date: 2022/7/22 01:58:01
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/22 01:58:01
 */

package s3

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Option set Options
type Option func(*Options)

type Options struct {
	region  string
	bucket  string
	dir     string
	timeout time.Duration
	client  *s3.Client
}

// WithRegion set s3 region
func WithRegion(region string) Option {
	return func(args *Options) {
		args.region = region
	}
}

// WithBucket set s3 bucket
func WithBucket(bucket string) Option {
	return func(args *Options) {
		args.bucket = bucket
	}
}

// WithDir set s3 dir
func WithDir(dir string) Option {
	return func(args *Options) {
		args.dir = dir
	}
}

// WithTimeout set s3 client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(args *Options) {
		args.timeout = timeout
	}
}

// WithClient set s3 client
func WithClient(client *s3.Client) Option {
	return func(args *Options) {
		args.client = client
	}
}
