package admin

import (
	"github.com/superyyk/yishougai/db"
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"

	"github.com/gin-gonic/gin"
)

var Db = db.Db

func Ping(c *gin.Context) {
	tel := c.Query("tel")
	pass := c.Query("pass")
	res := make(map[string]interface{})
	res["tel"] = tel
	res["pass"] = tool.Md5_salt(pass)
	tool.Success(c, "success", res)
}

func Login(c *gin.Context) {
	tel := c.Query("tel")
	pass := c.Query("pass")
	pass = tool.Md5_salt(pass)
	var user []model.UserInfos
	res := make(map[string]interface{})
	Db.Table("users").Where("tel=? AND ty=? AND pass=?", tel, 1, pass).First(&user)
	if len(user) > 0 { //找到
		token := utils.GetJWT(24*10, user[0].Uid, "yyk")
		res["token"] = token
		res["info"] = user
		tool.Success(c, "登录成功", res)
		return

	} else { //账号密码不匹配
		tool.Fail(c, "账户密码错误或没有权限登录", 400)
		return
	}
}
