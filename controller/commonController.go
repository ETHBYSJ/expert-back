package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/util"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonController struct {
	commonService service.CommonService
}

// 审核专家推荐
func (controller *CommonController) ReviewRecommend(c *gin.Context) {
	var reviewRecommendVO vo.ReviewRecommendVO
	if err := c.ShouldBind(&reviewRecommendVO); err == nil {
		res := controller.commonService.ReviewReCommend(c, &reviewRecommendVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 审核专家申请
func (controller *CommonController) ReviewApply(c *gin.Context) {
	var reviewApplyVO vo.ReviewApplyVO
	if err := c.ShouldBind(&reviewApplyVO); err == nil {
		res := controller.commonService.ReviewApply(c, &reviewApplyVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 设置cookie
func (controller *CommonController) SetCookie(c *gin.Context) {
	c.SetCookie("iadmin", "e415c0ff-ae61-48f2-9208-d4335b8d1792", 3*24*3600, "/", "", false, true)
	// c.SetCookie("iadmin", "a45f05cb-3e32-4b11-85ce-f59d96b4bd41", 3 * 24 * 3600, "/", "", false, true)
	c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{}))
}

// 测试获得用户信息
func (controller *CommonController) GetAccountProfile(c *gin.Context) {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		}))
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Data: profile,
		}))
	}
}
