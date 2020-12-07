// 专家推荐
package model

import (
	"expert-back/db"
	"expert-back/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 单位信息
type RecommendCompany struct {
	CompanyID	primitive.ObjectID	`bson:"_id"`
	vo.RecommendCompanyVO
	UserID 		string 	`json:"userID" bson:"userID"`		// 当前用户id
}

// 专家推荐
type RecommendExpert struct {
	vo.RecommendExpertVO
	CompanyID 		string 	`json:"companyID" bson:"companyID"`			// 推荐单位id
}

// 根据用户id获得单位信息
func GetCompanyByUserID(userID string) *RecommendCompany {
	var recommendCompany RecommendCompany
	if err := db.DBConn.DB.Collection("companies").
		FindOne(db.DBConn.Context, bson.D{{"userID", userID}}).
		Decode(&recommendCompany); err != nil {
			return nil
	}
	return &recommendCompany
}

// 清空某单位对应的专家列表
func ClearExpertsByCompanyID(companyID string) error {
	if _, err := db.DBConn.DB.Collection("experts").
		DeleteMany(db.DBConn.Context, bson.D{{"companyID", companyID}}); err != nil {
			return err
	}
	return nil
}


// 保存单位信息
func SaveOrUpdateCompanyInfo(recommendCompany *RecommendCompany) (string, error) {
	updateRes, err := db.DBConn.DB.Collection("companies").
		UpdateOne(db.DBConn.Context, bson.D{{"userID", recommendCompany.UserID}}, bson.D{{"$set", bson.D{{"recommendcompanyvo", recommendCompany.RecommendCompanyVO}}}}, options.Update().SetUpsert(true))
	if err != nil {
		return "", err
	}
	// util.Log().Info("%v %v %v %v", updateRes.UpsertedID, updateRes.UpsertedCount, updateRes.MatchedCount, updateRes.ModifiedCount)
	if updateRes.UpsertedCount != 0 {
		objectID := updateRes.UpsertedID.(primitive.ObjectID)
		return objectID.Hex(), nil
	}
	oldCompany := GetCompanyByUserID(recommendCompany.UserID)
	if oldCompany == nil {
		return "", nil
	}
	return oldCompany.CompanyID.Hex(), nil
}

// 保存专家信息
func SaveExpertsInfo(recommendExperts []interface{}) error {
	if _, err := db.DBConn.DB.Collection("experts").
		InsertMany(db.DBConn.Context, recommendExperts); err != nil {
		return err
	}
	return nil
}


