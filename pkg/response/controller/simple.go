/*
 * @Author: lwnmengjing
 * @Date: 2023/1/26 01:22:21
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/26 01:22:21
 */

package controller

import (
	"strings"

	"github.com/kamva/mgm/v3"

	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
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

// GetAction get action
func (e *Simple) GetAction(key string) response.Action {
	if action := e.options.getAction(key); action != nil {
		return action
	}
	// default action
	switch key {
	case response.Get:
		return actions.NewGet(e.options.model, e.GetKey())
	case response.Control:
		return actions.NewControl(e.options.model, e.GetKey())
	case response.Delete:
		return actions.NewDelete(e.options.model, e.GetKey())
	case response.Search:
		return actions.NewSearch(e.options.model, e.options.search)
	default:
		return nil
	}
}
