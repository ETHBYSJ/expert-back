package model

import (
	"expert-back/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ApplyFile = 1		// 专家申请上传文件
	RecommendFile = 2 	// 专家推荐上传文件
)

// 上传文件相关
type FileRecord struct {
	Type 		int 				`bson:"type"`
	UserID 		primitive.ObjectID	`bson:"userID"`
	SubmitID    string             	`bson:"submitID"`
	Name 		string 				`bson:"name"`
}

// 根据用户id获得文件记录
func GetFileRecordByUserID(userID primitive.ObjectID) (*FileRecord, error) {
	var fileRecord FileRecord
	filter := bson.D{{"userID", userID}}
	if err := db.DBConn.DB.Collection("files").
		FindOne(db.DBConn.Context, filter).
		Decode(&fileRecord); err != nil {
		return nil, err
	}
	return &fileRecord, nil
}

// 根据提交id获得问卷记录
func GetFileRecordBySubmitID(submitID string) (*FileRecord, error) {
	var fileRecord FileRecord
	filter := bson.D{{"submitID", submitID}}
	if err := db.DBConn.DB.Collection("files").
		FindOne(db.DBConn.Context, filter).
		Decode(&fileRecord); err != nil {
		return nil, err
	}
	return &fileRecord, nil
}

// 根据提交id更新文件记录
func SaveOrUpdateFileRecordBySubmitID(fileRecord *FileRecord) error {
	filter := bson.D{{"submitID", fileRecord.SubmitID}}
	update := bson.D{{"$set", bson.D{{"type", fileRecord.Type}, {"userID", fileRecord.UserID}, {"submitID", fileRecord.SubmitID}, {"name", fileRecord.Name}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := db.DBConn.DB.Collection("files").
		UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
		return err
	}
	return nil
}

// 根据用户id更新文件记录
func SaveOrUpdateFileRecordByUserID(fileRecord *FileRecord) error {
	filter := bson.D{{"userID", fileRecord.UserID}}
	update := bson.D{{"$set", bson.D{{"type", fileRecord.Type}, {"userID", fileRecord.UserID}, {"submitID", fileRecord.SubmitID}, {"name", fileRecord.Name}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := db.DBConn.DB.Collection("files").
		UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
		return err
	}
	return nil
}
