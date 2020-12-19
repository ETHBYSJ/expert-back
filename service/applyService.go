// 专家申请
package service

import (
	"expert-back/model"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	util2 "expert-back/pkg/util"
	"expert-back/util"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
)

type ApplyService struct {
	fileService FileService
}

// 上传照片
func (service *ApplyService) ApplyUpload(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.UploadPhoto(c, profile.Id)
	return res
}

// 创建专家申请并返回id
func (service *ApplyService) ApplyCreate(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	err = model.CreateApply(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyCreate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 提交基本信息
func (service *ApplyService) ApplySubmitBase(c *gin.Context, applyBaseVO *vo.ApplyBaseVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	if err := model.SaveApplyBase(profile.Id, applyBaseVO); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 提交专业类别
func (service *ApplyService) ApplySubmitMajor(c *gin.Context, applyMajorVO *vo.ApplyMajorVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	if err := model.SaveApplyMajor(profile.Id, applyMajorVO); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 提交专攻领域
func (service *ApplyService) ApplySubmitResearchField(c *gin.Context, applyResearchFieldVO *vo.ApplyResearchFieldVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	if err := model.SaveApplyResearchField(profile.Id, applyResearchFieldVO); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 提交个人履历
func (service *ApplyService) ApplySubmitResume(c *gin.Context, applyResumeVO *vo.ApplyResumeVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	if err := model.SaveApplyResume(profile.Id, applyResumeVO); err != nil {
		util2.Log().Info("submit resume error %v", err)
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 提交意见评价
func (service *ApplyService) ApplySubmitOpinion(c *gin.Context, applyOpinionVO *vo.ApplyOpinionVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	if err := model.SaveApplyOpinion(profile.Id, applyOpinionVO); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}
