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

}

// 专家推荐
func (controller *RecommendController) Recommend(c *gin.Context) {
	var recommendVO vo.RecommendVO
	var recommendService service.RecommendService
	if err := c.ShouldBindJSON(&recommendVO); err == nil {
		res := recommendService.Recommend(c, &recommendVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.CODE: e.HTTP_BADREQUEST,
		}))
	}
}