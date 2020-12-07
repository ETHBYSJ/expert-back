package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	fileController := controller.FileController{}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		file := v1.Group("/file")
		{
			// 下载文件
			file.GET("/download", fileController.DownloadFile)
		}
	}

	return r
}