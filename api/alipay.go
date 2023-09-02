package api

import (
	"context"
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"github.com/superyyk/yishougai/api/cert"
	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/db"
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/jinzhu/gorm"

	alipay1 "github.com/smartwalle/alipay/v2"
	alipay3 "github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/ncrypto"
)

func GoAliPay(c *gin.Context) {
	// 初始化支付宝客户端
	//
	//	appid：应用ID
	//	privateKey：应用私钥，支持PKCS1和PKCS8
	//	isProd：是否是正式环境
	client, err := alipay.NewClient(config.Alipay.MyAppId, config.Alipay.MyPrivateKey, true)
	if err != nil {
		// xlog.Error(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	client.SetBodySize(10) // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//
	//	注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
		//SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
		SetNotifyUrl("http://www.feimiertech.com/fmr/public/index.php/index/index/alipay_notify") // 设置异步通知URL
		// SetAppAuthToken()                           // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	RSA2, _ := os.Open("alipayCertPublicKey_RSA2.crt")
	AlipayPublicContentRSA2, _ := ioutil.ReadAll(RSA2)
	client.AutoVerifySign(AlipayPublicContentRSA2)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("/appCertPublicKey_2021004108655660.crt", "/alipayRootCert.crt", "/alipayCertPublicKey_RSA2.crt")
	// 证书内容
	//err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
	bm := make(gopay.BodyMap)
	bm.Set("subject", "测试支付").
		Set("out_trade_no", utils.Uuid(23)).
		Set("total_amount", "0.01")
	ctx := context.Background()
	rep, _ := client.TradeAppPay(ctx, bm)
	res := make(map[string]interface{})
	res["rep"] = rep
	res["er"] = err
	tool.Success(c, "success", res)

}

func GoAliPay2(c *gin.Context) {

	price := c.Query("price")
	title := c.Query("title")
	ty := c.Query("ty")
	if ty == "0" {
		var privateKey = config.Alipay.MyPrivateKey // 这个地方填写用软件生成的应用私钥 RSA2 私钥
		var appId = config.Alipay.MyAppId
		// 创建支付宝客户端，沙箱环境下第三个参数填false，生产环境要换成true
		client, _ := alipay3.New(appId, privateKey, true)

		// 如果涉及到资金的支出，必须用公钥证书模式，如果只涉及支付推荐使用普通公钥模式，也就是刚刚生成的RSA2密钥
		//  公钥证书模式

		er := client.LoadAppPublicCertFromFile("appCertPublicKey_2021004108655660.crt") // 加载应用公钥证书
		er1 := client.LoadAliPayRootCertFromFile("alipayRootCert.crt")                  // 加载支付宝根证书
		er2 := client.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt")      // 加载支付宝公钥证书

		//普通公钥模式，加载支付宝公钥，在沙箱环境上传软件生成的应用公钥RSA2公钥后生成的支付宝公钥

		//client.LoadAliPayPublicKey()
		trade_no := utils.Uuid(30)
		var p = alipay3.TradeAppPay{}
		uid := utils.Uuid(20)
		p.NotifyURL = "http://www.feimiertech.com/fmr/public/index.php/index/index/alipay_notify" //回调地址
		p.Subject = "测试支付"
		//template := "2006-01-02 15:04:05"
		p.OutTradeNo = trade_no
		//p.OutTradeNo = uid     //自己生成一个唯一的交易单号
		p.TotalAmount = "0.01" //交易金额
		p.ProductCode = "p_1010101"

		//标准参数
		//p.NotifyURL = "http://203.86.24.181:3000/alipay"
		p.Body = "body"
		//p.Subject = "商品标题"
		//p.OutTradeNo = "01010101"
		//p.TotalAmount = "100.00"
		//p.ProductCode = "p_1010101"

		values := url.Values{} //回调参数，这里只能这样写，要进行urlEncode才能传给支付宝
		// 需要回传的参数
		values.Add("aaa", "aaa")
		values.Add("bbb", "bbb")
		p.PassbackParams = values.Encode() //支付宝会把passback_params={aaa=aaa&bbb=bbb}发送到回调函数

		// 这里返回的url中会包含sign，直接返回给前端就ok
		var url, err = client.TradeAppPay(p)
		if err != nil {
			fmt.Println(err)
		}
		res := make(map[string]interface{})
		res["url"] = url
		res["er"] = er
		res["er1"] = er1
		res["er2"] = er2
		res["price"] = price
		res["title"] = title
		res["order_id"] = uid
		res["trad_no"] = trade_no
		tool.Success(c, "success", res)
	}

}

func GoAlipayV2(c *gin.Context) {
	//支付宝公钥
	//var aliPublicKey = "xxx" // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
	//应用私钥
	//var privateKey = "xxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	//应用id
	//var appId = "xxx"
	price := c.Query("price")
	title := c.Query("title")
	ty := c.Query("ty")
	//最后一个参数 false 表示沙箱环境 true表示正式环境

	client, _ := alipay1.New(config.Alipay.MyAppId, config.Alipay.MyPub, config.Alipay.MyPrivateKey, true)
	//调用Page 支付(根据不同情况,选择不同支付)
	var p = alipay1.TradeAppPay{}
	//配置回调地址
	//p.NotifyURL = "http://xxx/alipayNotify" //post
	//配置支付完成后的跳转地址
	//p.ReturnURL = "http://xxx/alipayReturn"
	p.Subject = title
	template := "2006-01-02 15:04:05"
	//订单号
	p.OutTradeNo = time.Now().Format(template)
	p.TotalAmount = price //元
	//产品code
	//p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var resP, err4 = client.TradeAppPay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	res := make(map[string]interface{})
	res["resP"] = resP
	res["ty"] = ty
	tool.Success(c, "success", res)
}

// go-pay
func GetGoPay(c *gin.Context) {
	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	title := c.Query("title")
	price := c.Query("price")
	uid := utils.Uuid(20)
	client, err := alipay.NewClient(config.Alipay.MyAppId, config.Alipay.MyPrivateKey, true)
	res := make(map[string]interface{})
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
		// 设置字符编码，不设置默认 utf-8
		SetSignType(alipay.RSA2).                                             // 设置签名类型，不设置默认 RSA2
		SetReturnUrl("https://lzscxb.js.cool").                               // 设置返回URL，付款结束后跳转的url
		SetNotifyUrl("http://3n3swpx.nat.ipyingshe.com/v1/pay/alipay/notify") // 设置异步通知URL

		// 自动同步验签（只支持证书模式）
		// 传入 alipayCertPublicKey_RSA2.crt 内容
	// res["AppPublicContent"] = cert.AppPublicContent
	// res["AlipayRootContent"] = cert.AlipayRootContent
	// res["AlipayPublicContentRSA2"] = cert.AlipayPublicContentRSA2

	// tool.Success(c, "success", res)
	client.AutoVerifySign(cert.AlipayPublicContentRSA2)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	// 证书内容
	err = client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)

	if err != nil {
		panic(err)
	}

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", title). // 标题
					Set("out_trade_no", uid).                     // 订单号，支付成功后会返回
					Set("total_amount", price).                   // 订单金额
					Set("timeout_express", "2m").                 // 支付超时时间
					Set("product_code", "FAST_INSTANT_TRADE_PAY") // 必填 具体参考文档

	aliRsp, err := client.TradeAppPay(c, bm)
	if err != nil {
		//xlog.Error("err:", err)
		return
	}

	res["aliRsp"] = aliRsp
	res["AppPublicContent"] = cert.AppPublicContent
	res["AlipayRootContent"] = cert.AlipayRootContent
	res["AlipayPublicContentRSA2"] = cert.AlipayPublicContentRSA2

	tool.Success(c, "success", res)
	//fmt.Println(aliRsp)

}
func LoadAppPublicCertFromFile(filename string) (error, string) {
	return LoadAppCertPublicKeyFromFile(filename)
}

