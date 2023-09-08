package zap

import (
	"fmt"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/logger/level"
	"github.com/mss-boot-io/mss-boot/core/logger/writer"
)

func TestName(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	if l.String() != "zap" {
		t.Errorf("name is error %s", l.String())
	}

	t.Logf("test logger name: %s", l.String())
}

func TestLogf(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	logger.DefaultLogger = l
	logger.Logf(level.Info, "test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	logger.Init(logger.WithLevel(level.Debug))
	l.Logf(level.Debug, "test show debug: %s", "debug msg")

	logger.Init(logger.WithLevel(level.Info))
	l.Logf(level.Debug, "test non-show debug: %s", "debug msg")
}

func TestWithReportCaller(t *testing.T) {
	var err error
	logger.DefaultLogger, err = NewLogger(WithCallerSkip(2))
	if err != nil {
		t.Fatal(err)
	}

	logger.Logf(level.Info, "testing: %s", "WithReportCaller")
}

func TestFields(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l.Fields(map[string]interface{}{
		"x-request-id": "123456abc",
	})
	logger.DefaultLogger.Log(level.Info, "hello")
}

func TestFile(t *testing.T) {
	output, err := writer.NewFileWriter(
		writer.WithPath("testdata"),
		writer.WithSuffix("log"),
	)
	if err != nil {
		t.Errorf("logger setup error: %s", err.Error())
	}
	//var err error
	logger.DefaultLogger, err = NewLogger(logger.WithLevel(level.Trace), WithOutput(output))
	if err != nil {
		t.Errorf("logger setup error: %s", err.Error())
	}
	logger.DefaultLogger = logger.DefaultLogger.Fields(map[string]interface{}{
		"x-request-id": "123456abc",
	})
	fmt.Println(logger.DefaultLogger)
	logger.DefaultLogger.Log(level.Info, "hello")
}

func Test_zapToLoggerLevel(t *testing.T) {
	type args struct {
		l zapcore.Level
	}
	tests := []struct {
		name string
		args args
		want level.Level
	}{
		{
			name: "test zapToLoggerLevel",
			args: args{
				l: zapcore.DebugLevel,
			},
			want: level.Debug,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zapToLoggerLevel(tt.args.l); got != tt.want {
				t.Errorf("zapToLoggerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
