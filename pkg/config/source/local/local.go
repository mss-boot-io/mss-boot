/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:25:06
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:25:06
 */

package local

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

type Source struct {
	opt source.Options
}

func (s *Source) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(s.opt.Dir, name))
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.opt.Dir, name))
}

func New(options ...source.Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	return s, nil
}
