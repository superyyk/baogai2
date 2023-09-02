package api

import (
	"encoding/json"
	"errors"
	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/db"
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"
	"regexp"
	"time"

	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var Db = db.Db
var codes []model.Codes
var emailcodes []model.EmailCodes

func UserRegister(c *gin.Context) {
	res := make(map[string]interface{})
	res["ip"] = c.RemoteIP()
	tool.Success(c, "success", res)
}

func UserLogin(c *gin.Context) {
	tel := c.Query("tel")
	pass := c.Query("pass")
	pass = tool.Md5_salt(pass)

	ty := c.Query("ty")
	t := tool.String_int(ty)

	var userdata []model.UsersInfo

	if t == 0 {
		Db.Table("users").Where("tel=? AND pass=?", tel, pass).First(&userdata)
	}
	if t == 1 {
		Db.Table("users").Where("email=? AND pass=?", tel, pass).First(&userdata)
	}
	if len(userdata) != 0 {
		sign := utils.GetJWT(24*10, userdata[0].Uid, "yyk")
		res := make(map[string]interface{})
		res["sign"] = sign
		res["uid"] = userdata[0].Uid
		res["user_info"] = userdata
		tool.Success(c, "登陆成功！", res)

	} else {

		tool.Fail(c, "用户名或密码错误!", 400)
	}
}

func Regist(c *gin.Context) {
	tel := c.Query("tel")
	pass := c.Query("pass")
	pass = tool.Md5_salt(pass)
	pay_pass := tool.Md5_salt("000000")
	code := c.Query("code")
	ty := c.Query("ty")
	name := c.DefaultQuery("name", tel)
	is_code := c.Query("is_code")
	if is_code == "1" { //需要邀请码
		//校验邀请码
		var u []model.UsersInfo
		Db.Table("users").Where("code=?", code).Find(&u)
		if len(u) > 0 { //匹配到邀请码
			if ty == "tel" {
				err := Is_Tel(tel)
				if err != nil {
					tool.Success(c, "账号已存在，不能注册", 401)
					return
				} else {
					uuid := utils.Only()
					user_uuid := uuid
					co := utils.RandUid(7)
					// rr := tool.RandNum(12)
					// pass := tool.Md5(rr)
					user := model.UsersInfo{
						Tel:     tel,
						Name:    name,
						Uid:     user_uuid,
						Pass:    pass,
						Time:    tool.Int64_string(utils.GetNowTimeStamp()),
						Head:    config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
						Code:    co,
						PayPass: pay_pass,
					}
					if err := Db.Table("users").Create(&user).Error; err != nil {
						tool.Success(c, "系统繁忙，稍后再试", 401)
					} else {
						//添加推广关系
						err := AddTuiguang(code, u[0].Uid, co, user_uuid)
						if err != nil {
							tool.Success(c, "注册失败", 401)
						} else {
							tool.Success(c, "注册成功", 200)
						}

					}
				}

			}
			if ty == "email" {
				err := Is_Email(tel)
				if err != nil {
					tool.Success(c, "账号已存在，不能注册", 401)
					return
				} else {
					uuid := utils.Only()
					user_uuid := uuid
					co := utils.RandUid(7)
					// rr := tool.RandNum(12)
					// pass := tool.Md5(rr)
					user := model.UsersInfo{
						Email:   tel,
						Name:    name,
						Uid:     user_uuid,
						Pass:    pass,
						Time:    tool.Int64_string(utils.GetNowTimeStamp()),
						Head:    config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
						Code:    co,
						PayPass: pay_pass,
					}
					if err := Db.Table("users").Create(&user).Error; err != nil {
						tool.Success(c, "系统繁忙，稍后再试", 401)
					} else {
						//添加推广关系

						err := AddTuiguang(code, u[0].Uid, co, user_uuid)
						if err != nil {
							tool.Success(c, "注册失败", 401)
						} else {
							tool.Success(c, "注册成功", 200)
						}
					}
				}
			}

		} else {
			tool.Success(c, "无效邀请码", 401)
		}
	} else { //不需要邀请码
		uuid := utils.Only()
		user_uuid := uuid
		co := utils.RandUid(7)
		// rr := tool.RandNum(12)
		// pass := tool.Md5(rr)
		//user := model.UsersInfo{}
		if ty == "tel" {
			err := Is_Tel(tel)
			if err != nil {
				tool.Success(c, "账号已存在，不能注册", 401)
				return
			} else {
				user := model.UsersInfo{
					Tel:     tel,
					Name:    name,
					Uid:     user_uuid,
					Pass:    pass,
					Time:    tool.Int64_string(utils.GetNowTimeStamp()),
					Head:    config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
					Code:    co,
					PayPass: pay_pass,
				}
				if err := Db.Table("users").Create(&user).Error; err != nil {
					tool.Success(c, "系统繁忙，稍后再试", 401)
				} else {
					tool.Success(c, "注册成功", 200)
				}
			}

		}
		if ty == "email" {
			err := Is_Email(tel)
			if err != nil {
				tool.Success(c, "账号已存在，不能注册", 401)
				return
			} else {
				user := model.UsersInfo{
					Email:   tel,
					Name:    name,
					Uid:     user_uuid,
					Pass:    pass,
					Time:    tool.Int64_string(utils.GetNowTimeStamp()),
					Head:    config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
					Code:    co,
					PayPass: pay_pass,
				}
				if err := Db.Table("users").Create(&user).Error; err != nil {
					tool.Success(c, "系统繁忙，稍后再试", 401)
				} else {
					tool.Success(c, "注册成功", 200)
				}
			}

		}

	}
}

// 判断邮箱或手机号是否存在
func Is_Tel(tel string) error {
	var u []model.UsersInfo
	Db.Table("users").Where("tel=?", tel).Find(&u)
	if len(u) > 0 { //已存在
		return errors.New("已存在")
	} else {
		return nil
	}
}
func Is_Email(email string) error {
	var u []model.UsersInfo
	Db.Table("users").Where("email=?", email).Find(&u)
	if len(u) > 0 { //已存在
		return errors.New("已存在")
	} else {
		return nil
	}
}

func AddTuiguang(father, father_uid, child, child_uid string) error {
	saoma := &model.Saoma{
		Father:    father,
		Father_id: father_uid,
		Child:     child,
		Child_id:  child_uid,
		Time:      time.Now().Unix(),
	}
	if err := Db.Table("saoma").Create(&saoma).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func LoginSms(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	uid := c.Query("uid")
	var userdata []model.UsersInfo
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
			Db.Table("users").Where("tel=?", tel).First(&userdata)
			if len(userdata) != 0 { //找到用户

				user_uid := userdata[0].Uid
				sign := utils.GetJWT(24*10, user_uid, "yyk")
				res["sign"] = sign
				res["uid"] = user_uid
				res["user_info"] = userdata
				//CheckManage(c,user_uid)
				tool.Success(c, "登陆成功！", res)
			} else { //未找到，即注册新用户

				// uuid := utils.Only()
				// user_uuid := uuid
				// rr := tool.RandNum(12)
				// pass := tool.Md5(rr)
				// user := model.UsersInfo{
				// 	Tel:  tel,
				// 	Name: "游客",
				// 	Uid:  user_uuid,
				// 	Pass: pass,
				// 	Time: tool.Int64_string(utils.GetNowTimeStamp()),
				// 	Head: config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
				// 	Code: utils.RandUid(7),
				// }
				// if err := Db.Table("users").Create(&user).Error; err != nil { //注册失败
				tool.Fail(c, "未找到该用户！", 400)
				// } else { //注册成功
				// 	sign := utils.GetJWT(24*10, user_uuid, "yyk")
				// 	//user_uid:=user_uuid
				// 	res["sign"] = sign
				// 	res["uid"] = user_uuid
				// 	//CheckManage(c,user_uuid)
				// 	tool.Success(c, "注册并登陆成功！", res)
				// }
			}
		}
	} else { //验证码不匹配
		tool.Fail(c, "手机号与验证码不符", 400)

	}
}

