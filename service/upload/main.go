package main

import (
	cfg "LeiliNetdisk/config"
	"LeiliNetdisk/route"
)

func startApiService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	// api 服务
	go startApiService()

	// rpc 服务
	//startRpcService()
}
