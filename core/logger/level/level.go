/*
 * @Author: lwnmengjing
 * @Date: 2021/6/15 5:58 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/15 5:58 下午
 */

package level

import (
	"fmt"
	"strings"
)

//go:generate stringer -type level -output level_string.go

// Level level for logger
type Level int8

const (
	// Trace level. Designates finer-grained informational events than the Debug.
	Trace Level = iota - 2
	// Debug level. Usually only enabled when debugging. Very verbose logging.
	Debug
	// Info is the default logging priority.
	// General operational entries about what's going on inside the application.
	Info
	// Warn level. Non-critical entries that deserve eyes.
	Warn
	// Error level. Logs. Used for errors that should definitely be noted.
	Error
	// Fatal level. Logs and then calls `logger.Exit(1)`. highest level of severity.
	Fatal
)

// ToGorm trans to gorm log level
func (l Level) ToGorm() int {
	switch l {
	case Fatal, Error:
		return 2
	case Warn:
		return 3
	case Info, Debug, Trace:
		return 4
	default:
		return 1
	}
}

// Enabled returns true if the given level is at or above this level.
func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

// GetLevel converts a level string into a logger Level value.
// returns an error if the input string does not match known values.
func GetLevel(levelStr string) (Level, error) {
	switch strings.ToLower(levelStr) {
	case strings.ToLower(Trace.String()):
		return Trace, nil
	case strings.ToLower(Debug.String()):
		return Debug, nil
	case strings.ToLower(Info.String()):
		return Info, nil
	case strings.ToLower(Warn.String()):
		return Warn, nil
	case strings.ToLower(Error.String()):
		return Error, nil
	case strings.ToLower(Fatal.String()):
		return Fatal, nil
	}
	return Info, fmt.Errorf("unknown level String: '%s', defaulting to InfoLevel", levelStr)
}
