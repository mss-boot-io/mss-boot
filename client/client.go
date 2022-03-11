/*
 * @Author: lwnmengjing
 * @Date: 2021/6/22 3:24 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/22 3:24 下午
 */

package client

// Default default client
var Default = &Client{}

// Client client
type Client struct {
	opts Options
}

// Init init
func (e *Client) Init(opts ...Option) {
	e.opts = SetDefault()
	for i := range opts {
		opts[i](&e.opts)
	}
}

// Store get store client
func (e *Client) Store() *StoreService {
	return e.opts.store
}

// Tenant get tenant client
//func (e *Client) Tenant() *TenantService {
//	return e.opts.tenant
//}

// Store client
func Store() *StoreService {
	return Default.opts.store
}

// Tenant tenant client
//func Tenant() *TenantService {
//	return Default.opts.tenant
//}
