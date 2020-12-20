package model

import (
	"expert-back/db"
	"expert-back/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Apply = 1
	Recommend = 2
)

// 专家推荐记录/专家申请记录
type Record struct {
	Type 		int 				`json:"-" bson:"type"`
	UserID 		primitive.ObjectID	`json:"-" bson:"userID"`
	SubmitID	string 				`json:"-" bson:"_id"`
	CompanyName	string 				`json:"-" bson:"companyName"`
	vo.CommonRecordVO				`bson:"commonRecord"`
}

// 保存或更新
func SaveOrUpdateRecord(record *Record) error {
	recordDoc := bson.D{{"type", record.Type}, {"userID", record.UserID}, {"_id", record.SubmitID}, {"companyName", record.CompanyName}, {"commonRecord", record.CommonRecordVO}}
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



