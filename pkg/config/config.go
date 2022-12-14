/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

package config

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
	sourceFS "github.com/mss-boot-io/mss-boot/pkg/config/source/fs"
	sourceLocal "github.com/mss-boot-io/mss-boot/pkg/config/source/local"
	sourceS3 "github.com/mss-boot-io/mss-boot/pkg/config/source/s3"
)

// Init 初始化配置
func Init(cfg any, options ...source.Option) (err error) {
	opts := &source.Options{}
	for _, opt := range options {
		opt(opts)
	}
	var f fs.ReadFileFS
	switch opts.Provider {
	case source.FS:
		f, err = sourceFS.New(options...)
	case source.Local:
		f, err = sourceLocal.New(options...)
	case source.S3:
		s := &Storage{
			Type:            ProviderType(os.Getenv("s3_provider")),
			SigningMethod:   os.Getenv("s3_signing_method"),
			Region:          os.Getenv("s3_region"),
			Bucket:          os.Getenv("s3_bucket"),
			AccessKeyID:     os.Getenv("s3_access_key_id"),
			SecretAccessKey: os.Getenv("s3_secret_access_key"),
		}
		s.Init()
		options = append(options,
			source.WithBucket(s.Bucket), source.WithClient(s.GetClient()))
		f, err = sourceS3.New(options...)
	}
	if err != nil {
		return err
	}

	var rb []byte
	rb, err = f.ReadFile("application.yml")
	if err != nil {
		err = nil
		rb, err = f.ReadFile("application.yaml")
	}
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(rb, cfg)
	if err != nil {
		return err
	}
	stage := os.Getenv("stage")
	if stage == "" {
		stage = os.Getenv("STAGE")
	}
	if stage == "" {
		stage = "local"
	}
	rb, err = f.ReadFile(fmt.Sprintf("application-%s.yml", stage))
	if err != nil {
		err = nil
		rb, err = f.ReadFile(fmt.Sprintf("application-%s.yaml", stage))
		if err != nil {
			return nil
		}
	}
	return yaml.Unmarshal(rb, cfg)
}