func LoginEmail(c *gin.Context) {
	email := c.Query("tel")
	code := c.Query("code")
	uid := c.Query("uid")
	var userdata []model.UsersInfo
	res := make(map[string]interface{})

	Db.Table("email_codes").Where("email=? AND code=? AND uid=?", email, code, uid).First(&codes)
	if len(codes) != 0 { //验证码匹配正确
		//验证码是否过期
		nowtime := time.Now().Unix()

		if nowtime-codes[0].Time > 60*5 { //5分钟后失效
			tool.Fail(c, "验证码已失效", nowtime-codes[0].Time)
			return
		} else {

			//查看是否有该用户
			Db.Table("users").Where("email=?", email).First(&userdata)
			if len(userdata) != 0 { //找到用户

				user_uid := userdata[0].Uid
				sign := utils.GetJWT(24*10, user_uid, "yyk")
				res["sign"] = sign
				res["uid"] = user_uid
				res["user_info"] = userdata
				//CheckManage(c,user_uid)
				tool.Success(c, "登陆成功！", res)
			} else { //未找到，即注册新用户
				// uuid := utils.Only()
				// user_uuid := uuid
				// rr := tool.RandNum(12)
				// pass := tool.Md5(rr)
				// code := utils.RandUid(7)
				// user := model.UsersInfo{
				// 	Email: email,
				// 	Name:  "游客",
				// 	Uid:   user_uuid,
				// 	Pass:  pass,
				// 	Time:  tool.Int64_string(utils.GetNowTimeStamp()),
				// 	Head:  config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/static/logo.png",
				// 	Code:  code,
				// }
				// if err := Db.Table("users").Create(&user).Error; err != nil { //注册失败
				tool.Fail(c, "未找到该用户！", 400)
				// } else { //注册成功
				// 	sign := utils.GetJWT(24*10, user_uuid, "yyk")
				// 	//user_uid:=user_uuid
				// 	res["sign"] = sign
				// 	res["uid"] = user_uuid
				// 	//CheckManage(c,user_uuid)
				// 	tool.Success(c, "注册并登陆成功！", res)
				// }
			}
		}
	} else { //验证码不匹配
		tool.Fail(c, "邮箱账号与验证码不符", 400)

	}
}

