package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	cfg "LeiliNetdisk/config"
	"LeiliNetdisk/handler"
)

func main() {

	// 静态资源处理
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	pwd, _ := os.Getwd()
	fmt.Println(pwd + " " + os.Args[0])
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "./"))))

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta/", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	http.HandleFunc("/file/mpupload/uppart", handler.UploadPartHandler)
	http.HandleFunc("/file/mpupload/complete", handler.CompleteUploadHandler)

	// 用户相关接口
	http.HandleFunc("/", handler.SignInHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))


	fmt.Printf("上传服务启动中，开始监听监听[%s]...\n", cfg.UploadServiceHost)
	err := http.ListenAndServe(cfg.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
