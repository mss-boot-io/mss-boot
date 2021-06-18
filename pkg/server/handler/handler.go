/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 3:13 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 3:13 下午
 */

package handler

import (
	"context"

	log "github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/core-go/server/grpc/interceptors/logging/ctxlog"
)

// Handler 基类
type Handler struct {
	ID  string
	Log *log.Helper
}

// Make 构建
func (e *Handler) Make(c context.Context) {
	e.Log = ctxlog.Extract(c)
}