func LoadAppCertPublicKeyFromFile(filename string) (error, string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err, ""
	}

	return loadAppCertPublicKey(b)
}

func loadAppCertPublicKey(b []byte) (error, string) {
	cert, err := ncrypto.DecodeCertificate(b)
	if err != nil {
		return err, ""
	}
	appCertSN := getCertSN(cert)
	return nil, appCertSN
}
func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}

func CheckOrderStatus(c *gin.Context) {
	baoguo_id := c.Query("baoguo_id")
	var dabao []model.DaBao
	Db.Table("dabao").Where("baoguo_id=? AND pay_status==1", baoguo_id).Find(&dabao)
	if len(dabao) > 0 { //找到
		tool.Success(c, "success", 200)
		return
	} else {
		tool.Success(c, "success", 401)
		return
	}
}

//微信支付

func GoWeixinpay(c *gin.Context) {

}

// 修改包裹订单 支付状态
type PriceModel struct {
	Order_id string "json:order_id"
	Price    string "json:price"
}

func ChangeBaoguoPayStatus(c *gin.Context) {
	baoguo_id := c.Query("baoguo_id")
	//trade_no := c.Query("trade_no")
	yunfei := c.Query("yunfei")
	//yunfei_org := c.Query("yunfei_org")
	content := c.Query("content")
	price_list := c.Query("price_list")
	var pp []PriceModel
	err := json.Unmarshal([]byte(price_list), &pp)
	db.CheckErr(err)
	row := Db.Table("dabao").Where("baoguo_id=?", baoguo_id).Updates(map[string]interface{}{
		"pay_status": 1,

		"pay_time":     time.Now().Unix(),
		"wuliu_status": 1,
		"status":       2,

		"content": content,
	})
	if row.RowsAffected == 1 {
		var count int = 0
		for _, p := range pp {
			r := Db.Table("chushou").Where("order_id=?", p.Order_id).Update("price", p.Price)
			if r.RowsAffected == 1 {
				count += 1
			}

		}
		tool.Success(c, "支付成功", 200)

	} else {
		tool.Success(c, "支付失败", yunfei)
	}
}

