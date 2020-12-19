// 专家推荐
package model

import (
	"expert-back/db"
	util2 "expert-back/pkg/util"
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
func GetCompanyByUserID(userID string) (*RecommendCompany, error) {
	var recommendCompany RecommendCompany
	if err := db.DBConn.DB.Collection("companies").
		FindOne(db.DBConn.Context, bson.D{{"userID", userID}}).
		Decode(&recommendCompany); err != nil {
			return nil, err
	}
	return &recommendCompany, nil
}

// 根据单位id获得专家信息
func GetExpertsByCompanyID(companyID string) ([]*RecommendExpert, error) {
	experts := []*RecommendExpert{}
	cursor, err := db.DBConn.DB.Collection("experts").Find(db.DBConn.Context, bson.D{{Key: "companyID", Value: companyID}})
	if err != nil {
		return experts, err
	}
	defer cursor.Close(db.DBConn.Context)
	for cursor.Next(db.DBConn.Context) {
		expert := RecommendExpert{}
		if err := cursor.Decode(&expert); err != nil {
			util2.Log().Info("err = %v", err)
			return experts, err
		}
		experts = append(experts, &expert)
	}
	return experts, nil
}

// 根据用户id获得专家信息
func GetExpertsByUserID(userID string) ([]*RecommendExpert, error) {
	company, err := GetCompanyByUserID(userID)
	if err != nil {
		// util.Log().Info("err1 = %v", err)
		return []*RecommendExpert{}, nil
	}
	experts, err := GetExpertsByCompanyID(company.CompanyID.Hex())
	if err != nil {
		// util.Log().Info("err2 = %v", err)
		return []*RecommendExpert{}, err
	}
	return experts, nil
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
	if updateRes.UpsertedCount != 0 {
		objectID := updateRes.UpsertedID.(primitive.ObjectID)
		return objectID.Hex(), nil
	}
	oldCompany, err := GetCompanyByUserID(recommendCompany.UserID)
	if err != nil {
		return "", err
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


