// 专家申请
package model

import (
	"expert-back/db"
	"expert-back/pkg/util"
	"expert-back/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 专家申请
type ApplyExpert struct {
	vo.ApplyBaseVO          `bson:"applyBase"`          // 基本信息
	vo.ApplyMajorVO         `bson:"applyMajor"`         // 专业类别
	vo.ApplyResearchFieldVO `bson:"applyResearchField"` // 专攻领域
	vo.ApplyResumeVO        `bson:"applyResume"`        // 个人履历
	vo.ApplyOpinionVO       `bson:"applyOpinion"`       // 意见评价
	UserID                  primitive.ObjectID          `bson:"userID"` // 用户id
}

// 根据用户id获得申请信息
func GetApplyByUserID(userID primitive.ObjectID) (*ApplyExpert, error) {
	var applyExpert ApplyExpert
	if err := db.DBConn.DB.Collection("apply").
		FindOne(db.DBConn.Context, bson.D{{"userID", userID}}).
		Decode(&applyExpert); err != nil {
		return nil, err
	}
	return &applyExpert, nil
}

// 创建申请
func createApply(userID primitive.ObjectID) error {
	apply, err := GetApplyByUserID(userID)
	// 没有记录，新建
	if err != nil {
		util.Log().Info("model create apply %v", err)
		apply = &ApplyExpert{
			UserID: userID,
		}
		// 新建
		if _, err = db.DBConn.DB.Collection("apply").
			InsertOne(db.DBConn.Context, apply); err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

// 通用函数
func saveApplyInfo(userID primitive.ObjectID, key string, value interface{}) error {
	if _, err := db.DBConn.DB.Collection("apply").
		UpdateOne(db.DBConn.Context, bson.D{{"userID", userID}}, bson.D{{"$set", bson.D{{key, value}}}}); err != nil {
		return err
	}
	return nil
}


// 保存基本信息
func SaveApplyBase(userID primitive.ObjectID, applyBaseVO *vo.ApplyBaseVO) error {
	if err := createApply(userID); err != nil {
		return err
	}
	return saveApplyInfo(userID, "applyBase", applyBaseVO)
}

// 获取基本信息
func GetApplyBase(userID primitive.ObjectID) (*vo.ApplyBaseVO, error) {
	expert, err := GetApplyByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &expert.ApplyBaseVO, nil
}

// 保存专业类别
func SaveApplyMajor(userID primitive.ObjectID, applyMajorVO *vo.ApplyMajorVO) error {
	if err := createApply(userID); err != nil {
		return err
	}
	return saveApplyInfo(userID, "applyMajor", applyMajorVO)
}

// 获取专业类别
func GetApplyMajor(userID primitive.ObjectID) (*vo.ApplyMajorVO, error) {
	expert, err := GetApplyByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &expert.ApplyMajorVO, nil
}

// 保存专攻领域
func SaveApplyResearchField(userID primitive.ObjectID, applyResearchFieldVO *vo.ApplyResearchFieldVO) error {
	if err := createApply(userID); err != nil {
		return err
	}
	return saveApplyInfo(userID, "applyResearchField", applyResearchFieldVO)
}

// 获取专业类别
func GetApplyResearchField(userID primitive.ObjectID) (*vo.ApplyResearchFieldVO, error) {
	expert, err := GetApplyByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &expert.ApplyResearchFieldVO, nil
}

// 保存个人履历
func SaveApplyResume(userID primitive.ObjectID, applyResumeVO *vo.ApplyResumeVO) error {
	if err := createApply(userID); err != nil {
		return err
	}
	return saveApplyInfo(userID, "applyResume", applyResumeVO)
}

// 获取个人履历
func GetApplyResume(userID primitive.ObjectID) (*vo.ApplyResumeVO, error) {
	expert, err := GetApplyByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &expert.ApplyResumeVO, nil
}

// 保存意见评价
func SaveApplyOpinion(userID primitive.ObjectID, applyOpinionVO *vo.ApplyOpinionVO) error {
	if err := createApply(userID); err != nil {
		return err
	}
	return saveApplyInfo(userID, "applyOpinion", applyOpinionVO)
}

// 获取意见评价
func GetApplyOpinion(userID primitive.ObjectID) (*vo.ApplyOpinionVO, error) {
	expert, err := GetApplyByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &expert.ApplyOpinionVO, nil
}
