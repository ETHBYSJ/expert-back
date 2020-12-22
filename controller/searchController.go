package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchController struct {
	searchService service.SearchService
}

// 搜索专家信息
func (controller *SearchController) SearchExperts(c *gin.Context) {
	var searchVO vo.SearchVO
	if err := c.ShouldBindJSON(&searchVO); err == nil {
		res := controller.searchService.SearchExperts(c, &searchVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}
