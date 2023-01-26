/*
 * @Author: lwnmengjing
 * @Date: 2023/1/25 20:05:14
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/25 20:05:14
 */

package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	Get     = "get"
	Base    = "base"
	Delete  = "delete"
	Search  = "search"
	Control = "control"
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
