package controller

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/26 01:22:21
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/26 01:22:21
 */

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"

	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions/gorm"
	mgmActions "github.com/mss-boot-io/mss-boot/pkg/response/actions/mgm"
)

// Simple controller
type Simple struct {
	Base
	options Options
}

// NewSimple new simple
func NewSimple(options ...Option) *Simple {
	s := &Simple{
		options: DefaultOptions(),
	}
	for i := range options {
		options[i](&s.options)
	}
	return s
}

// Path route path
func (e *Simple) Path() string {
	if e.options.model == nil {
		return ""
	}
	return strings.ReplaceAll(strings.ToLower(mgm.CollName(e.options.model)), "_", "-")
}

// Handlers return handlers
func (e *Simple) Handlers() gin.HandlersChain {
	return nil
}

// GetAction get action
func (e *Simple) GetAction(key string) response.Action {
	if action := e.options.getAction(key); action != nil {
		return action
	}
	switch e.options.modelProvider {
	case actions.ModelProviderMgm:
		return e.getActionMgm(key)
	case actions.ModelProviderGorm:
		return e.getActionGorm(key)
	default:
		return nil
	}
}

func (e *Simple) getActionMgm(key string) response.Action {
	b := mgmActions.Base{
		Model: e.options.model,
	}
	switch key {
	case response.Get:
		return mgmActions.NewGet(b, e.GetKey())
	case response.Control:
		return mgmActions.NewControl(b, e.GetKey())
	case response.Delete:
		return mgmActions.NewDelete(b, e.GetKey())
	case response.Search:
		return mgmActions.NewSearch(b, e.options.search)
	default:
		return nil
	}
}

func (e *Simple) getActionGorm(key string) response.Action {
	opts := []gorm.Option{
		gorm.WithModel(e.options.model),
		gorm.WithScope(e.options.scope),
		gorm.WithTreeField(e.options.treeField),
		gorm.WithDepth(e.options.depth),
	}
	if e.options.needAuth(key) {
		opts = append(opts, gorm.WithHandlers(gin.HandlersChain{response.AuthHandler}))
	}
	switch key {
	case response.Get:
		opts = append(opts, gorm.WithKey(e.GetKey()))
		return gorm.NewGet(opts...)
	case response.Control:
		opts = append(opts, gorm.WithKey(e.GetKey()))
		return gorm.NewControl(opts...)
	case response.Delete:
		opts = append(opts, gorm.WithKey(e.GetKey()))
		return gorm.NewDelete(opts...)
	case response.Search:
		opts = append(opts, gorm.WithSearch(e.options.search))
		return gorm.NewSearch(opts...)
	default:
		return nil
	}
}
