package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Rdb_8 struct {
	rdb *redis.Client
}

// func init() {
// 	RedisInit()
// }

func (this *Rdb_8) SetRedis(key string, value string, t int64) bool {
	this.rdb = MyRedis

	expire := time.Duration(t) * time.Second
	if err := this.rdb.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func (this *Rdb_8) GetRedis(key string) string {
	this.rdb = MyRedis
	result, err := this.rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (this *Rdb_8) DelRedis(key string) bool {
	this.rdb = MyRedis
	_, err := this.rdb.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (this *Rdb_8) ExpireRedis(key string, t int64) bool {
	// 延长过期时间
	this.rdb = MyRedis
	expire := time.Duration(t) * time.Second
	if err := this.rdb.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
