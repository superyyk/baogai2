package api

import (
	"poker/db"
	"poker/tool"

	"github.com/gin-gonic/gin"
)

var Rdb db.Rdb_8

func SetRedis(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	var t int64 = 10
	tt := c.Query("t")
	if tt == "" {
		t = 10
	} else {
		t = tool.String_int64(tt)
	}

	err := Rdb.SetRedis(key, value, t)
	if err {
		tool.Success(c, "success", 200)
	} else {
		tool.Success(c, "err", err)
	}
}

func Get_redis(c *gin.Context) {
	key := c.Query("key")
	re := Rdb.GetRedis(key)
	tool.Success(c, "success", re)
}

func Yanshi(c *gin.Context) {
	key := c.Query("key")
	t := tool.String_int64(c.Query("t"))
	err := Rdb.ExpireRedis(key, t)
	if err {
		tool.Success(c, "success", 200)
	} else {
		tool.Success(c, "err", err)
	}
}
