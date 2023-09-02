package api

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/superyyk/baogai/model"
	"github.com/superyyk/baogai/tool"
	"time"

	"github.com/gin-gonic/gin"
)

//var Rdb db.Rdb

func GetXiangyan(c *gin.Context) {
	var xiangyan []model.Xiangyan
	var xiangyan_big []model.Xiangyan
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	ty := c.Query("ty")
	offset := (pageNum - 1) * pagesize
	if ty == "0" {
		Db.Table("xiangyan").Where("ty=0 AND status=1").Offset(offset).Limit(pagesize).Order("sort desc").Order("status_vip desc").Find(&xiangyan)
	}
	if ty == "1" {
		Db.Table("xiangyan").Where("ty=1 AND status=1").Offset(offset).Limit(pagesize).Order("sort desc").Order("status_vip desc").Find(&xiangyan_big)
	}
	if ty == "" {
		Db.Table("xiangyan").Where("ty=0 AND status=1").Offset(offset).Limit(pagesize).Order("sort desc").Order("status_vip desc").Find(&xiangyan)
		Db.Table("xiangyan").Where("ty=1 AND status=1").Offset(offset).Limit(pagesize).Order("sort desc").Order("status_vip desc").Find(&xiangyan_big)
	}

	res := make(map[string]interface{})
	res["small"] = xiangyan
	res["big"] = xiangyan_big
	tool.Success(c, "success", res)
}

func GetXiangyanByName(c *gin.Context) {
	var xiangyan []model.Xiangyan
	var xiangyan_big []model.Xiangyan
	var list []model.Xiangyan
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	name := c.Query("name")
	ty := c.Query("ty")
	offset := (pageNum - 1) * pagesize
	res := make(map[string]interface{})

	if ty == "0" {
		Db.Table("xiangyan").Where("name like ? AND ty=0", "%"+name+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan)
		res["name"] = xiangyan
	}

	if ty == "1" {
		Db.Table("xiangyan").Where("name like ? AND ty=1", "%"+name+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_big)
		res["name"] = xiangyan_big
	}
	if ty == "" {
		Db.Table("xiangyan").Where("name like ? AND ty=0", "%"+name+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan)
		Db.Table("xiangyan").Where("name like ? AND ty=1", "%"+name+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_big)
		list = append(list, xiangyan...)
		list = append(list, xiangyan_big...)
		res["name"] = list
	}

	tool.Success(c, "success", res)
}
func GetXiangyanByCode(c *gin.Context) {
	var xiangyan_small []model.Xiangyan
	var xiangyan_big []model.Xiangyan
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	code := c.Query("code")
	ty := c.Query("ty")
	offset := (pageNum - 1) * pagesize
	res := make(map[string]interface{})
	if ty == "0" {
		Db.Table("xiangyan").Where("code like ? AND ty=0", "%"+code+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_small)
	}

	if ty == "1" {
		Db.Table("xiangyan").Where("code like ? AND ty=1", "%"+code+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_big)

	}
	if ty == "" {
		Db.Table("xiangyan").Where("code like ? AND ty=0", "%"+code+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_small)
		Db.Table("xiangyan").Where("code like ? AND ty=1", "%"+code+"%").Offset(offset).Limit(pagesize).Order("sort desc").Find(&xiangyan_big)
	}

	res["code_small"] = xiangyan_small
	res["code_big"] = xiangyan_big
	tool.Success(c, "success", res)
}
func GetXiangYangBySaoma(c *gin.Context) {
	code := c.Query("code")
	var result model.Xiangyan
	Db.Table("xiangyan").Where("code like ?", "%"+code+"%").First(&result)
	tool.Success(c, "success", result)
}

func UpdateYanUrl(c *gin.Context) {
	var xiangyan []model.Xiangyan
	Db.Table("xiangyan").Find(&xiangyan)
	var count int = 0
	for _, v := range xiangyan {
		//img := v.Img
		row := Db.Table("xiangyan").Where("id=?", v.Id).Update("img", "http://www.etmoc.com"+v.Img)
		if row.RowsAffected == 1 {
			count += 1
		}
	}
	tool.Success(c, "success", count)
}

func GetDistence(c *gin.Context) {
	lat1 := c.Query("lat1")
	lgt1 := c.Query("lgt1")
	lat2 := c.Query("lat2")
	lgt2 := c.Query("lgt2")
	d := tool.EarthDistance(tool.String_float64(lat1), tool.String_float64(lgt1), tool.String_float64(lat2), tool.String_float64(lgt2))
	d2 := tool.GeoDistance(tool.String_float64(lat1), tool.String_float64(lgt1), tool.String_float64(lat2), tool.String_float64(lgt2), "K")
	r := make(map[string]interface{})
	r["d"] = d
	r["geo_d"] = d2
	tool.Success(c, "success", r)
}

type version struct {
	Id      int `gorm:"column:id" json:"id"`
	Version int `gorm:"column:version" json:"version"`
	Status  int `gorm:"column:status" json:"status"`
}

var res bool
var is_up int = 0
var down_url string = ""

