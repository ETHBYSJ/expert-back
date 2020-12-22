package model

import (
	"errors"
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

const (
	AcceptedText = "accept"
	ReviewingText = "reviewing"
	FailedText = "failed"
)

const (
	Accepted = 1
	Reviewing = 2
	Failed = 3
)

var StatusMap = map[int]string {
	Accepted: AcceptedText,
	Reviewing: ReviewingText,
	Failed: FailedText,
}

var errInvalidStatus = errors.New("无效的审核状态")

// 专家推荐记录/专家申请记录
type Record struct {
	Type              int                `json:"-" bson:"type"`				// 记录类型
	UserID            primitive.ObjectID `json:"-" bson:"userID"`			// 用户id
	SubmitID          string             `json:"submitID" bson:"submitID"`	// 提交id
	Name    		  string             `json:"-" bson:"name"`				// 代表单位名(专家推荐表)或人名(专家申请表)
	vo.CommonRecordVO `bson:"commonRecord"`
}

// 修改专家推荐审核状态
func UpdateRecommendRecordStatus(submitID string, status int) error {
	if status < Accepted || status > Failed {
		return errInvalidStatus
	}
	filter := bson.D{{"submitID", submitID}, {"type", Recommend}}
	update := bson.D{{"$set", bson.D{{"commonRecord.status", StatusMap[status]}}}}
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, filter, update); err != nil {
		return err
	}
	return nil
}

// 修改专家申请审核状态
func UpdateApplyRecordStatus(userID primitive.ObjectID, status int) error {
	if status < Accepted || status > Failed {
		return errInvalidStatus
	}
	filter := bson.D{{"userID", userID}, {"type", Apply}}
	update := bson.D{{"$set", bson.D{{"commonRecord.status", StatusMap[status]}}}}
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, filter, update); err != nil {
		return err
	}
	return nil
}

// Type UserID SubmitID Name CommonRecordVO
func SaveOrUpdateRecommendRecordInfo(record *Record) error {
	filter := bson.D{{"submitID", record.SubmitID}, {"type", Recommend}}
	update := bson.D{{"$set", bson.D{{"type", record.Type}, {"userID", record.UserID}, {"submitID", record.SubmitID}, {"name", record.Name}, {"commonRecord", record.CommonRecordVO}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
		return err
	}
	return nil
}

// Type UserID SubmitID Name CommonRecordVO
func SaveOrUpdateApplyRecordInfo(record *Record) error {
	filter := bson.D{{"userID", record.UserID}, {"type", Apply}}
	update := bson.D{{"$set", bson.D{{"type", record.Type}, {"userID", record.UserID}, {"submitID", record.SubmitID}, {"name", record.Name}, {"commonRecord", record.CommonRecordVO}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := db.DBConn.DB.Collection("records").
		UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
		return err
	}
	return nil
}

// 获得记录，通用函数
func getRecordsByUserID(userID primitive.ObjectID, recordType int) ([]*Record, error) {
	records := []*Record{}
	filter := bson.D{{"userID", userID}, {"type", recordType}}
	opts := options.Find().SetSort(bson.D{{"commonRecord.timestamp", -1}})
	// 按时间降序
	cursor, err := db.DBConn.DB.Collection("records").Find(db.DBConn.Context, filter, opts)
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
