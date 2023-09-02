package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func init() {
	var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "yyk*2012", DB: 1})
	res, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("ping_init 出错：", err)
	}
	fmt.Println(res)
	Rdb = rdb
}

// Limiter 定义属性
type Limiter struct {
	// Redis client connection.
	rc *redis.Client
}

// 根据redisURL创建新的limiter并返回
func NewLimiter() (*Limiter, error) {
	// opts, err := redis.ParseURL(redisURL)
	// if err != nil {
	// 	return nil, err
	// }
	opts := redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "yyk*2012",       // Redis 密码，没有设置忽略即可
		DB:       0,                // 使用的数据库编号，默认为0
	}

	rc := redis.NewClient(&opts)
	if err := rc.Ping().Err(); err != nil {
		return nil, err
	}

	return &Limiter{rc: rc}, nil
}

// 通过redis的value判断第几次访问并返回是否允许访问
func (l *Limiter) Allow(key string, events int64, per time.Duration) bool {
	curr := l.rc.LLen(key).Val()
	if curr >= events {
		return false
	}

	if v := l.rc.Exists(key).Val(); v == 0 {
		pipe := l.rc.TxPipeline() //开启事务，同时执行
		pipe.RPush(key, key)

		//设置过期时间
		pipe.Expire(key, per)
		_, _ = pipe.Exec()
	} else {
		l.rc.RPushX(key, key)
	}

	return true
}
