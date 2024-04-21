package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/4/20 22:46:00
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/4/20 22:46:00
 */

type ActionHook func(ctx *gin.Context, db *gorm.DB, m schema.Tabler) error

type Option func(*Options)

type Options struct {
	Model           schema.Tabler
	Scope           func(ctx *gin.Context, table schema.Tabler) func(db *gorm.DB) *gorm.DB
	Handlers        gin.HandlersChain
	TreeField       string
	Depth           int
	Key             string
	Search          response.Searcher
	BeforeCreate    ActionHook
	AfterCreate     ActionHook
	BeforeUpdate    ActionHook
	AfterUpdate     ActionHook
	BeforeGet       ActionHook
	AfterGet        ActionHook
	BeforeDelete    ActionHook
	AfterDelete     ActionHook
	BeforeSearch    ActionHook
	AfterSearch     ActionHook
	handlers        gin.HandlersChain
	controlHandlers gin.HandlersChain
	getHandlers     gin.HandlersChain
	deleteHandlers  gin.HandlersChain
	searchHandlers  gin.HandlersChain
}

func WithModel(m schema.Tabler) Option {
	return func(o *Options) {
		o.Model = m
	}
}

func WithScope(scope func(ctx *gin.Context, table schema.Tabler) func(db *gorm.DB) *gorm.DB) Option {
	return func(o *Options) {
		o.Scope = scope
	}
}

func WithHandlers(handlers gin.HandlersChain) Option {
	return func(o *Options) {
		o.Handlers = handlers
	}
}

func WithTreeField(treeField string) Option {
	return func(o *Options) {
		o.TreeField = treeField
	}
}

func WithDepth(depth int) Option {
	return func(o *Options) {
		o.Depth = depth
	}
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

func WithSearch(search response.Searcher) Option {
	return func(o *Options) {
		o.Search = search
	}
}

func WithBeforeCreate(hook ActionHook) Option {
	return func(o *Options) {
		o.BeforeCreate = hook
	}
}

func WithAfterCreate(hook ActionHook) Option {
	return func(o *Options) {
		o.AfterCreate = hook
	}
}

func WithBeforeUpdate(hook ActionHook) Option {
	return func(o *Options) {
		o.BeforeUpdate = hook
	}
}

func WithAfterUpdate(hook ActionHook) Option {
	return func(o *Options) {
		o.AfterUpdate = hook
	}
}

func WithBeforeGet(hook ActionHook) Option {
	return func(o *Options) {
		o.BeforeGet = hook
	}
}

func WithAfterGet(hook ActionHook) Option {
	return func(o *Options) {
		o.AfterGet = hook
	}
}

func WithBeforeDelete(hook ActionHook) Option {
	return func(o *Options) {
		o.BeforeDelete = hook
	}
}

func WithAfterDelete(hook ActionHook) Option {
	return func(o *Options) {
		o.AfterDelete = hook
	}
}

func WithBeforeSearch(hook ActionHook) Option {
	return func(o *Options) {
		o.BeforeSearch = hook
	}
}

func WithAfterSearch(hook ActionHook) Option {
	return func(o *Options) {
		o.AfterSearch = hook
	}
}

func WithControlHandlers(handlers gin.HandlersChain) Option {
	return func(o *Options) {
		o.controlHandlers = handlers
	}
}

func WithGetHandlers(handlers gin.HandlersChain) Option {
	return func(o *Options) {
		o.getHandlers = handlers
	}
}

func WithDeleteHandlers(handlers gin.HandlersChain) Option {
	return func(o *Options) {
		o.deleteHandlers = handlers
	}
}

func WithSearchHandlers(handlers gin.HandlersChain) Option {
	return func(o *Options) {
		o.searchHandlers = handlers
	}
}
