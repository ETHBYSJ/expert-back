package service

import (
	"expert-back/model"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchService struct {

}

// 搜索专家信息，暂不考虑排序
func (service *SearchService) SearchExperts(c *gin.Context, searchVO *vo.SearchVO) response.Response {
	expertsByKeyword, err := model.GetValidExpertsByName(searchVO.Keyword)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorSearch,
		})
	}
	expertsByLabels, err := model.GetValidExpertsByLabels(searchVO.Labels)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorSearch,
		})
	}
	experts := mergeSearchResults(expertsByKeyword, expertsByLabels)
	res := []*vo.SearchResultVO{}
	for _, expert := range experts {
		result := &vo.SearchResultVO{
			Labels:       append(expert.ResearchLabels, expert.MajorLabels...),
			Name:         expert.Name,
			Photo:        expert.Photo,
			Introduction: expert.Dept + expert.AdminPost,
		}
		res = append(res, result)
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: res,
	})
}

// 合并两部分查询的结果
func mergeSearchResults(expertsByKeyword []*model.ApplyExpert, expertsByLabels []*model.ApplyExpert) []*model.ApplyExpert {
	res := []*model.ApplyExpert{}
	m := make(map[primitive.ObjectID]*model.ApplyExpert)
	for _, expert := range expertsByKeyword {
		m[expert.UserID] = expert
	}
	for _, expert := range expertsByLabels {
		m[expert.UserID] = expert
	}
	for _, expert := range m {
		res = append(res, expert)
	}
	return res
}