func ChangePass(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	uid := c.Query("uid")
	pass := c.Query("pass")
	ty := c.Query("ty")
	t := tool.String_int(ty)
	if t == 0 {

		Db.Table("codes").Where("tel=? AND code=? AND uid=?", tel, code, uid).First(&codes)
		if len(codes) != 0 { //成功匹配
			pass = tool.Md5_salt(pass)
			res := Db.Table("users").Where("tel=?", tel).Update("pass", pass)
			if res.RowsAffected == 1 { //修改成功
				tool.Success(c, "修改成功！", 200)
			} else {
				tool.Fail(c, "修改失败！", codes)
			}
		} else {
			tool.Fail(c, "手机验证码不符", 400)
		}
	}
	if t == 1 {

		Db.Table("email_codes").Where("email=? AND code=? AND uid=?", tel, code, uid).First(&emailcodes)
		if len(emailcodes) != 0 { //成功匹配
			pass = tool.Md5_salt(pass)
			res := Db.Table("users").Where("email=?", tel).Update("pass", pass)
			if res.RowsAffected == 1 { //修改成功
				tool.Success(c, "修改成功！", 200)
			} else {
				tool.Fail(c, "修改失败！", codes)
			}
		} else {
			tool.Fail(c, "邮箱验证码不符", 400)
		}
	}

}

