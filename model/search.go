package model

import (
	"expert-back/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvertedIndex struct {
	Label string               `bson:"label"`
	List  []primitive.ObjectID `bson:"list"`
}

// 合并用户id
func mergeUserID(list []primitive.ObjectID, userID primitive.ObjectID) []primitive.ObjectID {
	exist := false
	for _, id := range list {
		if id == userID {
			exist = true
			break
		}
	}
	if !exist {
		list = append(list, userID)
	}
	return list
}

// 开发阶段使用，在出现数据不一致时使用
func TryToRepairInvertedIndex(oldLabels []string, labels []string, userID primitive.ObjectID) error {
	// 删除旧索引
	for _, label := range oldLabels {
		var old InvertedIndex
		filter := bson.D{{"label", label}}
		list := []primitive.ObjectID{}
		if err := db.DBConn.DB.Collection("indexes").
			FindOne(db.DBConn.Context, filter).
			Decode(&old); err == nil {
			list = old.List
		} else {
			continue
		}
		tmp := []primitive.ObjectID{}
		// 从list中删掉用户id
		for _, id := range list {
			if id != userID {
				tmp = append(tmp, id)
			}
		}
		list = tmp
		if len(list) != 0 {
			update := bson.D{{"$set", bson.D{{"label", label}, {"list", list}}}}
			opts := options.Update().SetUpsert(true)
			if _, err := db.DBConn.DB.Collection("indexes").
				UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
				return err
			}
		} else {
			if _, err := db.DBConn.DB.Collection("indexes").
				DeleteOne(db.DBConn.Context, filter); err != nil {
				return err
			}
		}

	}

	// 新增索引
	for _, label := range labels {
		var old InvertedIndex
		filter := bson.D{{"label", label}}
		list := []primitive.ObjectID{}
		if err := db.DBConn.DB.Collection("indexes").
			FindOne(db.DBConn.Context, filter).
			Decode(&old); err == nil {
			list = old.List
		} else if err != mongo.ErrNoDocuments {
			return err
		}
		// list = append(list, userID)
		list = mergeUserID(list, userID)
		update := bson.D{{"$set", bson.D{{"label", label}, {"list", list}}}}
		opts := options.Update().SetUpsert(true)
		if _, err := db.DBConn.DB.Collection("indexes").
			UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
			return err
		}
	}
	return nil
}

// 保存倒排索引信息
func SaveOrUpdateInvertedIndex(oldLabels []string, labels []string, userID primitive.ObjectID) error {
	oldMap := make(map[string]int)
	newMap := make(map[string]int)
	for _, label := range oldLabels {
		oldMap[label]++
	}
	for _, label := range labels {
		newMap[label]++
	}
	delList := []string{}
	addList := []string{}
	// 需要删除和新增的标签
	for label, _ := range oldMap {
		if _, ok := newMap[label]; !ok {
			delList = append(delList, label)
		}
	}
	for label, _ := range newMap {
		if _, ok := oldMap[label]; !ok {
			addList = append(addList, label)
		}
	}
	// 删除旧索引
	for _, label := range delList {
		var old InvertedIndex
		filter := bson.D{{"label", label}}
		list := []primitive.ObjectID{}
		if err := db.DBConn.DB.Collection("indexes").
			FindOne(db.DBConn.Context, filter).
			Decode(&old); err == nil {
			list = old.List
		} else {
			return err
		}
		tmp := []primitive.ObjectID{}
		// 从list中删掉用户id
		for _, id := range list {
			if id != userID {
				tmp = append(tmp, id)
			}
		}
		list = tmp
		if len(list) != 0 {
			update := bson.D{{"$set", bson.D{{"label", label}, {"list", list}}}}
			opts := options.Update().SetUpsert(true)
			if _, err := db.DBConn.DB.Collection("indexes").
				UpdateOne(db.DBConn.Context, filter, update, opts); err != nil {
				return err
			}
		} else {
			if _, err := db.DBConn.DB.Collection("indexes").
				DeleteOne(db.DBConn.Context, filter); err != nil {
				return err
			}
		}

	}
	// 新增索引
	for _, label := range addList {
		var old InvertedIndex
		filter := bson.D{{"label", label}}
		list := []primitive.ObjectID{}
		if err := db.DBConn.DB.Collection("indexes").
			FindOne(db.DBConn.Context, filter).
			Decode(&old); err == nil {
			list = old.List
		} else if err != mongo.ErrNoDocuments {
			return err
		}
		// list = append(list, userID)
		list = mergeUserID(list, userID)
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
		/*
			// 检查审核状态
			records, err := GetApplyRecordsByUserID(id)
			if err != nil {
				continue
			}
			if records[0].Status == "accepted" {
				experts = append(experts, expert)
			}
		*/
		experts = append(experts, expert)
	}
	return experts, nil

}
