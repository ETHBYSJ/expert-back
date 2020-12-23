// 专家申请
package service

import (
	"expert-back/model"
	"expert-back/pkg/conf"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/util"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"time"
)

type ApplyService struct {
	fileService FileService
}

// 获得申请文件名
func (service *ApplyService) ApplyFileName(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	record, err := model.GetFileRecordByUserIDAndType(profile.Id, model.Apply)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Data: "",
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: record.Name,
	})
}

// 获得图片url
func (service *ApplyService) ApplyPhotoUrl(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	photoRecord, err := model.GetFileRecordByUserIDAndType(profile.Id, model.ApplyPhoto)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Data: "",
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: photoRecord.Name,
	})
}

// 上传申请表
func (service *ApplyService) ApplyUpload(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.UploadApplyFile(c, profile.Id)
	return res
}

// 删除申请表
func (service *ApplyService) ApplyDeleteFile(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.DeleteApplyFile(c, profile.Id)
	return res
}

// 下载申请表
func (service *ApplyService) ApplyDownload(c *gin.Context) response.Response {
	res := service.fileService.DownloadFile(c, conf.SystemConfig.File.Download.Apply.Path, conf.SystemConfig.File.Download.Apply.Name)
	return res
}

// 上传照片
func (service *ApplyService) ApplyUploadPhoto(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.UploadPhoto(c, profile.Id)
	return res
}

// 删除照片
func (service *ApplyService) ApplyDeletePhoto(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.DeletePhoto(c, profile.Id)
	return res
}

/*
// 创建专家申请
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
*/

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

// 获取基本信息
func (service *ApplyService) ApplyGetBase(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	_ = model.CreateApply(profile.Id)
	applyBase, err := model.GetApplyBase(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: applyBase,
	})
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

// 获取专业类别
func (service *ApplyService) ApplyGetMajor(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	_ = model.CreateApply(profile.Id)
	applyMajor, err := model.GetApplyMajor(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: applyMajor,
	})
}

// 合并标签
func mergeLabels(majorLabels []string, researchLabels []string) []string {
	res := []string{}
	m := make(map[string]int)
	for _, label := range majorLabels {
		m[label]++
	}
	for _, label := range researchLabels {
		m[label]++
	}
	for label, _ := range m {
		res = append(res, label)
	}
	return res
}

// 提交专攻领域
func (service *ApplyService) ApplySubmitResearchField(c *gin.Context, applyResearchFieldVO *vo.ApplyResearchFieldVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	oldLabels := []string{}
	fields, err := model.GetApplyResearchField(profile.Id)
	if err == nil {
		oldLabels = mergeLabels(fields.MajorLabels, fields.ResearchLabels)
	}
	if err := model.SaveApplyResearchField(profile.Id, applyResearchFieldVO); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	// 保存倒排索引信息
	labels := mergeLabels(applyResearchFieldVO.ResearchLabels, applyResearchFieldVO.MajorLabels)
	err = model.SaveOrUpdateInvertedIndex(oldLabels, labels, profile.Id)
	if err != nil {
		err = model.TryToRepairInvertedIndex(oldLabels, labels, profile.Id)
		if err != nil {
			return response.BuildResponse(map[int]interface{}{
				response.Code: e.ErrorApplySaveInvertedIndex,
			})
		}
		return response.BuildResponse(map[int]interface{}{})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 获取专攻领域
func (service *ApplyService) ApplyGetResearchField(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	_ = model.CreateApply(profile.Id)
	applyResearchField, err := model.GetApplyResearchField(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: applyResearchField,
	})
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
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyUpdate,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 获取个人履历
func (service *ApplyService) ApplyGetResume(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	_ = model.CreateApply(profile.Id)
	applyResume, err := model.GetApplyResume(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: applyResume,
	})
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
	// 需要获得姓名
	applyBase, err := model.GetApplyBase(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	// 保存记录
	record := &model.Record{
		Type:           model.Apply,
		UserID:         profile.Id,
		SubmitID:       "",
		Name: 			applyBase.Name,
		CommonRecordVO: vo.CommonRecordVO{Title: applyBase.Name + "的专家申请", Status: model.ReviewingText, Timestamp: time.Now().Unix()},
	}
	err = model.SaveOrUpdateApplyRecordInfo(record)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendRecordSet,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 获取意见评价
func (service *ApplyService) ApplyGetOpinion(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	_ = model.CreateApply(profile.Id)
	applyOpinion, err := model.GetApplyOpinion(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: applyOpinion,
	})
}

// 获取申请记录
func (service *ApplyService) ApplyRecords(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	records, err := model.GetApplyRecordsByUserID(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyRecordsGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: records,
	})
}
