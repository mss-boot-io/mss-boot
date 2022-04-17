package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"log"

	"github.com/mss-boot-io/mss-boot/service/tenant/cfg"
	"github.com/mss-boot-io/mss-boot/service/tenant/router"
)

// @title tenant API
// @version 0.0.1
// @description tenant接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:9094
// @BasePath
func main() {
	c := &cfg.Config{}
	err := config.Init(flag.Lookup("c").Value.String(), c)
	if err != nil {
		log.Printf("cfg init failed, %s\n", err.Error())
		return
	}
	ctx := context.Background()

	r := gin.Default()
	router.Init(r.Group("/tenant"))

	c.Init(r)

	log.Println("starting tenant manage")

	err = server.Manage.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