func APPUpdate(c *gin.Context) {

	v := tool.String_int(c.Query("v"))
	vv := c.Query("v")
	plat := c.Query("plat")
	res = false
	var ver []version
	Db.Table("version").Where("status=?", 0).Find(&ver)
	if v < ver[0].Version {

		res = true
		is_up = 1
	} else {
		res = false
		is_up = 0
	}
	if plat == "android" {
		down_url = "https://www.kissnet.cn:39700/static/baogai.apk"
	} else {
		down_url = "https://www.kissnet.cn:39700/static/baogai.ipk" //appstore 的下载链接
	}
	m := make(map[string]interface{})
	m["update_flag"] = is_up //0 不需要升级，1 需要升级
	m["update_url"] = down_url
	m["forceupdate"] = 1 //0 不强制升级，1 强制升级
	m["wgt_flag"] = 0
	m["wgt_url"] = "https://www.kissnet.cn:39700/static/baogai.wgt"
	m["update_tips"] = "优化首页搜索问题\n优化运费余额支付\n优化卖货结算、提现功能\n优化金币结算、提现功能\n完善了会员充值\n优化各个订单详情查看"
	m["version"] = "1.0." + vv //当前版本
	m["size"] = 0
	m["system_version"] = tool.Int_string(ver[0].Version)
	m["input_version"] = vv
	// re := make(map[string]interface{})
	// re["code"] = 100
	// re["msg"] = ""
	// re["data"] = m
	tool.Success(c, "success", m)
}

func LoadPub(c *gin.Context) {
	var pub model.Pub
	Db.Table("version").Where("id=1").First(&pub)
	tool.Success(c, "success", pub)
}
func CheckCode(c *gin.Context) {
	code := c.Query("code")
	var user []model.UsersInfo
	Db.Table("users").Where("code=?", code).Find(&user)
	if len(user) > 0 { //找到邀请码
		tool.Success(c, "success", 200)
	} else {
		//无效邀请码
		tool.Success(c, "无效邀请码", 401)
	}
}

func CheckTel(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	ty := c.Query("ty")
	var u []model.UsersInfo
	Db.Table("users").Where("tel=? AND code=?", tel, code).Find(&u)
	if len(u) > 0 { //同一个人code和tel
		tool.Success(c, "不能推荐自己，请跟换", 400)
	} else {
		//判断是否被占用
		if ty == "tel" {
			Db.Table("users").Where("tel=?", tel).Find(&u)
			if len(u) > 0 {
				tool.Success(c, "手机号已被注册，请跟换", 400)
				return
			} else {
				tool.Success(c, "手机号可注册", 200)
			}

		}
		if ty == "email" {
			Db.Table("users").Where("email=?", tel).Find(&u)
			if len(u) > 0 {
				tool.Success(c, "邮箱已被注册，请跟换", 400)
				return
			} else {
				tool.Success(c, "手机号可注册", 200)
			}

		}

	}
}

func WebCheckSmsCode(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	var sms []model.Codes
	Db.Table("codes").Where("tel=? AND code=?", tel, code).Find(&sms)
	if len(sms) > 0 {
		tool.Success(c, "ok", 200)
	} else {
		tool.Success(c, "无效验证码", 400)
	}
}

func WebCheckEmailCode(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	var sms []model.EmailCodes
	Db.Table("email_codes").Where("email=? AND code=?", email, code).Find(&sms)
	if len(sms) > 0 {
		tool.Success(c, "ok", 200)
	} else {
		tool.Success(c, "无效验证码", 400)
	}
}

func CheckCodeNum(c *gin.Context) {
	// 1.初始化数据库

	// // 2.生成验证码
	// code := createCode()
	// // 3.将验证码存入redis,并统计当前用户请求验证码的次数
	// phone := c.Query("tel")
	// SetCode(code, phone)
	// // 4.判断用户输入的验证码是否正确
	// input := c.Query("code")
	// res := isEquals(input, "key"+input)
	// fmt.Println("res=", res)
}

// createCode 生成6位验证码
func createCode() (code string) {
	rand.Seed(time.Now().Unix()) //设置随机种子
	code = fmt.Sprintf("%6v", rand.Intn(600000))
	return
}

// isEquals 验证用户输入的验证码是否匹配
func isEquals(input, codeKey string, ctx context.Context) bool {
	// code := Rdb.Get(ctx, codeKey).String()
	// if code == input {
	// 	return true
	// }
	return false
}

// logic 用来处理请求部分
func SetCode(code, phone string, ctx context.Context) int {
	//	1.定义验证码和手机号的key的形式

	// codeKey := "key" + code
	// phoneKey := "key" + phone
	// //	2.查看用户发送验证码的数量
	// ctn := Rdb.Get(ctx, phoneKey).String()
	// //  ctn为零值的时候，说明该用户还没有发送过验证码
	// if ctn == "" {
	// 	Rdb.Set(ctx, phoneKey, "1", time.Minute*60*24) //设置手机请求验证码的时间为一天。
	// }
	// ctn1, _ := strconv.ParseInt(ctn, 10, 32)
	// //  发送次数没有超过3次
	// if ctn1 <= 2 {
	// 	Rdb.Incr(ctx, phoneKey)
	// 	return int(ctn1)
	// } else {
	// 	//fmt.Println("您今天发送的次数已经超过三次，不可再发送请求")

	// 	return int(ctn1)
	// }

	// //	将验证码存入redis,设置过期时间为2分钟
	// Rdb.Set(ctx, codeKey, code, time.Minute*2)
	return int(1)
}
