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
	vo.RecommendCompanyVO	`bson:"recommendCompany"`
	UserID 		primitive.ObjectID 	`bson:"userID"`		// 当前用户id
}

// 专家推荐
type RecommendExpert struct {
	vo.RecommendExpertVO	`bson:"recommendExpert"`
	UserID 	primitive.ObjectID	`json:"-" bson:"userID"`	// 当前用户id
	SubmitID 		string 	`json:"-" bson:"submitID"`							// 本次提交id
}

// 根据某次提交id获得单位名
func GetRecommendCompanyNameBySubmitID(submitID string) (string, error) {
	var record Record
	if err := db.DBConn.DB.Collection("records").
		FindOne(db.DBConn.Context, bson.D{{"_id", submitID}}).
		Decode(&record); err != nil {
		return "", err
	}
	return record.CompanyName, nil
}

// 根据某次提交id获得单位信息
func GetRecommendCompanyByName(name string) (*RecommendCompany, error) {
	var recommendCompany RecommendCompany
	if err := db.DBConn.DB.Collection("companies").
		FindOne(db.DBConn.Context, bson.D{{"recommendCompany.name", name}}).
		Decode(&recommendCompany); err != nil {
		return nil, err
	}
	return &recommendCompany, nil
}

// 根据某次提交id获得专家推荐信息
func GetRecommendExpertsBySubmitID(submitID string) ([]*RecommendExpert, error) {
	experts := []*RecommendExpert{}
	cursor, err := db.DBConn.DB.Collection("experts").Find(db.DBConn.Context, bson.D{{"submitID", submitID}})
	if err != nil {
		return experts, err
	}
	defer cursor.Close(db.DBConn.Context)
	for cursor.Next(db.DBConn.Context) {
		var expert RecommendExpert
		if err := cursor.Decode(&expert); err != nil {
			return experts, err
		}
		experts = append(experts, &expert)
	}
	return experts, nil
}

// 保存或更新单位信息
func SaveOrUpdateRecommendCompany(recommendCompany *RecommendCompany) error {
	if _, err := db.DBConn.DB.Collection("companies").
		UpdateOne(db.DBConn.Context, bson.D{{"recommendCompany.name", recommendCompany.Name}}, bson.D{{"$set", bson.D{{"userID", recommendCompany.UserID}, {"recommendCompany", recommendCompany.RecommendCompanyVO}}}}, options.Update().SetUpsert(true)); err != nil {
		return err
	}
	return nil
}


// 根据提交id删除专家
func DeleteRecommendExpertsBySubmitID(submitID string) error {
	if _, err := db.DBConn.DB.Collection("experts").
		DeleteMany(db.DBConn.Context, bson.D{{"submitID", submitID}}); err != nil {
		return err
	}
	return nil
}

// 保存专家信息
func SaveRecommendExperts(recommendExperts []interface{}) error {
	if _, err := db.DBConn.DB.Collection("experts").
		InsertMany(db.DBConn.Context, recommendExperts); err != nil {
		return err
	}
	return nil
}



