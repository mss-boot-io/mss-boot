/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 2:33 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 2:33 下午
 */

package client

import (
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
	pb "github.com/mss-boot-io/mss-boot/proto/store/v1"
)

var _store *StoreService

// StoreString store string
const StoreString = "store"

// GetStore get store client
func GetStore() *StoreService {
	return _store
}

// SetStore set store client
func SetStore(c *StoreService) {
	_store = c
}

// StoreService store service client
type StoreService struct {
	grpc.Service
	client pb.StoreClient
}

// String string
func (StoreService) String() string {
	return StoreString
}

// NewStore new a store client
func NewStore(c grpc.Service) *StoreService {
	s := &StoreService{
		Service: c,
	}
	s.client = pb.NewStoreClient(s.Connection)
	return s
}

// GetClient get client
func (e *StoreService) GetClient() pb.StoreClient {
	return e.client
}
