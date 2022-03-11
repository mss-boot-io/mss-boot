/*
 * @Author: lwnmengjing
 * @Date: 2021/6/23 11:22 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/23 11:22 上午
 */

package main

import (
	"context"
	"flag"
	"log"
	"oauth2/manage"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/pkg/config"

	c "oauth2/cfg"
	"oauth2/router"
)

// @title oauth2 API
// @version 0.0.1
// @description oauth2接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:9096
// @BasePath /oauth2
func main() {
	f := flag.Lookup("c")
	cfg := &c.Config{}
	err := config.Init(f.Value.String(), cfg)
	if err != nil {
		log.Printf("cfg init failed, %s\n", err.Error())
		return
	}
	ctx := context.Background()

	mgr := server.New()

	r := gin.Default()
	router.Init(r.Group("/oauth2"))

	err = cfg.Init(mgr, r)
	if err != nil {
		log.Printf("cfg enity init failed, %s\n", err.Error())
		return
	}
	manage.Init()

	log.Println("starting oauth2 manage")

	err = mgr.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
