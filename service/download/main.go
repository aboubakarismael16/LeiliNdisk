package main

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"time"

	micro "github.com/micro/go-micro"

	cfg "LeiliNetdisk/service/download/config"
	dlProto "LeiliNetdisk/service/download/proto"
	"LeiliNetdisk/service/download/route"
	dlRpc "LeiliNetdisk/service/download/rpc"
)

func startRpcService() {

	registry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.download"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Registry(registry),
	)
	service.Init()

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startApiService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startApiService()

	// rpc 服务
	startRpcService()
}
