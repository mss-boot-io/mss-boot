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
)

// Init 初始化配置
// fixme 配置刷新功能比较鸡肋，去除
func Init(f fs.ReadFileFS, cfg interface{}) (err error) {
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
