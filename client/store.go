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
	pb "github.com/lwnmengjing/mss-boot/proto/store/v1"
)

type StoreService struct {
	grpc.Service
	client pb.StoreClient
}

func NewStoreService(endpoint string, callTimeout time.Duration) (*StoreService, error) {
	s := &StoreService{
		Service: grpc.Service{
			CallTimeout: callTimeout,
		},
	}
	if err := s.Dial(endpoint, callTimeout); err != nil {
		return nil, err
	}
	s.client = pb.NewStoreClient(s.Connection)
	return s, nil
}

func (e *StoreService) GetClient() pb.StoreClient {
	return e.client
}
