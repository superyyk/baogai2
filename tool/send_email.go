package tool

import (
	"fmt"
	_ "github.com/go-gomail/gomail"
	"gopkg.in/gomail.v2"
	//"net"
	"net/smtp"
	"strings"
)

type EmailParam struct {
	ServerHost string  //邮箱地址如：599260040@qq.com
	ServerPort int    //腾讯企业邮箱 465
	FromEmail string
	FromPasswd string
	Toers string
	CCers string //抄送给 多个用（"，"）隔开
}

var serverHost,fromEmail,fromPasswd string
var serverPort int
var m *gomail.Message

func InitEmail(ep *EmailParam)  {
	toers:=[]string{}
	serverHost=ep.ServerHost
	serverPort=ep.ServerPort
	fromEmail=ep.FromEmail
	fromPasswd=ep.FromPasswd
	m=gomail.NewMessage()
	if len(ep.Toers)==0 {
		return
	}
	for _,tmp:=range strings.Split(ep.Toers,","){
		toers=append(toers,strings.TrimSpace(tmp))
	}
	m.SetHeader("To",toers...)
	//抄送列表
	if len(ep.CCers)!=0 {
		for _,tmp:=range strings.Split(ep.CCers,","){
			toers=append(toers,strings.TrimSpace(tmp))
		}
	}
	//发邮件
	//第三个参数为发件人别名，如 李大锤 也可以为空
	m.SetAddressHeader("From",fromEmail,"通知提示")
}

//SendEmail body支持html格式字符串
func SendEmail(mail_to,subject,body,server_Host,from_Email,from_Passwd string,server_Port int) error  {


	m:=gomail.NewMessage()
	//设置发送给多个
	mailArrTo:=strings.Split(mail_to,",")
	m.SetHeader("From", m.FormatAddress(from_Email, "Gaoding.pro")) // 添加别名
	m.SetHeader("To",mailArrTo...)
	//主题
	m.SetHeader("Subject",subject)
	//正文
	m.SetBody("text/html",body)
	//m.Attach("/home/Alex/lolcat.jpg") 附件

	d:=gomail.NewDialer(server_Host,server_Port,from_Email,from_Passwd)
	//发送
	err:=d.DialAndSend(m)
	if err!=nil {
		fmt.Print(err)
		//Failed(c,"发送失败",400)
	}
	return err

}

//方法二
func SendToMail(user,passwd,host,to,subject,body,mailtype string) error {
      hp:=strings.Split(host,":")
      auth:=smtp.PlainAuth("",user,passwd,hp[0])
      var content_type string
      if mailtype=="html"{
      	content_type="Content-Type:text/"+mailtype+";charset=UTF-8"
      	} else {
      	content_type="Content-Type:text/plain"+";charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err



}




