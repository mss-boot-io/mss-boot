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

	"github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/logger/level"
	"github.com/mss-boot-io/mss-boot/core/logger/writer"
	"github.com/mss-boot-io/mss-boot/core/plugins/logger/logrus"
	"github.com/mss-boot-io/mss-boot/core/plugins/logger/zap"
	logrusCore "github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

// Logger logger配置
type Logger struct {
	Type      string `yaml:"type" json:"type"`
	Path      string `yaml:"path" json:"path"`
	Level     string `yaml:"level" json:"level"`
	Stdout    string `yaml:"stdout" json:"stdout"`
	Cap       uint   `yaml:"cap" json:"cap"`
	Formatter string `yaml:"formatter" json:"formatter"`
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

	opts := []logger.Option{
		logger.WithLevel(l),
		logger.WithOutput(output),
	}
	switch e.Type {
	case "zap":
		opts = append(opts, zap.WithCallerSkip(2))
		switch e.Formatter {
		case "json":
			opts = append(opts, zap.WithEncoder(zapcore.NewJSONEncoder(zapcore.EncoderConfig{})))
		}
		logger.DefaultLogger, err = zap.NewLogger(opts...)
		if err != nil {
			log.Fatalf("new zap logger error, %s", err.Error())
		}
	case "logrus":
		opts = append(opts, logrus.WithSkip(12), logrus.ReportCaller())
		switch e.Formatter {
		case "json":
			opts = append(opts, logrus.WithJSONFormatter(&logrusCore.JSONFormatter{}))
		}
		logger.DefaultLogger = logrus.NewLogger(opts...)
	default:
		logger.DefaultLogger = logger.NewLogger(opts...)
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
