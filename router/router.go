package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
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
			recommend.GET("/download", recommendController.RecommendDownload)
			recommend.POST("/commit", recommendController.RecommendSubmit)
			recommend.POST("/upload", recommendController.RecommendUpload)
			recommend.GET("/experts", recommendController.RecommendGet)
		}
		apply := v1.Group("/apply")
		{
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

	return r
}