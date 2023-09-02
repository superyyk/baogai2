package db

import (
	"context"
	"fmt"
	"log"

	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/model"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cn = model.ConnInfo{
	"baogai",   //本地 root gaoding
	"yyk*2012", //本地 yyk*2012  yyk2012
	"127.0.0.1",
	"baogai", //kuanku
	3306,
}
var Db *gorm.DB

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var (
	MyRedis *redis.Client
	ctx     = context.Background()
)

// 数据库连接池
func init() {

	Db = DbConn(cn.Username, cn.Password, cn.Host, cn.Db, cn.Port)

	RedisInit()
}

func DbConn(Username, Password, Host, Db string, Port int) *gorm.DB {

	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Username, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	CheckErr(err)

	db.SingularTable(true)
	//defer db.Close()
	db.DB().SetMaxIdleConns(1000000) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(1000000) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于1，超过的连接会被连接池关闭。
	return db
}

func RedisInit() {
	MyRedis = redis.NewClient(&redis.Options{
		Addr: config.Redis.Host,

		Password: config.Redis.Password, // no password set
		DB:       config.Redis.DB,       // use default DB
	})
	_, err := MyRedis.Ping(ctx).Result()
	if err != nil {
		//logging.Error("Redis connect ping failed, err:", err)
		return
	}
	//logging.Error("Redis connect succeeded")
	return

}
