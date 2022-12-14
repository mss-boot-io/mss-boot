/*
 * @Author: lwnmengjing
 * @Date: 2022/7/18 10:06:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/18 10:06:17
 */

package fs

import (
	"io/fs"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

type Source struct {
	opt source.Options
}

func (s *Source) Open(name string) (fs.File, error) {
	return s.opt.FS.Open(name)
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	return s.opt.FS.ReadFile(name)
}

// New source
func New(options ...source.Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	return s, nil
}