func GetUserInfo(c *gin.Context) {
	uid := c.Query("user_uuid")
	var sys_add []model.Address
	var my_add []model.Address
	var userdata []model.UsersInfo
	var wuliu_comp []model.WuliuComp
	var vip []model.Vip
	var dabao []model.DaBao
	var version []model.Version
	var saoma model.Saoma
	var saoma_childs []model.Saoma
	var tixian []model.Tixian
	var gonggao []model.Gonggao
	res := make(map[string]interface{})
	tx := Db.Begin()
	tx.Table("users").Where("uid=?", uid).Find(&userdata)
	tx.Table("address").Where("ty=? AND status=? AND used=?", 1, 0, 1).Find(&sys_add)
	tx.Table("address").Where("user_uuid=? AND status=?", uid, 0).Find(&my_add)
	tx.Table("wuliu_comp").Where("status=?", 0).Find(&wuliu_comp)
	tx.Table("vip").Where("status=?", 0).Find(&vip)
	tx.Table("dabao").Where("user_uuid=?", uid).Find(&dabao)
	tx.Table("version").Where("id=?", 1).Find(&version)
	tx.Table("saoma").Where("child_id=?", uid).First(&saoma)
	tx.Table("saoma").Where("father_id=?", uid).Find(&saoma_childs)
	tx.Table("tixian").Where("user_uuid=?", uid).Order("id desc").Find(&tixian)
	tx.Table("gonggao").Where("status=0").Find(&gonggao)
	for k, v := range dabao {
		ids := v.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)

		var list []model.ChuShou
		for _, kk := range nums {
			var chushou []model.ChuShou
			Db.Table("chushou").Where("order_id=?", kk).Find(&chushou)
			list = append(list, chushou[0])
		}
		l, _ := json.Marshal(list)
		dabao[k].Content = string(l)

	}

	yue := GetUseryue(uid)
	jinbi_yue := GetUserJinbiYue(uid)
	shangji := GetShangYiji(uid)
	if len(userdata) != 0 { //找到用户
		tx.Commit()
		res["info"] = userdata
		res["sys_add"] = sys_add
		res["my_add"] = my_add
		res["wuliu_comp"] = wuliu_comp
		res["vip"] = vip
		res["dabao"] = dabao
		res["huokuan"] = GetHuokuan(uid)
		res["version"] = version
		res["saoma_father"] = saoma
		res["saoma_child"] = saoma_childs
		res["yue"] = yue
		res["jinbi_yue"] = jinbi_yue
		res["tixian"] = tixian
		res["shangji"] = shangji
		res["gonggao"] = gonggao
		tool.Success(c, "success", res)
	} else {
		tx.Rollback()
		tool.Fail(c, "未找到该用户信息", 400)
	}

}
func GetShangYiji(user_uuid string) model.UsersInfo {
	var saoma model.Saoma
	Db.Table("saoma").Where("child_id=?", user_uuid).Take(&saoma)
	var info model.UsersInfo
	Db.Table("users").Select("uid,name,weixin_code,weixin,qq").Where("uid=?", saoma.Father_id).Take(&info)
	return info

}

// 获取用户账户余额
func GetUseryue(user_uuid string) float64 {
	//自己的卖货收益，除去上级和上二级提成

	maihuo_shouru := GetMaihuoShouru(user_uuid)
	//下级和下二级的卖货提成
	yiji := YijiMaiHuoTiCheng(user_uuid)
	erji := ErjiMaiHuoTiCheng(user_uuid)
	//已提现
	tixian := GetUserTixian(user_uuid)
	yue := maihuo_shouru + yiji + erji - tixian

	return yue

}
func GetUserTixian(user_uuid string) float64 {
	var count float64 = 0
	var result []float64
	Db.Table("tixian").Where("user_uuid=? AND status in (0,1) AND ty=0", user_uuid).Pluck("price", &result)
	for _, v := range result {
		count += v
	}
	return count
}
func GetUserJinbiYue(user_uuid string) float64 {
	//下级和下二级大于7天的VIP会员 充值提成
	yiji := GetYijiVipTiCheng(user_uuid)
	erji := GetErjiVipTiCheng(user_uuid)
	//已提现
	tibi := GetJinbiTiBi(user_uuid)
	//素材购买
	sucai := GetSucaiPay(user_uuid)
	return yiji + erji - tibi - sucai
}

func GetSucaiPay(user_uuid string) float64 {
	var sucai_pay float64 = 0
	var result []float64
	Db.Table("buy_sucai").Where("user_uuid=? AND status=?", user_uuid, 1).Pluck("money", &result)
	for _, v := range result {
		sucai_pay += v
	}
	return sucai_pay
}
func GetJinbiTiBi(user_uuid string) float64 {
	var tibi float64 = 0
	var result []float64
	Db.Table("tixian").Where("user_uuid=? AND status in (0,1) AND ty=1", user_uuid).Pluck("price", &result)
	for _, v := range result {
		tibi += v
	}
	return tibi
}

func GetErjiVipTiCheng(user_uuid string) float64 {
	var saoma []model.Saoma
	var erji_ticheng float64 = 0
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for _, v := range saoma {
		var erji []model.Saoma
		Db.Table("saoma").Where("father_id=?", v.Child_id).Find(&erji)
		for _, vv := range erji {
			var vip_order model.VipOrders
			Db.Table("vip_orders").Where("user_uuid=? AND status=1", vv.Child_id).First(&vip_order)
			now := time.Now().Unix()
			d_7 := 7 * 24 * 60 * 60
			if (now - int64(d_7)) > vip_order.PayTime { //7天后
				erji_ticheng += tool.String_float64(vip_order.Price) * float64(vip_order.T_2) / 100
			}
		}

	}
	return erji_ticheng
}

