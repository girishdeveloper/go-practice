package main

import (
	"config/server"
	"fmt"
	"time"
)

func main() {
	serve := server.Server{}
	defaultConfig := serve.NewServer(serve.DefaultOpts())
	config := serve.NewServer(server.Opts{
		MaxConnection: 300,
		Id:            2,
		Tls:           true,
		Status:        "running",
		UpSince:       time.Now(),
	})
	fmt.Println("default configuration parameters are as follows")
	fmt.Println(defaultConfig)
	fmt.Println("configuration parameters are as follows")
	fmt.Println(config)
}
