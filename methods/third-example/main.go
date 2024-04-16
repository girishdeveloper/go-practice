package main

import (
	"conf/server"
	"fmt"
	"time"
)

func main() {
	serve := server.Server{}
	defaultConf := serve.NewServer(serve.DefaultOpts())
	conf := serve.NewServer(server.Opts{
		ID:            2,
		MaxConnection: 500,
		Tls:           true,
		Status:        "paused",
		UpSince:       time.Now(),
	})
	fmt.Println("Default configuration is as follows")
	fmt.Println(defaultConf)
	fmt.Println("New configuration is as follows")
	fmt.Println(conf)
}
