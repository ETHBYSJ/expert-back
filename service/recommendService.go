// 专家推荐
package service

import (
	"expert-back/model"
	"expert-back/pkg/conf"
	"expert-back/pkg/docx"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/util"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"github.com/unidoc/unioffice/document"
	"strconv"
	"time"
)

type RecommendService struct {
	fileService FileService
}




// 根据提交id获得信息
func (service *RecommendService) RecommendGetSubmit(c *gin.Context, recommendGetSubmitVO *vo.RecommendGetSubmitVO) response.Response {
	submitID := recommendGetSubmitVO.SubmitID
	// 获得单位名
	name, err := model.GetRecommendCompanyNameBySubmitID(submitID)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	recommendCompany, err := model.GetRecommendCompanyByName(name)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	recommendExperts, err := model.GetRecommendExpertsBySubmitID(submitID)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendSubmitGet,
		})
	}
	expertList := []vo.RecommendExpertVO{}
	for _, expert := range recommendExperts {
		expertList = append(expertList, expert.RecommendExpertVO)
	}
	recommend := vo.RecommendVO{
		RecommendCompanyVO: recommendCompany.RecommendCompanyVO,
		List:               expertList,
		SubmitID:           submitID,
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
	// 保存记录
	record := &model.Record{
		Type:           model.Recommend,
		UserID:         profile.Id,
		SubmitID:       recommendVO.SubmitID,
		CompanyName:    recommendVO.RecommendCompanyVO.Name,
		CommonRecordVO: vo.CommonRecordVO{Title: recommendVO.RecommendCompanyVO.Name + "单位的推荐", Status: "reviewing", Timestamp: time.Now().Unix()},
	}
	err = model.SaveOrUpdateRecord(record)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommend,
		})
	}
	// 保存或更新单位信息
	recommendCompany := &model.RecommendCompany{
		RecommendCompanyVO: recommendVO.RecommendCompanyVO,
		UserID: profile.Id,
	}
	err = model.SaveOrUpdateRecommendCompany(recommendCompany)
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
			UserID: profile.Id,
			SubmitID: recommendVO.SubmitID,
		}
	}
	err = model.SaveRecommendExperts(experts)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommend,
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
func (service *RecommendService) RecommendUpload(c *gin.Context) response.Response {
	profile, err := util.GinGetAccountProfile(c)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorGetAccountProfile,
		})
	}
	res := service.fileService.UploadRecommendFile(c, profile.Id)
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
		ageVal, err := strconv.Atoi(age)
		if err != nil {
			continue
		}
		qualification, err := docx.GetCell(table, i, 4)
		if err != nil {
			continue
		}
		company, err := docx.GetCell(table, i, 5)
		if err != nil {
			continue
		}
		title, err := docx.GetCell(table, i, 6)
		if err != nil {
			continue
		}
		duty, err := docx.GetCell(table, i, 7)
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
			Name:          name,
			Sex:           sex,
			Age:           ageVal,
			Qualification: qualification,
			Title:         title,
			Major:         major,
			Company:       company,
			Duty:          duty,
			Phone:         phone,
			Mobile:        mobile,
			Email:         email,
		}
		expertList = append(expertList, &expert)
	}
	return expertList
}

