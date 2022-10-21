/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:31:13
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:31:13
 */

package local

type Option func(*Options)

type Options struct {
	dir string
}

// WithDir set local dir
func WithDir(dir string) Option {
	return func(args *Options) {
		args.dir = dir
	}
}
