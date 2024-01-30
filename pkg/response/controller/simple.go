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
	b := gorm.Base{
		Model:     e.options.model,
		Scope:     e.options.scope,
		TreeField: e.options.treeField,
		Depth:     e.options.depth,
	}
	if e.options.needAuth(key) {
		b.Handlers = gin.HandlersChain{response.AuthHandler}
	}
	switch key {
	case response.Get:
		return gorm.NewGet(b, e.GetKey())
	case response.Control:
		return gorm.NewControl(b, e.GetKey())
	case response.Delete:
		return gorm.NewDelete(b, e.GetKey())
	case response.Search:
		return gorm.NewSearch(b, e.options.search)
	default:
		return nil
	}
}
