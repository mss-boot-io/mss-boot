package response

/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 20:05:14
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 20:05:14
 */

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/virtual/model"
)

const (
	// Get action
	Get = "get"
	// Base action
	Base = "base"
	// Delete action
	Delete = "delete"
	// Search action
	Search = "search"
	// Control action
	Control = "control"
	// Create action
	Create = "create"
	// Update action
	Update = "update"
)

// Action interface
type Action interface {
	fmt.Stringer
	Handler() gin.HandlerFunc
}

// Searcher search interface
type Searcher interface {
	GetPage() int64
	GetPageSize() int64
}

// VirtualAction virtual action
type VirtualAction interface {
	String() string
	Handler() gin.HandlerFunc
	SetModel(key string, m *model.Model)
	GetModel(ctx *gin.Context) *model.Model
}
