package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions/virtual"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/18 09:09:44
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/18 09:09:44
 */

// Virtual controller
type Virtual struct {
	Base
	options  Options
	Provider *virtual.Base
}

// NewVirtual new virtual controller
func NewVirtual(provider *virtual.Base, options ...Option) *Virtual {
	v := &Virtual{
		Provider: provider,
		options:  DefaultOptions(),
	}
	for i := range options {
		options[i](&v.options)
	}
	return v
}

// Path http path
func (v *Virtual) Path() string {
	return ":" + virtual.PathKey
}

// Handlers middlewares
func (v *Virtual) Handlers() gin.HandlersChain {
	if v.options.auth {
		return gin.HandlersChain{response.AuthHandler}
	}
	return nil
}

// GetAction get action
func (v *Virtual) GetAction(key string) response.Action {
	return v.getAction(key)
}

func (v *Virtual) getAction(key string) response.Action {
	switch key {
	case response.Get:
		return virtual.NewGet(v.Provider)
	case response.Create:
		return virtual.NewCreate(v.Provider)
	case response.Update:
		return virtual.NewUpdate(v.Provider)
	case response.Delete:
		return virtual.NewDelete(v.Provider)
	case response.Search:
		return virtual.NewSearch(v.Provider)
	default:
		return nil
	}
}
