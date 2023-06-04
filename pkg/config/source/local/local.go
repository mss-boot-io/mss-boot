/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:25:06
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:25:06
 */

package local

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

// Source is a local file source
type Source struct {
	opt *source.Options
}

// Open a file for reading
func (s *Source) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(s.opt.Dir, name))
}

// ReadFile read file
func (s *Source) ReadFile(name string) (rb []byte, err error) {
	for i := range source.Extends {
		path := filepath.Join(s.opt.Dir,
			fmt.Sprintf("%s.%s", name, source.Extends[i]))
		_, err = os.Stat(path)
		if err != nil {
			continue
		}
		rb, err = os.ReadFile(filepath.Join(s.opt.Dir,
			fmt.Sprintf("%s.%s", name, source.Extends[i])))
		if err == nil {
			s.opt.Extend = source.Extends[i]
			return rb, nil
		}
	}
	return nil, err
}

// GetExtend get extend
func (s *Source) GetExtend() source.Scheme {
	return s.opt.Extend
}

// New returns a new source
func New(options ...source.Option) (*Source, error) {
	s := &Source{
		opt: source.DefaultOptions(),
	}
	for _, opt := range options {
		opt(s.opt)
	}
	return s, nil
}
