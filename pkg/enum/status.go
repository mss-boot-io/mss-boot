package enum

/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:10
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:10
 */

//go:generate stringer -type Status -output status_string.go

// Status type for enum
type Status uint8

const (
	_ Status = iota
	// Enabled enabled
	Enabled
	// Disabled disabled
	Disabled
	// Locked lock status
	Locked
)
