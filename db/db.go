package db

import (
	"context"
	"expert-back/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"strings"
	"time"
)

// MongoDB连接
var DBConn *Database

// 连接数据库
func Init(connection string, dbName string) {
	backgroundCtx := context.Background()
	ctx, cancel := context.WithTimeout(backgroundCtx, 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		util.Log().Panic("连接数据库出错: %s", err)
	}
	ctxPing, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = client.Ping(ctxPing, readpref.Primary())
	if err != nil {
		util.Log().Panic("连接数据库失败: %s", err)
	}
	db := client.Database(dbName)
	DBConn = &Database{DB: db, Client: client, Context: backgroundCtx}
	// 创建集合
	// 单位
	_ = DBConn.DB.CreateCollection(DBConn.Context, "companies")
	// 专家推荐信息
	_ = DBConn.DB.CreateCollection(DBConn.Context, "experts")
	// 创建索引
	// CreateIndex("users", "openId")
}

// 联合索引
func CreateIndex(collection string, unique bool, keys ...string) {
	indexView := DBConn.DB.Collection(collection).Indexes()
	keysDoc := bsonx.Doc{}
	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}
	// 创建索引
	result, err := indexView.CreateOne(
		DBConn.Context,
		mongo.IndexModel{
			Keys: keysDoc,
			Options: options.Index().SetUnique(unique),
		},
	)
	if result == "" || err != nil {
		util.Log().Error("创建索引失败", err)
	} else {
		util.Log().Info("创建索引: %s", result)
	}
}


type Database struct {
	DB 		*mongo.Database
	Client 	*mongo.Client
	Context context.Context
}