package controller

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/26 01:22:21
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/26 01:22:21
 */

import (
	"github.com/mss-boot-io/mss-boot/pkg/response/actions/authentic"
	"strings"

	"github.com/gin-gonic/gin"
	mgm "github.com/kamva/mgm/v3"

	"github.com/mss-boot-io/mss-boot/pkg/response"
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
	return strings.ReplaceAll(strings.ToLower(mgm.CollName(e.options.model)), "_", "-")
}

// Handlers return handlers
func (e *Simple) Handlers() gin.HandlersChain {
	if e.options.auth {
		return gin.HandlersChain{response.AuthHandler}
	}
	return nil
}

// GetAction get action
func (e *Simple) GetAction(key string) response.Action {
	if action := e.options.getAction(key); action != nil {
		return action
	}
	switch e.options.modelProvider {
	case authentic.ModelProviderMgm:
		return e.getActionMgm(key)
	case authentic.ModelProviderGorm:
		return e.getActionGorm(key)
	default:
		return nil
	}
}

func (e *Simple) getActionMgm(key string) response.Action {
	switch key {
	case response.Get:
		return authentic.NewGetMgm(e.options.model, e.GetKey())
	case response.Control:
		return authentic.NewControlMgm(e.options.model, e.GetKey())
	case response.Delete:
		return authentic.NewDeleteMgm(e.options.model, e.GetKey())
	case response.Search:
		return authentic.NewSearchMgm(e.options.model, e.options.search)
	default:
		return nil
	}
}

func (e *Simple) getActionGorm(key string) response.Action {
	switch key {
	case response.Get:
		return authentic.NewGetGorm(e.options.model, e.GetKey())
	case response.Control:
		return authentic.NewControlGorm(e.options.model, e.GetKey())
	case response.Delete:
		return authentic.NewDeleteGorm(e.options.model, e.GetKey())
	case response.Search:
		return authentic.NewSearchGorm(e.options.model, e.options.search)
	default:
		return nil
	}
}
