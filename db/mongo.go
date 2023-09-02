package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mdb *mongo.Database

var MsgCollection *mongo.Collection

func init() {
	uri := "mongodb://poker_mongo:yyk*2012@127.0.0.1:27017/poker_mongo"
	name := "poker_mongo"
	maxTime := time.Duration(10) //超时
	table := "msg"
	db, err := ConnectToDB(uri, name, maxTime)
	if err != nil {
		panic("链接mongo错误")

	}
	MsgCollection = db.Collection(table)

}

func ConnectToDB(uri, name string, timeout time.Duration) (*mongo.Database, error) {
	//设置链接超时
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//链接uri
	o := options.Client().ApplyURI(uri)
	//发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	//返回client
	return client.Database(name), err

}
