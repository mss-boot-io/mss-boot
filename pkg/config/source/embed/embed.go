/*
 * @Author: lwnmengjing
 * @Date: 2022/7/18 10:06:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/18 10:06:17
 */

package embed

import (
	"io/fs"
)

type Source struct {
	opt Options
}

func (s *Source) Open(name string) (fs.File, error) {
	return s.opt.fs.Open(name)
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	return s.opt.fs.ReadFile(name)
}

// New source
func New(options ...Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	return s, nil
}
