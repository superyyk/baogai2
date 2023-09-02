package api

import (
	"flag"
	"github.com/superyyk/baogai/model"
	"github.com/superyyk/baogai/redis"
	"github.com/superyyk/baogai/tool"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetYunfeiTuiInfo(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var wuliu []model.YunfeiFanhuanInfo
	Db.Table("yunfei_fanhuan_info").Where("user_uuid=? AND ty=2", user_uuid).Find(&wuliu)
	for k, v := range wuliu {
		var baoguo model.DaBao
		Db.Table("dabao").Where("baoguo_id=?", v.BaoguoId).First(&baoguo)
		wuliu[k].Baoguo_info = baoguo
	}
	tool.Success(c, "success", wuliu)

}

func GetYunFeiYue(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	yue := GetYunfei_yue(user_uuid)
	tool.Success(c, "success", yue)
}

func ChangePayPass(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	pass := c.Query("password")
	pass = tool.Md5_salt(pass)
	row := Db.Table("users").Where("uid=?", user_uuid).Update("pay_pass", pass)
	if row.RowsAffected == 1 {
		tool.Success(c, "修改成功", 200)
	} else {
		tool.Success(c, "修改失败", 401)
	}
}

func UserTixianYue(c *gin.Context) {

	user_uuid := c.Query("user_uuid")
	money := c.Query("money")
	pass := c.Query("pass")
	ty := c.Query("ty")
	pass = tool.Md5_salt(pass)
	m := tool.String_float64(money)
	yue := GetUseryue(user_uuid)
	jinbi_yue := GetUserJinbiYue(user_uuid)
	if ty == "0" {
		if m <= yue { //余额不足
			//判断密码
			var u []model.UsersInfo
			Db.Table("users").Where("uid=? AND pay_pass=?", user_uuid, pass).Find(&u)
			if len(u) > 0 {
				tixian := &model.Tixian{
					Time:      time.Now().Unix(),
					Ty:        tool.String_int(ty),
					User_uuid: user_uuid,
					Price:     m,
				}
				if err := Db.Table("tixian").Create(&tixian).Error; err != nil {
					tool.Success(c, "系统繁忙，稍后再试", 401)
				} else {
					tool.Success(c, "申请成功", 200)
				}
			} else {
				tool.Success(c, "密码错误", 401)
			}

		} else {
			tool.Success(c, "可用余额不足", 401)
		}
	} else {

		if m <= jinbi_yue { //余额不足
			//判断密码
			var u []model.UsersInfo
			Db.Table("users").Where("uid=? AND pay_pass=?", user_uuid, pass).Find(&u)
			if len(u) > 0 {
				tixian := &model.Tixian{
					Time:      time.Now().Unix(),
					Ty:        tool.String_int(ty),
					User_uuid: user_uuid,
					Price:     m,
				}
				if err := Db.Table("tixian").Create(&tixian).Error; err != nil {
					tool.Success(c, "系统繁忙，稍后再试", 401)
				} else {
					tool.Success(c, "申请成功", 200)
				}
			} else {
				tool.Success(c, "密码错误", 401)
			}

		} else {
			tool.Success(c, "可用余额不足", 401)
		}
	}

}

func FanKui(c *gin.Context) {
	tel := c.Query("tel")
	content := c.Query("content")
	uid := c.Query("uid")
	user_uuid := c.Query("user_uuid")
	fk := &model.FanKui{
		User_uuid: user_uuid,
		Tel:       tel,
		Content:   content,
		Uid:       uid,
		Time:      time.Now().Unix(),
	}
	if err := Db.Table("fankui").Create(&fk).Error; err != nil {
		tool.Success(c, "反馈失败", 401)
	} else {
		tool.Success(c, "反馈成功", 200)
	}
}

func ChangePhone(c *gin.Context) {
	tel := c.Query("tel")
	user_uuid := c.Query("user_uuid")
	code := c.Query("code")
	code_uid := c.Query("code_uid")
	var sms []model.Codes
	Db.Table("codes").Where("uid=? AND code=? AND tel=?", code_uid, code, tel).Find(&sms)
	if len(sms) > 0 { //匹配成功
		row := Db.Table("users").Where("uid=?", user_uuid).Update("tel", tel)
		if row.RowsAffected == 1 {
			tool.Success(c, "修改成功", 200)
		} else {
			tool.Success(c, "修改失败", 401)
		}
	} else {
		tool.Success(c, "手机号和验证码不匹配", 401)
	}

}

func Checktel(c *gin.Context) {
	tel := c.Query("tel")
	var u []model.UsersInfo
	Db.Table("users").Where("tel=?", tel).Find(&u)
	if len(u) > 0 {
		tool.Success(c, "手机号已被注册，请更换", 401)
	} else {
		tool.Success(c, "可以注册", 200)
	}
}

func SettingInfo(c *gin.Context) {
	ty := c.Query("ty")
	qq := c.Query("qq")
	weixin := c.Query("weixin")
	weixin_code := c.Query("weixin_code")
	shoukuan_code := c.Query("shoukuan_code")
	name := c.Query("name")
	tixian_ty := c.Query("tixian_ty")
	user_uuid := c.Query("user_uuid")
	if ty == "0" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("qq", qq)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
	if ty == "1" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("weixin", weixin)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
	if ty == "2" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("weixin_code", weixin_code)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
	if ty == "3" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("shoukuan_code", shoukuan_code)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
	if ty == "4" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("name", name)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
	if ty == "5" {
		row := Db.Table("users").Where("uid=?", user_uuid).Update("tixian_ty", tixian_ty)
		if row.RowsAffected == 1 {
			tool.Success(c, "设置成功", 200)
		} else {
			tool.Success(c, "设置失败", 401)
		}
	}
}

func TestSaoma(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var u model.Users
	Db.Preload("Saoma").Where(&model.Users{Uid: user_uuid}).First(&u)
	tool.Success(c, "suceess", u)
}

func GetJiaoCheng(c *gin.Context) {
	var jiaocheng []model.JiaoCheng
	Db.Table("jiaocheng").Where("status=?", 0).Order("id asc").Find(&jiaocheng)
	tool.Success(c, "success", jiaocheng)
}

func GetSucai(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	offset := (pageNum - 1) * pagesize
	yue := GetUserJinbiYue(user_uuid)
	var sucai []model.Sucai
	var buy_sucai []model.BySucai
	Db.Table("sucai").Where("status=?", 0).Offset(offset).Limit(pagesize).Order("id desc").Find(&sucai)
	Db.Table("buy_sucai").Select("sucai_id,money").Where("user_uuid=? AND status=1", user_uuid).Find(&buy_sucai)
	res := make(map[string]interface{})
	res["yue"] = yue
	res["sucai"] = sucai
	res["buy_sucai"] = buy_sucai
	tool.Success(c, "success", res)
}

func GetMySucai(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var sucai []model.Sucai
	Db.Table("sucai").Where("status=? AND user_uuid=?", 0, user_uuid).Order("id desc").Find(&sucai)
	tool.Success(c, "success", sucai)
}

func PubSucai(c *gin.Context) {
	ty := c.Query("ty")
	content := c.Query("content")
	urls := c.Query("urls")
	user_uuid := c.Query("user_uuid")
	price := c.Query("price")
	sc := &model.Sucai{
		User_uuid: user_uuid,
		Ty:        tool.String_int(ty),
		Time:      time.Now().Unix(),
		Content:   content,
		Price:     tool.String_float64(price),
		Urls:      urls,
	}
	if err := Db.Table("sucai").Create(&sc).Error; err != nil {
		tool.Success(c, "发布失败", 401)
	} else {
		tool.Success(c, "发布成功", 200)
	}
}
func SetSucai(c *gin.Context) {
	id := c.Query("id")
	status := c.Query("status")
	row := Db.Table("sucai").Where("id=?", id).Update("status", status)
	if row.RowsAffected == 1 {
		tool.Success(c, "设置成功", 200)
	} else {
		tool.Success(c, "设置失败", 401)
	}
}

func SucaiPay(c *gin.Context) {
	sucai_id := c.Query("sucai_id")
	user_uuid := c.Query("user_uuid")
	money := c.Query("money")
	pass := c.Query("pass")
	pass = tool.Md5_salt(pass)
	var user []model.UsersInfo
	Db.Table("users").Where("uid=? AND pay_pass=?", user_uuid, pass).First(&user)
	if len(user) > 0 {
		buy_sucai := &model.BySucai{
			Money:     tool.String_float64(money),
			User_uuid: user_uuid,
			Sucai_id:  tool.String_int(sucai_id),
			Time:      time.Now().Unix(),
			Status:    1,
		}
		if err := Db.Table("buy_sucai").Create(&buy_sucai).Error; err != nil {
			tool.Success(c, "购买失败", 401)
		} else {

			tool.Success(c, "购买成功", 200)
		}
	} else {
		tool.Success(c, "密码错误！", 401)
	}

}

func GetMyBuySucai(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var my_buy []model.BySucai
	Db.Table("buy_sucai").Where("user_uuid=?", user_uuid).Find(&my_buy)
	for k, v := range my_buy {
		var sucai model.Sucai
		Db.Table("sucai").Where("id=?", v.Sucai_id).Take(&sucai)
		my_buy[k].Info = sucai
	}
	tool.Success(c, "success", my_buy)
}

func GetBaoGuoList(c *gin.Context) {
	var baoguo []model.DaBao
	user_uuid := c.Query("user_uuid")
	status := c.Query("status")
	if status == "10" {
		Db.Table("dabao").Where("user_uuid=? AND status in (0,1,2,3,4,5,6)", user_uuid).Find(&baoguo)
	} else {
		Db.Table("dabao").Where("user_uuid=? AND status=?", user_uuid, status).Find(&baoguo)
	}

	tool.Success(c, "success", baoguo)
}

func TestMysql(c *gin.Context) {
	//val := c.Query("val")
	v := c.Request.URL
	rdb := redis.Rdb
	ip := c.ClientIP()

	ip_key := "key_" + ip
	ctn, _ := rdb.Get(ip_key).Result()
	res := make(map[string]interface{})
	if ctn == "" {
		rdb.Set(ip_key, "1", time.Hour*24) //一小时内访问超5次
	}
	ctn1, _ := strconv.ParseInt(ctn, 10, 64)
	if ctn1 <= 500 {
		pipe := rdb.TxPipeline()
		pipe.Incr(ip_key) //累加+1
		err := pipe.LPush(ip, v).Err()
		pipe.Expire(ip_key, time.Second*60)
		if err != nil {
			tool.Fail(c, "fail", 401)
			return
		}
		pipe.Expire(ip, time.Second*10)
		_, _ = pipe.Exec()
		tool.Success(c, "success", ctn)
	} else {
		ip_len := rdb.LLen(ip_key).Val()
		//ip_val := rdb.LPop(ip).Val()
		res["code"] = 401
		res["ctn"] = ctn
		res["ip_len"] = ip_len
		res["ip_val"] = ""
		tool.Fail(c, "当天IP访问次数已超过三次", res)
		return
	}

}

var (
	//etcdAddr = flag.String("etcd_addr", "47.106.160.38:2379", "etcd address")
	addr = flag.String("base_addr", "47.106.160.38:39100", "base address")
)

func RpcxTest(c *gin.Context) {
	num := c.Query("num")
	//n := tool.String_int(num)
	p := tool.Producer
	err := p.Publish("test", num)
	if err != nil {
		tool.Success(c, "faid", err)
	} else {
		tool.Success(c, "success", 200)
	}
}
