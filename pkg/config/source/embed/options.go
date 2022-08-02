/*
 * @Author: lwnmengjing
 * @Date: 2022/7/22 02:11:50
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/22 02:11:50
 */

package embed

import "embed"

// Option set Options
type Option func(*Options)

type Options struct {
	fs embed.FS
}

// WithFrom set embed.FS
func WithFrom(fs embed.FS) Option {
	return func(args *Options) {
		args.fs = fs
	}
}
