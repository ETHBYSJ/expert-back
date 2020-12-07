// 专家推荐
package service

import (
	"expert-back/model"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
)

type RecommendService struct {

}

func (service *RecommendService) Recommend(c *gin.Context, recommendVO *vo.RecommendVO) response.Response {
	recommendCompany := &model.RecommendCompany{UserID: recommendVO.UserID, RecommendCompanyVO: recommendVO.RecommendCompanyVO}
	// 保存或更新单位信息
	companyID, err := model.SaveOrUpdateCompanyInfo(recommendCompany)
	if companyID == "" || err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_RECOMMEND,
		})
	}
	// 更新专家列表
	_ = model.ClearExpertsByCompanyID(companyID)
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


