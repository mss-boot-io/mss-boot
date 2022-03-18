/*
 * @Author: lwnmengjing
 * @Date: 2022/3/19 1:24
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/19 1:24
 */

package handlers

import (
	"context"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/server/handler"
	pb "github.com/mss-boot-io/mss-boot/proto/store/v1"
)

// StoreHandler store handler
type StoreHandler struct {
	handler.Handler
	pb.UnimplementedStoreServer
}

// NewStoreHandler new store handler
func NewStoreHandler(id string) *StoreHandler {
	return &StoreHandler{
		Handler: handler.Handler{
			ID:  id,
			Log: log.NewHelper(log.DefaultLogger),
		},
	}
}

func (e *StoreHandler) Get(c context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	e.Make(c)
	e.Log.Info(req)
	resp := &pb.GetResp{Value: "123"}
	return resp, nil
}
