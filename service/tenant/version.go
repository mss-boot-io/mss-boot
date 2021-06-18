/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 2:56 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 2:56 下午
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const version = "0.0.1"

var ver bool

func init() {
	flag.String("c", "config/store-service.yaml",
		"Read configuration from specified `FILE`, supports JSON/YAML/TOML formats")
	flag.BoolVar(&ver, "v", false, "Print the server version information")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if ver {
		fmt.Println(version)
		os.Exit(-1)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
