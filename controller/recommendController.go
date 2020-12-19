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

// 获得专家
func (controller *RecommendController) RecommendGet(c *gin.Context) {
	res := controller.recommendService.RecommendGet(c)
	c.JSON(http.StatusOK, res)
}

// 专家推荐
func (controller *RecommendController) RecommendCommit(c *gin.Context) {
	var recommendVO vo.RecommendVO
	if err := c.ShouldBindJSON(&recommendVO); err == nil {
		res := controller.recommendService.RecommendCommit(c, &recommendVO)
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
	res := controller.recommendService.RecommendUpload(c)
	c.JSON(http.StatusOK, res)
}

