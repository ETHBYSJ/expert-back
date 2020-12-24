package model

import (
	"expert-back/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ApplyFile     = 1 // 专家申请上传文件
	ApplyPhoto    = 2 // 专家申请上传照片
	RecommendFile = 3 // 专家推荐上传文件
)

// 上传文件相关
type FileRecord struct {
	Type     int                `bson:"type"`
	UserID   primitive.ObjectID `bson:"userID"`
	SubmitID string             `bson:"submitID"`
	Name     string             `bson:"name"` // 文件名或图片url
}

// 根据用户id获得文件记录
func GetFileRecordByUserIDAndType(userID primitive.ObjectID, fileType int) (*FileRecord, error) {
	var fileRecord FileRecord
	filter := bson.D{{"userID", userID}, {"type", fileType}}
	if err := db.DBConn.DB.Collection("files").
		FindOne(db.DBConn.Context, filter).
		Decode(&fileRecord); err != nil {
		return nil, err
	}
	return &fileRecord, nil
}

// 根据提交id获得文件记录
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

/*
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
*/

// 根据用户id和文件类型更新文件记录
func SaveOrUpdateFileRecordByUserIDAndType(fileRecord *FileRecord) error {
	filter := bson.D{{"userID", fileRecord.UserID}, {"type", fileRecord.Type}}
	update := bson.D{{"$set", bson.D{{"type", fileRecord.Type}, {"userID", fileRecord.UserID}, {"submitID", fileRecord.SubmitID}, {"name", fileRecord.Name}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := db.DBConn.DB.Collection("files").
		UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
		return err
	}
	return nil
}

// 根据提交id删除文件记录
func DeleteFileRecordBySubmitID(submitID string) error {
	filter := bson.D{{"submitID", submitID}}
	if _, err := db.DBConn.DB.Collection("files").
		DeleteOne(db.DBConn.Context, filter); err != nil {
		return err
	}
	return nil
}

// 根据用户id和文件类型删除文件记录
func DeleteFileRecordByUserIDAndType(userID primitive.ObjectID, fileType int) error {
	filter := bson.D{{"userID", userID}, {"type", fileType}}
	if _, err := db.DBConn.DB.Collection("files").
		DeleteOne(db.DBConn.Context, filter); err != nil {
		return err
	}
	return nil
}
