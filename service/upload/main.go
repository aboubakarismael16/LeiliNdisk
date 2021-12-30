package main

import (
	cfg "LeiliNetdisk/config"
	"LeiliNetdisk/route"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func startApiService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {

	pwd, _ := os.Getwd()
	fmt.Println(pwd + " " + os.Args[0])
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "./"))))

	router := route.Router()
	router.Run(cfg.UploadServiceHost)

}
