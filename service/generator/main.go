package main

import (
	"context"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/pkg/config"

	"generator/cfg"
	"generator/router"
)

// @title generator API
// @version 0.0.1
// @description generator接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:8001
// @BasePath
func main() {
	c := &cfg.Config{}
	err := config.Init(flag.Lookup("c").Value.String(), c)
	if err != nil {
		log.Fatalf("cfg init failed, %s\n", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	r := gin.Default()
	router.Init(r.Group("/generator"))

	c.Init(r)

	log.Info("starting generator manage")

	err = server.Manage.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
