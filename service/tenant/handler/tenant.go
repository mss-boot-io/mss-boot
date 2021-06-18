/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 5:25 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 5:25 下午
 */

package handler

import (
	"context"

	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"

	"tenant/models"
	"tenant/service"
)

func (e Tenant) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	e.Make(ctx)
	s := &service.Tenant{}
	s.Make(models.Orm.WithContext(ctx), e.Log)
	err := s.Create(req)
	resp := &pb.CreateResp{}
	return resp, err
}

func (e Tenant) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	e.Make(ctx)
	s := &service.Tenant{}
	s.Make(models.Orm.WithContext(ctx), e.Log)
	resp := &pb.GetResp{}
	err := s.Get(req, resp)
	return resp, err
}
