package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplyController struct {
	applyService service.ApplyService
}

// 创建申请
func (controller *ApplyController) ApplyCreate(c *gin.Context) {
	res := controller.applyService.ApplyCreate(c)
	c.JSON(http.StatusOK, res)
}

// 提交基本信息
func (controller *ApplyController) ApplySubmitBase(c *gin.Context) {
	var applyBaseVO vo.ApplyBaseVO
	if err := c.ShouldBindJSON(&applyBaseVO); err == nil {
		res := controller.applyService.ApplySubmitBase(c, &applyBaseVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 提交专业类别
func (controller *ApplyController) ApplySubmitMajor(c *gin.Context) {
	var applyMajorVO vo.ApplyMajorVO
	if err := c.ShouldBindJSON(&applyMajorVO); err == nil {
		res := controller.applyService.ApplySubmitMajor(c, &applyMajorVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 提交专攻领域
func (controller *ApplyController) ApplySubmitResearchField(c *gin.Context) {
	var applyResearchFieldVO vo.ApplyResearchFieldVO
	if err := c.ShouldBindJSON(&applyResearchFieldVO); err == nil {
		res := controller.applyService.ApplySubmitResearchField(c, &applyResearchFieldVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 提交个人履历
func (controller *ApplyController) ApplySubmitResume(c *gin.Context) {
	var applyResumeVO vo.ApplyResumeVO
	if err := c.ShouldBindJSON(&applyResumeVO); err == nil {
		res := controller.applyService.ApplySubmitResume(c, &applyResumeVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}

// 提交意见评价
func (controller *ApplyController) ApplySubmitOpinion(c *gin.Context) {
	var applyOpinionVO vo.ApplyOpinionVO
	if err := c.ShouldBindJSON(&applyOpinionVO); err == nil {
		res := controller.applyService.ApplySubmitOpinion(c, &applyOpinionVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.Code: e.HttpBadRequest,
		}))
	}
}
