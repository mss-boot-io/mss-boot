/*
 * @Author: lwnmengjing
 * @Date: 2021/5/19 11:42 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/19 11:42 上午
 */

package ctxlog

import "sync"

// Fields fields
type Fields struct {
	mux   sync.Mutex
	value map[string]interface{}
}

// NewFields make a new fields
func NewFields(key string, value interface{}) *Fields {
	f := &Fields{}
	f.Set(key, value)
	return f
}

// Set field
func (e *Fields) Set(key string, value interface{}) {
	if e.value == nil {
		e.value = make(map[string]interface{})
	}
	e.mux.Lock()
	e.value[key] = value
	e.mux.Unlock()
}

// Values return value
func (e *Fields) Values() map[string]interface{} {
	return e.value
}

// Merge merge fields
func (e *Fields) Merge(f *Fields) {
	if len(f.value) > 0 {
		for k, v := range f.value {
			e.Set(k, v)
		}
	}
}
