/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

package config

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/source"
	sourceFS "github.com/mss-boot-io/mss-boot/pkg/config/source/fs"
	sourceLocal "github.com/mss-boot-io/mss-boot/pkg/config/source/local"
	"github.com/mss-boot-io/mss-boot/pkg/config/source/mgdb"
	sourceS3 "github.com/mss-boot-io/mss-boot/pkg/config/source/s3"
	"gopkg.in/yaml.v3"
)

// Init 初始化配置
func Init(cfg any, options ...source.Option) (err error) {
	opts := source.DefaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	var f source.Sourcer
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
	case source.MGDB:
		f, err = mgdb.New(options...)
		if err != nil {
			return err
		}
		var rb []byte
		rb, err = f.ReadFile(opts.Name)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(rb, cfg)
	}
	if err != nil {
		return err
	}

	stage := getStage()

	var rb []byte
	rb, err = f.ReadFile(opts.Name)
	if err != nil {
		log.Errorf("read file error: %v", err)
		return err
	}
	var unm func([]byte, interface{}) error
	switch f.GetExtend() {
	case source.SchemeYaml, source.SchemeYml:
		unm = yaml.Unmarshal
	case source.SchemeJSOM:
		unm = json.Unmarshal
	}
	err = unm(rb, cfg)
	if err != nil {
		log.Errorf("unmarshal error: %v", err)
		return err
	}

	rb, err = f.ReadFile(fmt.Sprintf("%s-%s", opts.Name, stage))
	if err == nil {
		err = unm(rb, cfg)
		if err != nil {
			log.Errorf("unmarshal error: %v", err)
		}
	}
	return nil
}

func getStage() string {
	stage := os.Getenv("stage")
	if stage == "" {
		stage = os.Getenv("STAGE")
	}
	if stage == "" {
		stage = "local"
	}
	return stage
}
