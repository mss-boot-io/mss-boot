package main

import (
	"context"
	"flag"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	pb "github.com/mss-boot-io/mss-boot/proto/store/v1"

	"github.com/mss-boot-io/mss-boot/service/store/cfg"
	"github.com/mss-boot-io/mss-boot/service/store/handlers"
)

func main() {
	c := &cfg.Config{}
	err := config.Init(flag.Lookup("c").Value.String(), c)
	if err != nil {
		log.Fatalf("cfg init failed, %s\n", err.Error())
	}
	//ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()

	c.Init(func(srv *grpc.Server) {
		pb.RegisterStoreServer(srv.Server(), handlers.NewStoreHandler("store"))
	})

	log.Info("starting generator manage")

	err = server.Manage.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
