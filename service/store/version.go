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
	flag.String("stage", "local", "stage environment")
	flag.BoolVar(&ver, "v", false, "Print the server version information")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if ver {
		fmt.Println(version)
		os.Exit(0)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