func GetJiesuan(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var dabao []model.DaBao
	var tuikuan []model.DaBao
	var saoma []model.Saoma
	var pub model.Pub
	saoma = GetSaoma(user_uuid)
	Db.Table("dabao").Where("user_uuid=?", user_uuid).Find(&dabao)
	Db.Table("dabao").Where("user_uuid=? AND pay_status=?", user_uuid, 3).Find(&tuikuan)
	Db.Table("version").Where("id=?", 1).First(&pub)
	for k, v := range dabao {
		ids := v.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)

		var list []model.ChuShou
		for _, k := range nums {
			var chushou []model.ChuShou
			Db.Table("chushou").Where("order_id=?", k).Find(&chushou)
			list = append(list, chushou[0])
		}
		l, _ := json.Marshal(list)
		dabao[k].Content = string(l)

	}
	//获取直推vip已支付的钱，和人数
	m, count, mm, ccount := GetYijiVipMoneyAndCount(user_uuid)
	y := make(map[string]interface{})
	y["money"] = m
	y["count"] = count
	y["money_7"] = mm
	y["count_7"] = ccount
	//二级vip已支付的钱，和人数
	e := make(map[string]interface{})

	em, ec, eem, eec := GetErjiMoneyAndCount(user_uuid)
	e["money"] = em
	e["count"] = ec
	e["money_7"] = eem
	e["count_7"] = eec
	res := make(map[string]interface{})
	res["dabao"] = dabao
	res["tuikuan"] = tuikuan
	res["saoma"] = saoma
	res["pub"] = pub
	res["yiji"] = y
	res["erji"] = e

	tool.Success(c, "success", res)
}

func GetErjiMoneyAndCount(user_uuid string) (float64, int, float64, int) {
	var saoma []model.Saoma
	var m float64 = 0
	var count int = 0
	var mm float64 = 0
	var ccount int = 0
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for _, v := range saoma { //第一级
		var erji []model.Saoma
		Db.Table("saoma").Where("father_id=?", v.Child_id).Find(&erji) //二级
		for _, vv := range erji {
			//查vip订单
			var Vip_order model.VipOrders
			Db.Table("vip_orders").Where("user_uuid=? AND status=?", vv.Child_id, 1).First(&Vip_order)
			is_7 := Is7Day(Vip_order.PayTime)
			if is_7 { //大于7天的
				mm += tool.String_float64(Vip_order.Price)
				if Vip_order.Status == 1 {
					ccount += 1
				}

			} else { //小于7天的
				m += tool.String_float64(Vip_order.Price)
				if Vip_order.Status == 1 {
					count += 1
				}
			}

		}
	}
	return m, count, mm, ccount
}

func Is7Day(day int64) bool { //是否大于7天
	// 获取当前时间的时间戳
	//now := time.Now().Unix()
	// 计算七天前的时间戳
	sevenDaysAgo := time.Now().Add(-7 * 24 * time.Hour).Unix()
	if day < sevenDaysAgo {
		return true
	} else {
		return false
	}

}

func GetSaoma(user_uuid string) []model.Saoma {
	var saoma []model.Saoma
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for k, v := range saoma {
		var erji []model.Saoma
		Db.Table("saoma").Where("father_id=?", v.Child_id).Find(&erji)
		saoma[k].Erji = GetErji(v.Child_id)
		saoma[k].Info = GetUserDetail(v.Child_id)
	
	}
	return saoma
}

func CancelBaoGuo(c *gin.Context) {
	baoguo_id := c.Query("baoguo_id")
	ids := c.Query("ids")
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(ids, -1)
	//tx := Db.Begin()
	var count int = 0
	Db.Transaction(func(tx *gorm.DB) error {
		row := tx.Table("dabao").Where("baoguo_id=?", baoguo_id).Update("status", 6)
		if row.RowsAffected == 1 {

			for _, v := range nums {

				tx.Transaction(func(tx2 *gorm.DB) error {
					r := tx2.Table("chushou").Where("order_id=?", v).Update("status", 0)
					if r.RowsAffected == 1 {
						//tx.Commit()
						count += 1
					} else {
						tx2.Rollback()
						return nil
					}
					return nil
				})

			}

		} else {
			return nil
		}
		return nil
	})

	if count == len(nums) {
		tool.Success(c, "取消成功", 200)
	} else {
		tool.Success(c, "取消失败", 401)
	}

}