func GetYijiVipTiCheng(user_uuid string) float64 {
	var saoma []model.Saoma
	var vip_ticheng float64 = 0
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for _, v := range saoma {
		var vip_order model.VipOrders
		Db.Table("vip_orders").Where("user_uuid=? AND status=1", v.Child_id).First(&vip_order)
		now := time.Now().Unix()
		d_7 := 7 * 24 * 60 * 60
		if (now - int64(d_7)) > vip_order.PayTime { //7天后
			vip_ticheng += tool.String_float64(vip_order.Price) * float64(vip_order.T_1) / 100
		}
	}
	return vip_ticheng
}

func YijiMaiHuoTiCheng(user_uuid string) float64 {
	var saoma []model.Saoma
	var yiji_maihuo_ticheng float64 = 0
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for _, v := range saoma {
		yiji_maihuo_ticheng += GetMaihuoTiCheng(v.Child_id, 1)
	}
	return yiji_maihuo_ticheng
}

func ErjiMaiHuoTiCheng(user_uuid string) float64 {
	var saoma []model.Saoma
	var erji_maihuo_ticheng float64 = 0
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for _, v := range saoma {
		//获取二级用户
		var erji []model.Saoma
		Db.Table("saoma").Where("father_id=?", v.Child_id).Find(&erji)
		for _, vv := range erji {
			erji_maihuo_ticheng += GetMaihuoTiCheng(vv.Child_id, 2)
		}
	}
	return erji_maihuo_ticheng
}

func GetMaihuoShouru(user_uuid string) float64 {
	var dabao []model.DaBao
	var maichu_money float64 = 0
	Db.Table("dabao").Where("user_uuid=? AND status=?", user_uuid, 5).Find(&dabao)
	for _, v := range dabao {
		ids := v.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)
		for _, vv := range nums {
			var maichu model.ChuShou
			Db.Table("chushou").Where("order_id=?", vv).First(&maichu)
			mm := float64(maichu.OkNum) * tool.String_float64(maichu.Price)
			yiji := mm * float64(maichu.M_1) / 100
			erji := mm * float64(maichu.M_2) / 100
			maichu_money += mm - yiji - erji
		}

	}
	return maichu_money
}

func GetMaihuoTiCheng(user_uuid string, ty int) float64 {
	var dabao []model.DaBao
	var ticheng float64 = 0
	Db.Table("dabao").Where("user_uuid=? AND status=?", user_uuid, 5).Find(&dabao)
	for _, v := range dabao {
		ids := v.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)
		for _, vv := range nums {
			var maichu model.ChuShou
			Db.Table("chushou").Where("order_id=?", vv).First(&maichu)
			mm := float64(maichu.OkNum) * tool.String_float64(maichu.Price)
			yiji := mm * float64(maichu.M_1) / 100
			erji := mm * float64(maichu.M_2) / 100
			if ty == 1 {
				ticheng += yiji
			}
			if ty == 2 {
				ticheng += erji
			}
			//maichu_ticheng += yiji + erji
		}

	}
	return ticheng
}

// 计算货款
func GetHuokuan(user_uuid string) float64 {
	var dabao []model.DaBao
	//tx := Db.Begin()
	Db.Table("dabao").Where("user_uuid=?", user_uuid).Find(&dabao)
	var count float64 = 0
	var list []model.ChuShou
	for _, v := range dabao {
		if v.Status == "5" {
			ids := v.Orders_id
			re := regexp.MustCompile("[0-9]+")
			nums := re.FindAllString(ids, -1)
			for _, kk := range nums {
				var chushou []model.ChuShou
				Db.Table("chushou").Where("order_id=?", kk).Find(&chushou)
				list = append(list, chushou[0])

			}
		}

	}
	for _, vv := range list {

		count += float64(vv.OkNum) * tool.String_float64(vv.Price)
	}

	return count
}

