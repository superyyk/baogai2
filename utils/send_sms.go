package utils

import (
	"encoding/json"
	"fmt"
	"github.com/superyyk/baogai/config"
	"github.com/superyyk/baogai/db"
	"github.com/superyyk/baogai/model"
	"github.com/superyyk/baogai/tool"
	"time"

	"github.com/gin-gonic/gin"
)

type CodeRes struct {
	Res  int
	Uid  string
	Code int
	Data string
}

type Email struct {
	Code  string `gorm:"column:code" json:"code"`
	Uid   string `gorm:"column:uid" json:"uid"`
	Email string `gorm:"column:email" json:"email"`
	Time  int64  `gorm:"column:time" json:"time"`
}

var Db = db.Db

func SendSms(c *gin.Context) {

	tel := c.Query("tel")
	sign := c.Query("sign")
	//res:=make(map[string]interface{})
	var coderes = &CodeRes{}

	url := "http://47.106.160.38:39001/index.php/index/index/send_sign_sms?tel=" + tel + "&sign=" + sign
	str := HttpGet(url)
	if err := json.Unmarshal([]byte(str), coderes); err != nil {
		fmt.Println(err)
	}
	if coderes.Res == 200 {
		var codes = model.Codes{
			Uid:  coderes.Uid,
			Tel:  tel,
			Code: coderes.Code,
			Time: GetNowTimeStamp(),
			//Date:GetNowTime(),

		}
		if err := Db.Table("codes").Create(&codes).Error; err != nil {

			tool.Success(c, "短信发送成功！存库失败", coderes.Uid)
		} else {
			tool.Success(c, "短信验证码已发送！", coderes.Uid)
		}

	} else {
		tool.Fail(c, "系统繁忙！稍后再试", 400)
	}

}

func SendEmail(c *gin.Context) {
	serverHost := "smtp.qq.com"
	serverPort := 587
	fromEmail := "1432915284@qq.com" //
	fromPasswd := "ilacnmcyvbefgiij" ////QLCVMKUTEQWLLGIE
	myToers := c.Query("to_emails")
	//myCCers:="y18974197146@163.com,1432915284@qq.com" //抄送
	subject := c.Query("title")

	ty := SmsNum4()
	uid := RandNum(21)
	//body := `这是正文<br>
	//        <h3>这是标题</h3>
	//         Hello <a href = "http://www.latelee.org">主页</a><br>`
	a := `<h3> Email verification code：<a>` + ty + `</a></h3> <p><span style="color:green;font-size:large"> 【` + config.Base.EmailTitle + `平台】线上回收，一个就收！</span></p>`
	body := a
	//方法一
	//myEmail:=&tool.EmailParam{
	//	ServerHost:serverHost,
	//	ServerPort:serverPort,
	//	FromEmail:fromEmail,
	//	FromPasswd:fromPasswd,
	//	Toers:myToers,
	//	CCers:myCCers,
	//}
	//tool.InitEmail(myEmail)
	//tool.SendEmail(c,subject,body)
	//方法二

	err := tool.SendEmail(myToers, subject, body, serverHost, fromEmail, fromPasswd, serverPort)
	if err != nil {
		tool.Fail(c, "发送失败", err.Error())
	} else {
		emaildata := &Email{
			Code:  ty,
			Email: myToers,
			Time:  time.Now().Unix(),
			Uid:   uid,
		}

		if err := db.Db.Table("email_codes").Create(&emaildata).Error; err != nil {
			tool.Fail(c, "存储失败", 400)

		} else {
			res := make(map[string]interface{})
			res["uid"] = uid
			res["email"] = myToers

			tool.Success(c, "发送成功", res)
		}

	}
}
