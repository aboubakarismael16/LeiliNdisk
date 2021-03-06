package route

import (
	"LeiliNetdisk/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//gin framework 包括Logger,Recovery
	router := gin.Default()
	//处理静态资源
	router.Static("/static/", "./static")

	// 无需验证就能访问的接口
	router.GET("/user/signup",handler.SignupHandler)
	router.POST("/user/signup",handler.DoSignupHandler)
	router.GET("/user/signin",handler.SignInHandler)
	router.POST("/user/signin",handler.DoSignInHandler)
	router.GET("/", handler.SignInHandler)


	//加入中间件，用于效验token的拦截器
	router.Use(handler.HTTPInterceptor())

	// Use 中间件之后的接口，都需要通过拦截器
	// 用户信息
	router.POST("/user/info",handler.UserInfoHandler)


	//use之后的所有handler都会进去拦截

	////用户信息查询
	//router.GET("/user/info", handler.UserInfoHandler)

	//文件存取接口
	router.GET("/file/upload", handler.UploadHandler)
	router.POST("/file/upload", handler.DoUploadHandler)
	//router.GET("/file/query", handler.QueryFileHandler)
	//TODO:功能正常，与前端交互有问题
	//router.POST("/file/update", handler.FileMetaUpdataHandle)
	//TODO:功能正常，没有前端按钮
	router.POST("/file/delete", handler.FileDeleteHandler)
	//尚未测试
	router.POST("/file/download", handler.DownloadHandler)
	router.POST("/file/downloadurl", handler.DownloadURLHandler)

	//秒传接口
	router.POST("/file/fastupload", handler.TryFastUploadHandler)

	//分块上传接口
	router.POST("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	router.POST("/file/mpupload/uppart", handler.UploadPartHandler)
	router.POST("/file/mpupload/finshed", handler.CompleteUploadHandler)
	router.POST("/file/mpupload/cancel", handler.CancelUploadPartHandler)
	router.POST("/file/mpupload/status", handler.MultipartUploadStatusHanlder)

	return router
}
