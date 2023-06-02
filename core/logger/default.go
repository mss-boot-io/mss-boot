package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/mss-boot-io/mss-boot/core/logger/formatter"
	"github.com/mss-boot-io/mss-boot/core/logger/level"
)

func init() {
	lvl, err := level.GetLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = level.Info
	}

	DefaultLogger = NewHelper(NewLogger(WithLevel(lvl)))
}

type defaultLogger struct {
	sync.RWMutex
	opts Options
	log  *log.Logger
}

// Init (opts...) should only overwrite provided Options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	l.log = log.New(l.opts.Out, "", log.Lshortfile|log.LstdFlags)
	//l.log = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	return nil
}

// String string
func (l *defaultLogger) String() string {
	return "default"
}

// Fields Fields
func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	l.opts.Fields = copyFields(fields)
	l.Unlock()
	return l
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// logCallerFilePath returns a package/file:line description of the caller,
// preserving only the leaf directory Name and file Name.
func logCallerFilePath(loggingFilePath string) string {
	// To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	idx := strings.LastIndexByte(loggingFilePath, filepath.Separator)
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], filepath.Separator)
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

// Log log
func (l *defaultLogger) Log(level level.Level, v ...interface{}) {
	l.logf(level, "", v...)
}

// Logf logf
func (l *defaultLogger) Logf(level level.Level, format string, v ...interface{}) {
	l.logf(level, format, v...)
}

func (l *defaultLogger) logf(level level.Level, format string, v ...interface{}) {
	// TODO decide does we need to write message if log Level.Level not used?
	if !l.opts.Level.Enabled(level) {
		return
	}

	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["Level"] = level.String()

	//if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
	//	fields["file"] = fmt.Sprintf("%s:%d", logCallerFilePath(file), line)
	//}

	rec := formatter.Record{
		//Timestamp: time.Now(),
		Metadata: make(map[string]string, len(fields)),
	}
	if format == "" {
		rec.Message = fmt.Sprint(v...)
	} else {
		rec.Message = fmt.Sprintf(format, v...)
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for i, k := range keys {
		if i == 0 {
			metadata += fmt.Sprintf("%s:%v", k, fields[k])
		} else {
			metadata += fmt.Sprintf(" %s:%v", k, fields[k])
		}
	}

	var name string
	if l.opts.Name != "" {
		name = "[" + l.opts.Name + "]"
	}
	//t := rec.Timestamp.Format("2006-01-02 15:04:05.000Z0700")
	//fmt.Printf("%s\n", t)
	//fmt.Printf("%s\n", Name)
	//fmt.Printf("%s\n", metadata)
	//fmt.Printf("%v\n", rec.Message)
	logStr := ""
	if name == "" {
		logStr = fmt.Sprintf("%s %v\n", metadata, rec.Message)
	} else {
		logStr = fmt.Sprintf("%s %s %v\n", name, metadata, rec.Message)
	}
	err := l.log.Output(l.opts.CallerSkipCount+1, logStr)

	//_, err := l.opts.Out.Write([]byte(logStr))
	if err != nil {
		log.Printf("log [Logf] write error: %s \n", err.Error())
	}

}

// Options get Options
func (l *defaultLogger) Options() Options {
	// not guard against Options Context values
	l.RLock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	l.RUnlock()
	return opts
}

// NewLogger builds a new logger based on Options
func NewLogger(opts ...Option) Logger {
	// Default Options
	options := Options{
		Level:           level.Info,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 3,
		Context:         context.Background(),
		Name:            "",
	}

	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(level.Fatal, err)
	}

	return l
}