func UseAddress(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	id := c.Query("id")
	//使用中的地址
	var add []model.Address
	Db.Table("address").Where("user_uuid=? AND id=? AND used=?", user_uuid, id, 1).Find(&add)
	if len(add) > 0 { //已经被使用
		tool.Success(c, "正在使用，无须设置!", 401)
	} else {
		row := Db.Table("address").Where("user_uuid=? AND used=?", user_uuid, 1).Update("used", 0)
		if row.RowsAffected == 1 {
			ro := Db.Table("address").Where("user_uuid=? AND id=?", user_uuid, id).Update("used", 1)
			if ro.RowsAffected == 1 {
				tool.Success(c, "设置成功！", 200)
			} else {
				tool.Success(c, "设置失败！", 401)
			}
		} else {
			tool.Success(c, "设置失败！", 401)
		}
	}

}

func AddAddress(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	address := c.Query("address")
	lat := c.Query("lat")
	lgt := c.Query("lgt")
	name := c.Query("name")
	tel := c.Query("tel")
	tip := c.Query("tip")
	ty := c.Query("ty")
	id := c.Query("id")
	var add []model.Address
	Db.Table("address").Where("user_uuid=?", user_uuid).Find(&add)
	if len(add) > 0 { //不是首次添加
		if ty == "0" { //添加
			data := &model.Address{
				Name:      name,
				Tel:       tel,
				Address:   address,
				User_uuid: user_uuid,
				Tip:       tip,
				Lat:       tool.String_float64(lat),
				Lgt:       tool.String_float64(lgt),
			}
			if err := Db.Table("address").Create(&data).Error; err != nil {
				tool.Success(c, "系统繁忙，稍后再试！", 401)
				return
			} else {
				tool.Success(c, "添加成功！", 200)
			}
		}

		if ty == "1" { //编辑
			row := Db.Table("address").Where("user_uuid=? AND id=?", user_uuid, id).Updates(map[string]interface{}{
				"address": address,
				"name":    name,
				"tip":     tip,
				"tel":     tel,
			})
			if row.RowsAffected == 1 {
				tool.Success(c, "编辑成功！", 200)
			} else {
				tool.Success(c, "编辑失败！", 401)
			}
		}

		if ty == "2" { //删除
			row := Db.Table("address").Where("user_uuid=? AND id=?", user_uuid, id).Update("status", 1)
			if row.RowsAffected == 1 {
				tool.Success(c, "删除成功！", 200)
			} else {
				tool.Success(c, "删除失败！", 401)
			}
		}
	} else { //首次添加设置默认
		data := &model.Address{
			Name:      name,
			Tel:       tel,
			Address:   address,
			User_uuid: user_uuid,
			Tip:       tip,
			Used:      1,
		}
		if err := Db.Table("address").Create(&data).Error; err != nil {
			tool.Success(c, "系统繁忙，稍后再试！", 401)
			return
		} else {
			tool.Success(c, "添加成功！", 200)
		}

	}

}

func get_clube_menber(clube_id string) int {
	var count int
	Db.Table("clubs_joins").Where("clube_id=? AND status=?", clube_id, 1).Count(&count)
	return count
}

func CreateClube(c *gin.Context) {
	//var clube []model.Cludes
	user_uuid := c.Query("user_uuid")
	name := c.Query("name")
	head := c.Query("head")
	tip := c.Query("tip")
	lat := c.Query("lat")
	lgt := c.Query("lgt")
	address := c.Query("address")
	t := time.Now().Unix()
	uid := utils.Only()
	u := uid

	tx := Db.Begin()
	clube := model.Cludes{
		Name:    name,
		Puber:   user_uuid,
		Time:    t,
		Head:    head,
		Uid:     u,
		Tip:     tip,
		Count:   1,
		Lat:     lat,
		Lgt:     lgt,
		Address: address,
	}
	if err := tx.Table("clubs").Create(&clube).Error; err != nil {

		tx.Rollback()
		tool.Fail(c, "创建失败", 400)
		return
	} else {
		//加入自己的俱乐部
		jc := model.JoinCludes{
			Clube:  u,
			Ask:    user_uuid,
			Time:   time.Now().Unix(),
			Pass:   tool.Int64_string(time.Now().Unix()),
			Status: 1,
			Tip:    "自动同意",
		}
		if er := tx.Table("clubs_joins").Create(&jc).Error; er != nil {
			tx.Rollback()
			tool.Fail(c, "创建失败", 400)
			return
		}
		tx.Commit()
		tool.Success(c, "创建成功", 200)
	}

}

