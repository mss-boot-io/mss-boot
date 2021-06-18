/*
 * @Author: lwnmengjing
 * @Date: 2021/6/18 2:24 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/18 2:24 下午
 */

package test

import (
	"context"
	"testing"
	"time"

	"github.com/lwnmengjing/mss-boot/client"
	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"
	"github.com/sanity-io/litter"
	"google.golang.org/grpc/metadata"
)

func TestTenantCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := client.NewTenantService("127.0.0.1:9091", 10*time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	metadata.NewIncomingContext(ctx, metadata.MD{"requestId": []string{"1234"}})
	resp, err := c.GetClient().Create(ctx, &pb.CreateReq{
		Name:        "test",
		Status:      1,
		System:      1,
		Contact:     "18012959561",
		Description: "tenant test",
		Domains:     []string{"127.0.0.1:8080"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(litter.Sdump(resp))
}

func TestTenantGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := client.NewTenantService("127.0.0.1:9091", 10*time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	metadata.NewIncomingContext(ctx, metadata.MD{"requestId": []string{"1234"}})
	resp, err := c.GetClient().Get(ctx, &pb.GetReq{
		Id: "2ae161549d7f43e9a4a5b77afea75d3b",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(litter.Sdump(resp))
}
