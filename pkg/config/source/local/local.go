/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:25:06
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:25:06
 */

package local

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Source struct {
	opt Options
}

func (s *Source) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(s.opt.dir, name))
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(s.opt.dir, name))
}

func New(options ...Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	return s, nil
}
