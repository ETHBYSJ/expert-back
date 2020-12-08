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
)

type RecommendService struct {
	fileService FileService
}

// 提交专家推荐信息
func (service *RecommendService) RecommendCommit(c *gin.Context, recommendVO *vo.RecommendVO) response.Response {
	recommendCompany := &model.RecommendCompany{UserID: recommendVO.UserID, RecommendCompanyVO: recommendVO.RecommendCompanyVO}
	// 保存或更新单位信息
	companyID, err := model.SaveOrUpdateCompanyInfo(recommendCompany)
	if companyID == "" || err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_RECOMMEND,
		})
	}
	// 按照现在的需求，一旦提交专家推荐信息不能修改
	// _ = model.ClearExpertsByCompanyID(companyID)
	list := recommendVO.List
	experts := make([]interface{}, len(list))
	// 保存专家信息
	for i := 0; i < len(list); i++ {
		experts[i] = &model.RecommendExpert{CompanyID: companyID, RecommendExpertVO: list[i]}
	}
	if err := model.SaveExpertsInfo(experts); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_RECOMMEND,
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
	userID := recommendUploadVO.UserID
	res := service.fileService.UploadFile(c, userID)
	if res.Code != e.SUCCESS {
		return res
	}
	// 解析docx文档
	path := res.Data.(string)
	tables, err := docx.Parse(path)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_PARSE,
		})
	}
	expertList := ConstructExpertList(&tables[1])
	util.Log().Info("list: %v", expertList)
	return response.BuildResponse(map[int]interface{}{
		response.DATA: expertList,
	})
}

// 解析表格获取数据
func ConstructExpertList(table *document.Table) []vo.RecommendExpertVO {
	// 处理具体解析逻辑
	rows := table.Rows()
	// 数据行数
	rowNum := len(rows) - 1
	expertList := []vo.RecommendExpertVO{}
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
		expertList = append(expertList, expert)
	}
	return expertList
}

