// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 574fb16d86
// Version Date: 2019年 4月12日 星期五 00时42分59秒 UTC

package main

import (
	"flag"

	// This Service
	"github.com/akka/gokit-demo/hello/helloservice-service/svc/server"
	"github.com/akka/gokit-demo/hello/helloservice-service/svc/server/cli"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(cli.Config)
}
