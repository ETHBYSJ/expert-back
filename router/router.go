package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	fileController := controller.FileController{}
	recommendController := controller.RecommendController{}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		file := v1.Group("/file")
		{
			// 下载文件
			file.GET("/download", fileController.DownloadFile)
		}
		recommend := v1.Group("/recommend")
		{
			// 提交推荐信息
			recommend.POST("/commit", recommendController.Recommend)
		}
	}

	return r
}