package main

import (
	"LeiliNetdisk/common"
	"LeiliNetdisk/config_example"
	"LeiliNetdisk/mq"
	dbproxy "LeiliNetdisk/service/dbproxy/client"
	"LeiliNetdisk/service/transfert/process"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func main() {
	// 文件转移服务
	go startTransferService()

	// rpc 服务
	startRPCService()
}

func startRPCService() {

	registry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.transfer"),
		micro.RegisterTTL(time.Second*10),
		micro.Registry(registry),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...))

	service.Init(
		micro.Action(func(context *cli.Context) {
			//检查是否有指定的mqHost
			mqhost := context.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				mq.UpdateRabbitHost(mqhost)
			}
		}),
	)

	//初始化dbproxy client
	dbproxy.Init(service)
	//初始化mq client
	mq.Init()

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func startTransferService() {
	if !config_example.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")
	mq.StartConsume(
		config_example.TransOSSQueueName,
		"transfer_oss", process.Transfer)
}
