/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/12/11 07:33:01
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/12/11 07:33:01
 */

package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ProviderType storage provider type
type ProviderType string

const (
	// S3 aws s3
	S3 ProviderType = "s3"
	// OSS aliyun oss
	OSS ProviderType = "oss"
	// OOS ctyun oos
	OOS ProviderType = "oos"
	// KODO qiniu kodo
	KODO ProviderType = "kodo"
	// COS tencent cos
	COS ProviderType = "cos"
	// OBS huawei obs
	OBS ProviderType = "obs"
	// BOS baidu bos
	BOS ProviderType = "bos"
	// GCS google gcs fixme:not tested
	GCS ProviderType = "gcs"
	// KS3 kingsoft ks3
	KS3 ProviderType = "ks3"
)

// URLTemplate storage provider url template
var URLTemplate = map[ProviderType]string{
	OSS:  "https://%s.aliyuncs.com",
	OOS:  "https://oos-%s.ctyunapi.cn",
	KODO: "https://s3-%s.qiniucs.com",
	COS:  "https://cos.%s.myqcloud.com",
	OBS:  "https://obs.%s.myhuaweicloud.com",
	BOS:  "https://s3.%s.bcebos.com",
	GCS:  "https://storage.googleapis.com",
	KS3:  "https://ks3-%s.ksyuncs.com",
}

var endpointResolverFunc = func(urlTemplate, signingMethod string) s3.EndpointResolverFunc {
	return func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           fmt.Sprintf(urlTemplate, region),
			SigningRegion: region,
			SigningMethod: signingMethod,
		}, nil
	}
}

// Storage storage
type Storage struct {
	Type            ProviderType `yaml:"type"`
	SigningMethod   string       `yaml:"signingMethod"`
	Region          string       `yaml:"region"`
	Bucket          string       `yaml:"bucket"`
	AccessKeyID     string       `yaml:"accessKeyID"`
	SecretAccessKey string       `yaml:"secretAccessKey"`
	client          *s3.Client
}

// Init init
func (o *Storage) Init() {
	var endpointResolver s3.EndpointResolver
	if o.Type != S3 {
		if urlTemplate, exist := URLTemplate[o.Type]; exist && urlTemplate != "" {
			endpointResolver = endpointResolverFunc(urlTemplate, o.SigningMethod)
		}
	}
	if o.Region == "" || o.AccessKeyID == "" || o.SecretAccessKey == "" {
		//use default config
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("failed to load SDK configuration, %v", err)
		}
		o.client = s3.NewFromConfig(cfg)
		return
	}

	o.client = s3.New(s3.Options{
		Region: o.Region,
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     o.AccessKeyID,
				SecretAccessKey: o.SecretAccessKey,
			}, nil
		}),
		EndpointResolver: endpointResolver,
	}, func(o *s3.Options) {
		o.EndpointOptions.DisableHTTPS = true
	})
}

// GetClient get client
func (o *Storage) GetClient() *s3.Client {
	return o.client
}
