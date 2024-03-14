package config

/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:31 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:31 下午
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"text/template"
	"text/template/parse"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
	sourceFS "github.com/mss-boot-io/mss-boot/pkg/config/source/fs"
	"github.com/mss-boot-io/mss-boot/pkg/config/source/gorm"
	sourceLocal "github.com/mss-boot-io/mss-boot/pkg/config/source/local"
	"github.com/mss-boot-io/mss-boot/pkg/config/source/mgdb"
	sourceS3 "github.com/mss-boot-io/mss-boot/pkg/config/source/s3"
	"gopkg.in/yaml.v3"
)

// Init 初始化配置
func Init(cfg source.Entity, options ...source.Option) (err error) {
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
			Endpoint:        os.Getenv("s3_endpoint"),
			AccessKeyID:     os.Getenv("s3_access_key_id"),
			SecretAccessKey: os.Getenv("s3_secret_access_key"),
		}
		s.Init()
		options = append(options,
			source.WithBucket(s.Bucket), source.WithClient(s.GetClient()))
		f, err = sourceS3.New(options...)
	case source.MGDB:
		f, err = mgdb.New(options...)
	case source.GORM:
		f, err = gorm.New(options...)
	}
	if err != nil {
		return err
	}

	stage := getStage()

	var rb []byte
	rb, err = f.ReadFile(opts.Name)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	rb, err = parseTemplateWithEnv(rb)
	if err != nil {
		return err
	}
	var unm func([]byte, any) error
	switch f.GetExtend() {
	case source.SchemeYaml, source.SchemeYml:
		unm = yaml.Unmarshal
	case source.SchemeJSOM:
		unm = json.Unmarshal
	}
	err = unm(rb, cfg)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	rb, err = f.ReadFile(fmt.Sprintf("%s-%s", opts.Name, stage))
	if err == nil {
		rb, err = parseTemplateWithEnv(rb)
		if err != nil {
			return err
		}
		err = unm(rb, cfg)
		if err != nil {
			slog.Error(err.Error())
		}
	}
	if !opts.Watch {
		return nil
	}
	return f.Watch(cfg, unm)
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

func parseTemplateWithEnv(rb []byte) ([]byte, error) {
	t, err := template.New("env").Parse(string(rb))
	if err != nil {
		return nil, err
	}
	tree, err := parse.Parse("env", string(rb), "{{", "}}")
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	data := getValueFromEnv(getParseKeys(tree["env"].Root))
	err = t.Execute(&buffer, data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func getValueFromEnv(keys []string) any {
	env := make(map[string]string)
	for i := range keys {
		if !strings.Contains(strings.ToLower(keys[i]), "env.") {
			continue
		}
		keyArr := strings.Split(keys[i], ".")
		if len(keyArr) > 1 {
			keyArr = keyArr[1:]
		}
		key := strings.Join(keyArr, ".")
		var exist bool
		env[key], exist = os.LookupEnv(key)
		if exist {
			continue
		}
		env[key], exist = os.LookupEnv(strings.ToUpper(key))
		if exist {
			continue
		}
		env[key], exist = os.LookupEnv(strings.ToLower(key))
		if exist {
			continue
		}
		env[key] = ""
	}
	return map[string]any{
		"Env": env,
	}
}

// getParseKeys get parse keys from template text
func getParseKeys(nodes *parse.ListNode) []string {
	keys := make([]string, 0)
	if nodes == nil {
		return keys
	}
	for a := range nodes.Nodes {
		if actionNode, ok := nodes.Nodes[a].(*parse.ActionNode); ok {
			if actionNode == nil || actionNode.Pipe == nil {
				continue
			}
			for b := range actionNode.Pipe.Cmds {
				if strings.Index(actionNode.Pipe.Cmds[b].String(), ".") == 0 {
					keys = append(keys, actionNode.Pipe.Cmds[b].String()[1:])
				}
			}
		}
	}
	return keys
}
