// 专家推荐
package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecommendController struct {
	recommendService service.RecommendService
}

// 根据提交id获取信息
func (controller *RecommendController) RecommendGetSubmit(c *gin.Context) {
	var recommendGetSubmitVO vo.RecommendGetSubmitVO
	if err := c.ShouldBindQuery(&recommendGetSubmitVO); err == nil {
		res := controller.recommendService.RecommendGetSubmit(c, &recommendGetSubmitVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 专家推荐
func (controller *RecommendController) RecommendSubmit(c *gin.Context) {
	var recommendVO vo.RecommendVO
	if err := c.ShouldBindJSON(&recommendVO); err == nil {
		res := controller.recommendService.RecommendSubmit(c, &recommendVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 专家推荐文件下载
func (controller *RecommendController) RecommendDownload(c *gin.Context) {
	res := controller.recommendService.RecommendDownload(c)
	if res.Code != e.Success {
		c.JSON(http.StatusOK, res)
	}
}

// 专家推荐文件上传
func (controller *RecommendController) RecommendUpload(c *gin.Context) {
	var recommendUploadVO vo.RecommendUploadVO
	if err := c.ShouldBindQuery(&recommendUploadVO); err == nil {
		res := controller.recommendService.RecommendUpload(c, &recommendUploadVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 获取推荐记录
func (controller *RecommendController) RecommendRecords(c *gin.Context) {
	res := controller.recommendService.RecommendRecords(c)
	c.JSON(http.StatusOK, res)
}
