package config

/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 1:37 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 1:37 下午
 */

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/mss-boot-io/mss-boot/core/logger/writer"
)

// Logger logger配置
type Logger struct {
	Path      string     `yaml:"path" json:"path"`
	Level     slog.Level `yaml:"level" json:"level"`
	Stdout    string     `yaml:"stdout" json:"stdout"`
	AddSource bool       `yaml:"addSource" json:"addSource"`
	Cap       uint       `yaml:"cap" json:"cap"`
	Json      bool       `yaml:"json" json:"json"`
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
		output = os.Stderr
	}
	if e.Json {
		slog.SetDefault(slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
			AddSource: e.AddSource,
			Level:     e.Level,
		})))
		return
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
		AddSource: e.AddSource,
		Level:     e.Level,
	})))
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
