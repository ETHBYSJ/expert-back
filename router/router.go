package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	recommendController := controller.RecommendController{}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		recommend := v1.Group("/recommend")
		{
			// 提交推荐信息
			recommend.GET("/download", recommendController.RecommendDownload)
			recommend.POST("/commit", recommendController.RecommendCommit)
			recommend.POST("/upload", recommendController.RecommendUpload)
			// recommend.GET("/testParse", recommendController.TestParse)
		}
	}

	return r
}