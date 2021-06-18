/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 2:33 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 2:33 下午
 */

package client

import (
	"time"

	"github.com/lwnmengjing/core-go/server/grpc"
	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"
)

type TenantService struct {
	grpc.Service
	client pb.TenantClient
}

func NewTenantService(endpoint string, callTimeout time.Duration) (*TenantService, error) {
	s := &TenantService{
		Service: grpc.Service{
			CallTimeout: callTimeout,
		},
	}
	if err := s.Dial(endpoint, callTimeout); err != nil {
		return nil, err
	}
	s.client = pb.NewTenantClient(s.Connection)
	return s, nil
}

func (e *TenantService) GetClient() pb.TenantClient {
	return e.client
}
