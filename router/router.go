package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	recommendController := controller.RecommendController{}
	commonController := controller.CommonController{}
	applyController := controller.ApplyController{}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/cookie", commonController.SetCookie)
		v1.GET("/profile", commonController.GetAccountProfile)
		recommend := v1.Group("/recommend")
		{
			// 提交推荐信息
			recommend.POST("/commit", recommendController.RecommendSubmit)
			// 上传推荐表
			recommend.POST("/upload", recommendController.RecommendUpload)
			// 下载推荐表
			recommend.GET("/download", recommendController.RecommendDownload)
			// 获取推荐专家信息
			recommend.GET("/experts", recommendController.RecommendGet)
		}
		apply := v1.Group("/apply")
		{
			// 上传申请表

			// 下载申请表

			// 上传照片
			apply.POST("/uploadPhoto", applyController.ApplyUpload)
			// 提交申请信息
			apply.GET("/create", applyController.ApplyCreate)
			// 提交基本信息
			apply.POST("/submitBase", applyController.ApplySubmitBase)
			// 提交专业类别
			apply.POST("/submitMajor", applyController.ApplySubmitMajor)
			// 提交专攻领域
			apply.POST("/submitResearchField", applyController.ApplySubmitResearchField)
			// 提交个人履历
			apply.POST("/submitResume", applyController.ApplySubmitResume)
			// 提交意见评价
			apply.POST("/submitOpinion", applyController.ApplySubmitOpinion)
		}
	}
	r.StaticFS("/static", http.Dir("./static/upload/picture"))
	return r
}