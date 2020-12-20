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
type RecommendDepartment struct {
	vo.RecommendDepartmentVO `bson:"recommendDepartment"`
	UserID                   primitive.ObjectID `bson:"userID"` // 当前用户id
}

// 专家推荐
type RecommendExpert struct {
	vo.RecommendExpertVO `bson:"recommendExpert"`
	UserID               primitive.ObjectID `json:"-" bson:"userID"`   // 当前用户id
	SubmitID             string             `json:"-" bson:"submitID"` // 本次提交id
}

// 根据某次提交id获得提交记录
func GetRecommendRecordBySubmitID(submitID string) (*Record, error) {
	var record Record
	if err := db.DBConn.DB.Collection("records").
		FindOne(db.DBConn.Context, bson.D{{"_id", submitID}}).
		Decode(&record); err != nil {
		return nil, err
	}
	return &record, nil
}

// 根据某次提交id获得单位信息
func GetRecommendDepartmentByName(name string) (*RecommendDepartment, error) {
	var recommendDepartment RecommendDepartment
	if err := db.DBConn.DB.Collection("departments").
		FindOne(db.DBConn.Context, bson.D{{"recommendDepartment.name", name}}).
		Decode(&recommendDepartment); err != nil {
		return nil, err
	}
	return &recommendDepartment, nil
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
func SaveOrUpdateRecommendDepartment(recommendDepartment *RecommendDepartment) error {
	if _, err := db.DBConn.DB.Collection("departments").
		UpdateOne(db.DBConn.Context, bson.D{{"recommendDepartment.name", recommendDepartment.Name}}, bson.D{{"$set", bson.D{{"userID", recommendDepartment.UserID}, {"recommendDepartment", recommendDepartment.RecommendDepartmentVO}}}}, options.Update().SetUpsert(true)); err != nil {
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
