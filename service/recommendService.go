// 专家推荐
package service

import (
	"expert-back/model"
	"expert-back/pkg/conf"
	"expert-back/pkg/docx"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	util2 "expert-back/pkg/util"
	"expert-back/util"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"github.com/unidoc/unioffice/document"
	"time"
)

type RecommendService struct {
	fileService FileService
}

// 根据提交id获得信息
func (service *RecommendService) RecommendGetSubmit(c *gin.Context, recommendGetSubmitVO *vo.RecommendGetSubmitVO) response.Response {
	submitID := recommendGetSubmitVO.SubmitID
	fileName := ""
	// 获得文件名
	fileRecord, err := model.GetFileRecordBySubmitID(submitID)
	if err == nil {
		fileName = fileRecord.Name
	}
	// 获得单位名
	record, err := model.GetRecommendRecordBySubmitID(submitID)
	if err != nil {
		util2.Log().Info("err1 %v", err)
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	recommendDepartment, err := model.GetRecommendDepartmentByName(record.Name)
	if err != nil {
		util2.Log().Info("err2 %v", err)
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	recommendExperts, err := model.GetRecommendExpertsBySubmitID(submitID)
	if err != nil {
		util2.Log().Info("err3 %v", err)
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	expertList := []vo.RecommendExpertVO{}
	for _, expert := range recommendExperts {
		expertList = append(expertList, expert.RecommendExpertVO)
	}
	recommend := vo.RecommendRetVO{
		RecommendVO: vo.RecommendVO{
			RecommendDepartmentVO: recommendDepartment.RecommendDepartmentVO,
			List:                  expertList,
			SubmitID:              submitID,
		},
		File:                  fileName,
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: recommend,
	})
}

// 提交专家推荐信息
func (service *RecommendService) RecommendSubmit(c *gin.Context, recommendVO *vo.RecommendVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	// 保存或更新单位信息
	recommendDepartment := &model.RecommendDepartment{
		RecommendDepartmentVO: recommendVO.RecommendDepartmentVO,
		UserID:                profile.Id,
	}
	err = model.SaveOrUpdateRecommendDepartment(recommendDepartment)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommend,
		})
	}
	// 根据提交id删除专家
	err = model.DeleteRecommendExpertsBySubmitID(recommendVO.SubmitID)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommend,
		})
	}
	list := recommendVO.List
	experts := make([]interface{}, len(list))
	// 保存专家信息
	for i := 0; i < len(list); i++ {
		experts[i] = &model.RecommendExpert{
			RecommendExpertVO: list[i],
			UserID:            profile.Id,
			SubmitID:          recommendVO.SubmitID,
		}
	}
	err = model.SaveRecommendExperts(experts)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommend,
		})
	}
	// 保存记录
	record := &model.Record{
		Type:           model.Recommend,
		UserID:         profile.Id,
		SubmitID:       recommendVO.SubmitID,
		Name: 			recommendVO.RecommendDepartmentVO.Name,
		CommonRecordVO: vo.CommonRecordVO{Title: recommendVO.RecommendDepartmentVO.Name + "的推荐", Status: model.ReviewingText, Timestamp: time.Now().Unix()},
	}
	err = model.SaveOrUpdateRecommendRecordInfo(record)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendRecordSet,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 下载推荐表
func (service *RecommendService) RecommendDownload(c *gin.Context) response.Response {
	res := service.fileService.DownloadFile(c, conf.SystemConfig.File.Download.Recommend.Path, conf.SystemConfig.File.Download.Recommend.Name)
	return res
}

// 上传推荐表
func (service *RecommendService) RecommendUpload(c *gin.Context, recommendUploadVO *vo.RecommendUploadVO) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.UploadRecommendFile(c, recommendUploadVO.SubmitID, profile.Id)
	return res
	/*
	if res.Code != e.Success {
		return res
	}
	// 解析docx文档
	path := res.Data.(string)
	tables, err := docx.Parse(path)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendParse,
		})
	}
	expertList := ConstructExpertList(&tables[1])
	return response.BuildResponse(map[int]interface{}{
		response.Data: expertList,
	})
	*/
}

//获取推荐记录
func (service *RecommendService) RecommendRecords(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	records, err := model.GetRecommendRecordsByUserID(profile.Id)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendRecordsGet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: records,
	})
}

// 解析表格获取数据
func ConstructExpertList(table *document.Table) []*vo.RecommendExpertVO {
	// 处理具体解析逻辑
	rows := table.Rows()
	// 数据行数
	rowNum := len(rows) - 1
	expertList := []*vo.RecommendExpertVO{}
	for i := 1; i <= rowNum; i++ {
		name, err := docx.GetCell(table, i, 1)
		if err != nil {
			continue
		}
		sex, err := docx.GetCell(table, i, 2)
		if err != nil {
			continue
		}
		age, err := docx.GetCell(table, i, 3)
		if err != nil {
			continue
		}
		if err != nil {
			continue
		}
		edu, err := docx.GetCell(table, i, 4)
		if err != nil {
			continue
		}
		dept, err := docx.GetCell(table, i, 5)
		if err != nil {
			continue
		}
		title, err := docx.GetCell(table, i, 6)
		if err != nil {
			continue
		}
		post, err := docx.GetCell(table, i, 7)
		if err != nil {
			continue
		}
		major, err := docx.GetCell(table, i, 8)
		if err != nil {
			continue
		}
		phone, err := docx.GetCell(table, i, 9)
		if err != nil {
			continue
		}
		mobile, err := docx.GetCell(table, i, 10)
		if err != nil {
			continue
		}
		email, err := docx.GetCell(table, i, 11)
		if err != nil {
			continue
		}
		expert := vo.RecommendExpertVO{
			Name:   name,
			Sex:    sex,
			Age:    age,
			Edu:    edu,
			Title:  title,
			Major:  major,
			Dept:   dept,
			Post:   post,
			Phone:  phone,
			Mobile: mobile,
			Email:  email,
		}
		expertList = append(expertList, &expert)
	}
	return expertList
}
