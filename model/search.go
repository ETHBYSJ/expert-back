package model

import (
	"expert-back/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvertedIndex struct {
	Label 	string					`bson:"label"`
	List 	[]primitive.ObjectID	`bson:"list"`
}

// 保存倒排索引信息
func SaveOrUpdateInvertedIndex(labels []string, userID primitive.ObjectID) error {
	for _, label := range labels {
		var old InvertedIndex
		filter := bson.D{{"label", label}}
		list := []primitive.ObjectID{}
		if err := db.DBConn.DB.Collection("indexes").
			FindOne(db.DBConn.Context, filter).
			Decode(&old); err == nil {
			list = old.List
		}
		list = append(list, userID)
		update := bson.D{{"$set", bson.D{{"label", label}, {"list", list}}}}
		opts := options.Update().SetUpsert(true)
		if _, err := db.DBConn.DB.Collection("indexes").
			UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
			return err
		}
	}
	return nil
}

// 根据标签搜索专家信息
func GetValidExpertsByLabels(labels []string) ([]*ApplyExpert, error) {
	experts := []*ApplyExpert{}
	// 可以用于后续计算匹配程度
	m := make(map[primitive.ObjectID]int)
	for _, label := range labels {
		filter := bson.D{{"label", label}}
		cursor, err := db.DBConn.DB.Collection("indexes").Find(db.DBConn.Context, filter)
		if err != nil {
			return experts, err
		}
		for cursor.Next(db.DBConn.Context) {
			var invertedIndex InvertedIndex
			if err := cursor.Decode(&invertedIndex); err != nil {
				return experts, err
			}
			for _, id := range invertedIndex.List {
				m[id]++
			}
		}
		cursor.Close(db.DBConn.Context)
	}
	for id, _ := range m {
		expert, err := GetApplyByUserID(id)
		if err != nil {
			continue
		}
		records, err := GetApplyRecordsByUserID(id)
		if err != nil {
			continue
		}
		// 检查审核状态
		if records[0].Status == "accepted" {
			experts = append(experts, expert)
		}
	}
	return experts, nil

}
