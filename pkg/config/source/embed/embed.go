/*
 * @Author: lwnmengjing
 * @Date: 2022/7/18 10:06:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/18 10:06:17
 */

package embed

import (
	"embed"
	"io/fs"
)

type Source struct {
	from embed.FS
}

func (s *Source) Open(name string) (fs.File, error) {
	return s.from.Open(name)
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	return s.from.ReadFile(name)
}

// New source
func New(source embed.FS) (*Source, error) {
	return &Source{
		from: source,
	}, nil
}
