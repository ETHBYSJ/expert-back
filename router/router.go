package router

import (
	"expert-back/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	searchController := controller.SearchController{}
	recommendController := controller.RecommendController{}
	commonController := controller.CommonController{}
	applyController := controller.ApplyController{}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		// 设置cookie
		v1.GET("/cookie", commonController.SetCookie)
		// 获取用户信息
		v1.GET("/profile", commonController.GetAccountProfile)
		// 审核专家推荐
		v1.POST("/reviewRecommend", commonController.ReviewRecommend)
		// 审核专家申请
		v1.POST("/reviewApply", commonController.ReviewApply)
		search := v1.Group("/search")
		{
			// 查询
			search.POST("/doSearch", searchController.SearchExperts)
		}
		recommend := v1.Group("/recommend")
		{
			// 提交推荐信息
			recommend.POST("/commit", recommendController.RecommendSubmit)
			// 上传推荐表
			recommend.POST("/upload", recommendController.RecommendUpload)
			// 下载推荐表
			recommend.GET("/download", recommendController.RecommendDownload)
			// 根据提交id获取信息
			recommend.GET("/getSubmit", recommendController.RecommendGetSubmit)
			// 获取推荐记录
			recommend.GET("/records", recommendController.RecommendRecords)
		}
		apply := v1.Group("/apply")
		{
			// 获取申请文件名
			apply.GET("/fileName", applyController.ApplyFileName)
			// 上传申请表
			apply.POST("/upload", applyController.ApplyUpload)
			// 下载申请表
			apply.GET("/download", applyController.ApplyDownload)
			// 上传照片
			apply.POST("/uploadPhoto", applyController.ApplyUploadPhoto)
			// 获取照片url
			apply.GET("/photoUrl", applyController.ApplyPhotoUrl)
			// 创建申请信息
			// apply.GET("/create", applyController.ApplyCreate)
			// 提交基本信息
			apply.POST("/submitBase", applyController.ApplySubmitBase)
			// 获取基本信息
			apply.GET("/getBase", applyController.ApplyGetBase)
			// 提交专业类别
			apply.POST("/submitMajor", applyController.ApplySubmitMajor)
			// 获取专业类别
			apply.GET("/getMajor", applyController.ApplyGetMajor)
			// 提交专攻领域
			apply.POST("/submitResearchField", applyController.ApplySubmitResearchField)
			// 获取专攻领域
			apply.GET("/getResearchField", applyController.ApplyGetResearchField)
			// 提交个人履历
			apply.POST("/submitResume", applyController.ApplySubmitResume)
			// 获取个人履历
			apply.GET("/getResume", applyController.ApplyGetResume)
			// 提交意见评价
			apply.POST("/submitOpinion", applyController.ApplySubmitOpinion)
			// 获取意见评价
			apply.GET("/getOpinion", applyController.ApplyGetOpinion)
			// 获取推荐记录
			apply.GET("/records", applyController.ApplyRecords)
		}
	}
	r.StaticFS("api/v1/static", http.Dir("./static/upload/picture"))
	return r
}
