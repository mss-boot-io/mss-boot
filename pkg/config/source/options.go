package source

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:31:13
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:31:13
 */

import (
	"io/fs"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Provider provider
type Provider string

const (
	// FS fs
	FS Provider = "fs"
	// Local local
	Local Provider = "local"
	// S3 s3
	S3 Provider = "s3"
	// MGDB mongodb
	MGDB Provider = "mgdb"
	// GORM gorm
	GORM Provider = "gorm"
)

// Extends extends
var Extends = []Scheme{SchemeYaml, SchemeYml, SchemeJSOM}

// Sourcer source interface
type Sourcer interface {
	fs.ReadFileFS
	GetExtend() Scheme
}

// Option option
type Option func(*Options)

// Options options
type Options struct {
	Provider          Provider
	Name              string
	Extend            Scheme
	Dir               string
	Region            string
	Bucket            string
	ProjectName       string
	Timeout           time.Duration
	Client            *s3.Client
	FS                fs.ReadFileFS
	MongoDBURL        string
	MongoDBName       string
	MongoDBCollection string
	Datasource        string
}

// DefaultOptions default options
func DefaultOptions() *Options {
	return &Options{
		Provider: Local,
		Name:     "application",
		Dir:      "cfg",
	}
}

// WithDatasource set datasource
func WithDatasource(datasource string) Option {
	return func(args *Options) {
		args.Datasource = datasource
	}
}

// WithMongoDBURL set mongodb url
func WithMongoDBURL(url string) Option {
	return func(args *Options) {
		if url == "" {
			url = "mongodb://localhost:27017"
		}
		args.MongoDBURL = url
	}
}

// WithMongoDBName set mongodb name
func WithMongoDBName(name string) Option {
	return func(args *Options) {
		args.MongoDBName = name
	}
}

// WithMongoDBCollection set mongodb collection
func WithMongoDBCollection(collection string) Option {
	return func(args *Options) {
		args.MongoDBCollection = collection
	}
}

// WithProvider set provider
func WithProvider(provider Provider) Option {
	return func(args *Options) {
		args.Provider = provider
	}
}

// WithDir set dir
func WithDir(dir string) Option {
	return func(args *Options) {
		args.Dir = strings.ReplaceAll(dir, "\\", "/")
	}
}

// WithName set config name
func WithName(file string) Option {
	return func(args *Options) {
		args.Name = strings.ReplaceAll(file, "\\", "/")
	}
}

// WithProjectName set projectName
func WithProjectName(projectName string) Option {
	return func(args *Options) {
		args.ProjectName = projectName
	}
}

// WithRegion set s3 region
func WithRegion(region string) Option {
	return func(args *Options) {
		args.Region = region
	}
}

// WithBucket set s3 bucket
func WithBucket(bucket string) Option {
	return func(args *Options) {
		args.Bucket = bucket
	}
}

// WithTimeout set s3 client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(args *Options) {
		args.Timeout = timeout
	}
}

// WithClient set s3 client
func WithClient(client *s3.Client) Option {
	return func(args *Options) {
		args.Client = client
	}
}

// WithFrom set embed.FS
func WithFrom(fs fs.ReadFileFS) Option {
	return func(args *Options) {
		args.FS = fs
	}
}
