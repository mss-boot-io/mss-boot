/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 1:37 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 1:37 下午
 */

package config

import (
	"io"
	"log"
	"os"

	"github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/core-go/logger/level"
	"github.com/lwnmengjing/core-go/logger/writer"
)

// Logger logger配置
type Logger struct {
	Type   string `yaml:"type"`
	Path   string `yaml:"path"`
	Level  string `yaml:"level"`
	Stdout string `yaml:"stdout"`
	Cap    uint   `yaml:"cap"`
}

// Init 初始化日志
func (e *Logger) Init() {
	var err error
	var output io.Writer
	switch e.Stdout {
	case "file":
		if !pathExist(e.Path) {
			err := pathCreate(e.Path)
			if err != nil {
				log.Fatalf("create dir error: %s", err.Error())
			}
		}
		output, err = writer.NewFileWriter(
			writer.WithPath(e.Path),
			writer.WithSuffix("log"),
			writer.WithCap(e.Cap),
		)
		if err != nil {
			log.Fatalf("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	var l level.Level
	l, err = level.GetLevel(e.Level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}

	switch e.Type {
	//case "zap":
	//	setLogger, err = zap.NewLogger(logger.WithLevel(l), logger.WithOutput(output), zap.WithCallerSkip(2))
	//	if err != nil {
	//		log.Fatalf("new zap logger error, %s", err.Error())
	//	}
	//case "logrus":
	//	setLogger = logrus.NewLogger(logger.WithLevel(l), logger.WithOutput(output), logrus.ReportCaller())
	default:
		logger.DefaultLogger = logger.NewLogger(logger.WithLevel(l), logger.WithOutput(output))
	}
}

func pathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// pathExist 判断目录是否存在
func pathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return s.IsDir()
}
