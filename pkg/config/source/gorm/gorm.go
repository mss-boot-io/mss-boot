/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/11 01:22:28
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/11 01:22:28
 */

package gorm

import (
	"errors"
	"io/fs"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

// Source source
type Source struct {
	opt *source.Options
}

// Open method Get not implemented
func (s *Source) Open(string) (fs.File, error) {
	return nil, errors.New("method Get not implemented")
}

// ReadFile method Get not implemented
func (s *Source) ReadFile(name string) ([]byte, error) {
	return nil, errors.New("method Get not implemented")
}
