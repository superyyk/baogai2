package utils

import (
	"fmt"
	"github.com/superyyk/baogai/model"
	"github.com/superyyk/baogai/tool"

	"time"

	"github.com/gin-gonic/gin"
)

var codes []model.Codes

func ChangePass(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	uid := c.Query("uid")
	pass := c.Query("pass")
	Db.Table("codes").Where("tel=? AND code=? AND uid=?", tel, code, uid).First(&codes)
	if len(codes) != 0 { //成功匹配
		res := Db.Table("userdata").Where("tel=?", tel).Update("pass", pass)
		if res.RowsAffected == 1 { //修改成功
			tool.Success(c, "修改成功！", 200)
		} else {
			tool.Fail(c, "修改失败！", codes)
		}
	} else {
		tool.Fail(c, "手机号或验证码不符", 400)
	}

}

func LoginSms(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	uid := c.Query("uid")
	var userdata []model.UserData
	res := make(map[string]interface{})

	Db.Table("codes").Where("tel=? AND code=? AND uid=?", tel, code, uid).First(&codes)
	if len(codes) != 0 { //验证码匹配正确
		//验证码是否过期
		nowtime := time.Now().Unix()

		if nowtime-codes[0].Time > 60*5 { //5分钟后失效
			tool.Fail(c, "验证码已失效", nowtime-codes[0].Time)
			return
		} else {

			//查看是否有该用户
			Db.Table("userdata").Where("tel=?", tel).First(&userdata)
			if len(userdata) != 0 { //找到用户
				sign := GetJWT(24*10, "yyk2012", "yyk")
				user_uid := userdata[0].Uid
				res["sign"] = sign
				res["uid"] = user_uid
				res["user_info"] = userdata
				//CheckManage(c,user_uid)
				tool.Success(c, "登陆成功！", res)
			} else { //未找到，即注册新用户
				user_uuid, _ := GetUuid(20)
				user := model.UserData{
					Tel:  tel,
					Name: "游客",
					Uid:  user_uuid,
					Pass: "123456",
					Time: tool.Int64_string(GetNowTimeStamp()),
					Head: "https://www.kissnet.cn:39500/static/images/logo.png",
				}
				if err := Db.Table("userdata").Create(&user).Error; err != nil { //注册失败
					tool.Fail(c, "注册失败！", 400)
				} else { //注册成功
					sign := GetJWT(24*10, "yyk2012", "yyk")
					//user_uid:=user_uuid
					res["sign"] = sign
					res["uid"] = user_uuid
					//CheckManage(c,user_uuid)
					tool.Success(c, "注册并登陆成功！", res)
				}
			}
		}
	} else { //验证码不匹配
		tool.Fail(c, "手机号与验证码不符", 400)

	}
}

func CheckManage(c *gin.Context, uid string) {
	var manage []model.Manage
	var manage_base []model.ManageData
	//uid:=c.Query("uid")
	Db.Table("manage").Where("owner=?", uid).Find(&manage)
	Db.Table("manage_base").Where("status=?", 0).Find(&manage_base)
	time.Sleep(time.Second * 1)
	fmt.Println("已有：", len(manage), "条")
	if len(manage) != 0 {

		fmt.Println("不用添加")
	} else {

		for k, v := range manage_base {
			go func(v model.ManageData, k int) {

				vv := model.Manage{
					Uid:    v.Id,
					Name:   v.Name,
					Status: v.Status,
					Owner:  uid,
					Ind:    v.Id,
					Img:    v.Img,
				}
				if err := Db.Table("manage").Create(&vv).Error; err != nil {
					fmt.Println(k, "添加管理失败")
				} else {
					fmt.Println(k, "添加完成")
				}
			}(v, k)

		}

	}

}
func CheckManage1(c *gin.Context) {
	var manage []model.Manage
	var manage_base []model.ManageData
	uid := c.Query("uid")
	Db.Table("manage").Where("owner=?", uid).Find(&manage)
	if len(manage) == 0 {
		Db.Table("manage_base").Where("status=?", 0).Find(&manage_base)
		for k, v := range manage_base {
			go func(v model.ManageData, k int) {

				vv := model.Manage{
					Uid:    v.Id,
					Name:   v.Name,
					Status: v.Status,
					Owner:  uid,
					Ind:    v.Id,
					Img:    v.Img,
				}
				if err := Db.Table("manage").Create(&vv).Error; err != nil {
					fmt.Println(k, "添加管理失败")
				} else {
					fmt.Println(k, "添加完成")
				}
			}(v, k)

		}

	}
	fmt.Println("不用添加")

}