func GetClude(c *gin.Context) {
	user := c.Query("user_uuid")
	var clubs []model.Cludes
	var clubs_in []model.JoinCludes
	//    我创建的俱乐部
	Db.Table("clubs").Where("puber=? AND status=?", user, 0).Order("id desc").Find(&clubs)
	for key, v := range clubs {
		var menber int
		var shenqing int
		var jc_ok []model.JoinCludes
		var jc_shenqing []model.JoinCludes
		Db.Table("clubs_joins").Where("clube_id=? AND status=?", v.Uid, 1).Find(&jc_ok).Count(&menber)
		Db.Table("clubs_joins").Where("clube_id=? AND status=?", v.Uid, 0).Find(&jc_shenqing).Count(&shenqing)
		clubs[key].Count = menber
		clubs[key].Shenqing = shenqing

		for k, vv := range jc_ok {
			var user []model.UserData
			Db.Table("users").Where("uid=?", vv.Ask).Find(&user)
			jc_ok[k].Info = user
		}
		for k, vv := range jc_shenqing {
			var user []model.UserData
			Db.Table("users").Where("uid=?", vv.Ask).Find(&user)
			jc_shenqing[k].Info = user
		}
		clubs[key].Ok = jc_ok
		clubs[key].Shen = jc_shenqing
	}
	//我加入的俱乐部
	Db.Table("clubs_joins").Where("asker_id=?", user).Order("id desc").Find(&clubs_in)
	for k, v := range clubs_in {
		var clube_info model.Cludes
		Db.Table("clubs").Where("uid=?", v.Clube).First(&clube_info)
		clubs_in[k].CludesInfo = clube_info
	}
	res := make(map[string]interface{})
	res["clubs"] = clubs
	res["clubs_in"] = clubs_in
	tool.Success(c, "success", res)
}

func SearchClude(c *gin.Context) {
	//user_uuid := c.Query("user_uuid")
	clube_id := c.Query("clube_id")
	var clube []model.Cludes
	Db.Table("clubs").Where("uid=?", clube_id).Find(&clube)

	if len(clube) != 0 {
		for key, v := range clube {
			var menber int
			Db.Table("clubs_joins").Where("clube_id=? AND status=?", v.Uid, 1).Count(&menber)
			clube[key].Count = menber
		}
		tool.Success(c, "success", clube)
	} else {
		tool.Fail(c, "俱乐部不存在", 400)
	}

}

func JoinClude(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	clube_id := c.Query("clube_id")
	tip := c.Query("tip")
	var join_clube []model.JoinCludes
	//是否申请过
	Db.Table("clubs_joins").Where("clube_id=? AND asker_id=? AND status!=?", clube_id, user_uuid, 2).Find(&join_clube)

	if len(join_clube) > 0 { // 已经加入了
		tool.Fail(c, "已申请，无须重复申请!", 400)
	} else {
		jc := model.JoinCludes{
			Clube: clube_id,
			Ask:   user_uuid,
			Tip:   tip,
			Time:  time.Now().Unix(),
		}
		if err := Db.Table("clubs_joins").Create(&jc).Error; err != nil {
			tool.Fail(c, "申请失败！", 400)
			return
		} else {
			tool.Success(c, "申请成功！", 200)
		}
	}

}

func ClubePass(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	clube_id := c.Query("clube_id")
	ty := c.Query("ty")
	tx := Db.Begin()
	res := tx.Table("clubs_joins").Where("clube_id=? AND asker_id=?", clube_id, user_uuid).Update("status", ty)
	var cc int
	var clube model.Cludes
	tx.Table("clubs").Where("uid=?", clube_id).First(&clube)
	cc = clube.Count
	if res.RowsAffected == 1 {
		tx.Table("clubs").Where("uid=?", clube_id).Update("count", cc+1)
		tx.Commit()
		tool.Success(c, "操作成功！", 200)
	} else {
		tx.Rollback()
		tool.Fail(c, "操作失败", 400)
	}

}

func GetCode(c *gin.Context) {
	qrcode.WriteFile("https://blog.csdn.net/yang731227", qrcode.High, 200, "./qrcode.png")
	qrcode.WriteColorFile("https://blog.csdn.net/yang731227", qrcode.High, 256, color.Black, color.White, "./qrcode.png")
}
