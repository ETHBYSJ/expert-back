package service

import (
	"expert-back/model"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonService struct {
}

// 审核专家推荐
func (service *CommonService) ReviewReCommend(c *gin.Context, reviewRecommendVO *vo.ReviewRecommendVO) response.Response {
	if err := model.UpdateRecommendRecordStatus(reviewRecommendVO.SubmitID, reviewRecommendVO.Status); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorReview,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}

// 审核专家申请
func (service *CommonService) ReviewApply(c *gin.Context, reviewApplyVO *vo.ReviewApplyVO) response.Response {
	userID, _ := primitive.ObjectIDFromHex(reviewApplyVO.UserID)
	if err := model.UpdateApplyRecordStatus(userID, reviewApplyVO.Status); err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorReview,
		})
	}
	return response.BuildResponse(map[int]interface{}{})
}
