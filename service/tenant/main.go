/*
 * @Author: lwnmengjing
 * @Date: 2021/5/18 12:58 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/5/18 12:58 下午
 */

package main

import (
	"context"
	"flag"
	"log"
	"tenant/models"

	"github.com/lwnmengjing/core-go/server"
	"github.com/lwnmengjing/core-go/server/grpc"
	"github.com/lwnmengjing/mss-boot/pkg/config"
	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"

	configStore "tenant/config"
	"tenant/handler"
)

func main() {
	f := flag.Lookup("c")
	cfg := &configStore.Config{}
	err := config.Init(f.Value.String(), cfg)
	if err != nil {
		log.Printf("config init failed, %s\n", err.Error())
		return
	}
	ctx := context.Background()

	mgr := server.New()
	cfg.Init(mgr, func(srv *grpc.Server) {
		pb.RegisterTenantServer(srv.Server(), handler.NewTenant(cfg.Server.Name))
	})
	models.Orm.Debug().Migrator().AutoMigrate(&models.Tenant{})
	log.Println("starting tenant manager")
	err = mgr.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
