/*
 * @Author: lwnmengjing
 * @Date: 2021/6/21 11:46 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/21 11:46 上午
 */

package client

// Options options
type Options struct {
	store *StoreService
	//tenant *TenantService
}

// Option set options
type Option func(*Options)

// WithStore set store client
func WithStore(c *StoreService) Option {
	return func(o *Options) {
		o.store = c
	}
}

// WithTenant set tenant client
//func WithTenant(c *TenantService) Option {
//	return func(o *Options) {
//		o.tenant = c
//	}
//}

// SetDefault set default balance
func SetDefault() Options {
	return Options{}
}
