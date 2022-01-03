package main

import (
	"github.com/micro/go-micro/registry"
	"log"
	"time"

	micro "github.com/micro/go-micro"

	"LeiliNetdisk/service/account/handler"
	proto "LeiliNetdisk/service/account/proto"
	_ "github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {

	registry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*10),     //10s检查等待时间
		micro.RegisterInterval(time.Second*5), //服务每隔5秒发一次心跳
		//micro.Flags(common.CustomFlags...),
	)
	//初始化service,解析命令行参数等
	service.Init()

	//todo:初始化dbproxy client
	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

