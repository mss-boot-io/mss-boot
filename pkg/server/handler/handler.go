/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 3:13 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 3:13 下午
 */

package handler

import (
	"context"

	"github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server/grpc/interceptors/logging/ctxlog"
	"github.com/mss-boot-io/mss-boot/core/tools/utils"
)

// Handler 基类
type Handler struct {
	ID        string
	RequestID string
	Log       *logger.Helper
}

// Make 构建
func (e *Handler) Make(c context.Context) {
	e.Log = logger.NewHelper(ctxlog.Extract(c))
	e.RequestID = utils.GetRequestID(c)
}

// Make 构建
func Make(c context.Context) *Handler {
	h := &Handler{}
	h.Make(c)
	return h
}
