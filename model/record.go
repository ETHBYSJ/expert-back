package model

import (
	"expert-back/db"
	"expert-back/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Apply     = 1
	Recommend = 2
)

// 专家推荐记录/专家申请记录
type Record struct {
	Type              int                `json:"-" bson:"type"`
	UserID            primitive.ObjectID `json:"-" bson:"userID"`
	SubmitID          string             `json:"submitID" bson:"_id"`
	Name    		  string             `json:"-" bson:"name"`			// 代表单位名(专家推荐表)或人名(专家申请表)
	File              string             `json:"-" bson:"file"`
	vo.CommonRecordVO `bson:"commonRecord"`
}


// Type UserID SubmitID File
func SaveOrUpdateRecordBaseInfo(record *Record) error {
	recordDoc := bson.D{{"type", record.Type}, {"userID", record.UserID}, {"_id", record.SubmitID}, {"file", record.File}}
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, bson.D{{"_id", record.SubmitID}}, bson.D{{"$set", recordDoc}}, options.Update().SetUpsert(true)); err != nil {
		return err
	}
	return nil
}

// Type UserID SubmitID DepartmentName CommonRecordVO
func SaveOrUpdateRecordInfo(record *Record) error {
	recordDoc := bson.D{{"type", record.Type}, {"userID", record.UserID}, {"_id", record.SubmitID}, {"name", record.Name}, {"commonRecord", record.CommonRecordVO}}
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, bson.D{{"_id", record.SubmitID}}, bson.D{{"$set", recordDoc}}, options.Update().SetUpsert(true)); err != nil {
		return err
	}
	return nil
}

// 获得记录，通用函数
func getRecordsByUserID(userID primitive.ObjectID, recordType int) ([]*Record, error) {
	records := []*Record{}
	cursor, err := db.DBConn.DB.Collection("records").Find(db.DBConn.Context, bson.D{{"userID", userID}, {"type", recordType}})
	if err != nil {
		return records, err
	}
	defer cursor.Close(db.DBConn.Context)
	for cursor.Next(db.DBConn.Context) {
		var record Record
		if err := cursor.Decode(&record); err != nil {
			return records, err
		}
		records = append(records, &record)
	}
	return records, nil
}

// 根据用户id获得所有推荐记录
func GetRecommendRecordsByUserID(userID primitive.ObjectID) ([]*Record, error) {
	return getRecordsByUserID(userID, Recommend)
}

// 根据用户id获得所有申请记录
func GetApplyRecordsByUserID(userID primitive.ObjectID) ([]*Record, error) {
	return getRecordsByUserID(userID, Apply)
}
