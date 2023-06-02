package ctxlog

/*
 * @Author: lwnmengjing
 * @Date: 2021/5/19 11:42 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/19 11:42 上午
 */

import "sync"

// Fields fields
type Fields struct {
	value sync.Map
}

// NewFields make a new fields
func NewFields(key string, value interface{}) *Fields {
	f := &Fields{}
	f.Set(key, value)
	return f
}

// Set field
func (e *Fields) Set(key string, value interface{}) {
	e.value.Store(key, value)
}

// Values return value
func (e *Fields) Values() map[string]interface{} {
	result := make(map[string]interface{})
	e.value.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})
	return result
}

// Merge merge fields
func (e *Fields) Merge(f *Fields) {
	for k, v := range f.Values() {
		e.Set(k, v)
	}
}
