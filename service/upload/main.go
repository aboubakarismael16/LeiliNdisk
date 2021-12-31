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

	router := route.Router()
	router.Run(cfg.UploadServiceHost)

}
