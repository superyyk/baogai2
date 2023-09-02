package route

import (
	"log"
	"poker/admin"
	"poker/api"
	"poker/cors"
	"poker/middlewares"
	"poker/redis"
	"poker/test1"
	"poker/tool"
	"poker/utils"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {

	r := gin.Default()
	user := r.Group("/api/user")
	user1 := r.Group("v1/api/user")
	test := r.Group("test")
	ad := r.Group("v1/admin")
	ad1 := r.Group("v1/api/admin")
	ad.Use(cors.Cors()) //跨域
	ad1.Use(cors.Cors())
	{
		ad.GET("/ping", admin.Ping)
		ad.GET("/login", admin.Login)
	}

	ad1.Use(middlewares.AdminAuth()) //须鉴权
	{
		ad1.POST("/get_huishoufenlei", admin.GetHuishouFenlei)
	}
	user.Use(cors.Cors())
	{
		user.POST("/register", api.UserRegister)
		user.GET("/login", api.UserLogin)
		user.GET("/regist", api.Regist)
		user.GET("/login_sms", api.LoginSms)
		user.GET("/send_sms", utils.SendSms)
		user.GET("/send_email", utils.SendEmail)
		user.GET("/login_email", api.LoginEmail)
		user.GET("/change_pass", api.ChangePass)
		user.GET("/get_xiangyan", api.GetXiangyan)
		user.GET("/get_xiangyan_by_name", api.GetXiangyanByName)
		user.GET("/get_xiangyan_by_code", api.GetXiangyanByCode)
		user.GET("/update_yan_url", api.UpdateYanUrl)
		user.GET("/get_distence", api.GetDistence)
		user.GET("/app_update", api.APPUpdate)
		user.GET("load_pub", api.LoadPub)
		user.GET("check_code", api.CheckCode)
		user.GET("web_check_tel", api.CheckTel)
		user.GET("web_check_sms_code", api.WebCheckSmsCode)
		user.GET("web_check_email_code", api.WebCheckEmailCode)
		user.GET("check_code_num", api.CheckCodeNum) //检查验证码发送超3次数
		user.GET("get_xiangyan_by_saoma", api.GetXiangYangBySaoma)
		user.GET("get_gonggao", api.GetGongGao)
		user.GET("set_status_1", api.SetStatus_1)
		user.GET("rpcx_test", api.RpcxTest)
	}
	//创建一个新limiter.
	limiter, err := redis.NewLimiter()
	if err != nil {
		log.Fatal(err)
	}
	user1.Use(cors.Cors())
	user1.Use(middlewares.RateMiddleware(limiter, 5000, 1)) // 如果ip请求连接数在两秒内超过5次，返回429并抛出error
	user1.Use(middlewares.Auth())

	{
		user1.GET("/get_user_info", api.GetUserInfo)
		user1.GET("/create_clude", api.CreateClube)
		user1.GET("/get_clude", api.GetClude)
		user1.POST("/up_img", tool.Upload)
		user1.GET("/search_clude", api.SearchClude)
		user1.GET("/join_clube", api.JoinClude)
		user1.GET("/clube_pass", api.ClubePass)
		user1.POST("/up_yuyin", tool.UpYuyin)
		//业务
		user1.GET("/get_nav", api.GetNav)
		//宝盖
		user1.GET("/chushou", api.Chushou)
		user1.GET("/get_maichu", api.GetMaichu)
		user1.GET("/cancel_maichu", api.CancelMaichu)
		user1.GET("/dabao", api.Dabao)
		user1.GET("/get_baoguo_by_id", api.GetBaoGuoById)
		user1.GET("/get_baoguo", api.GetBaoGuo)
		user1.GET("/add_address", api.AddAddress)
		user1.GET("/use_address", api.UseAddress)
		user1.GET("/get_yunfei", api.GetYunFei)
		user1.GET("/jijian", api.JiJian)
		user1.GET("/jijian_cancel", api.CancelOrder)
		user1.GET("/jijian_search", api.SearchOrder)
		user1.GET("/change_wuliu_status", api.ChangeWuliuStatus)
		user1.GET("/get_baoguo_by_wuliu_status", api.GetBaoguoByWuliuStatus)
		user1.GET("/get_baoguo_all", api.GetBaoguoAll)
		user1.GET("/go_alipay", api.GoAliPay) //GetGoPay //GoAliPay //GoAlipayV2
		user1.GET("check_order_status", api.CheckOrderStatus)
		user1.GET("go_weixinpay", api.GoWeixinpay)
		user1.GET("change_baoguo_pay_status", api.ChangeBaoguoPayStatus)
		user1.GET("/get_jiesuan", api.GetJiesuan) //获取结算页面的数据
		user1.GET("cancel_baoguo", api.CancelBaoGuo)
		user1.GET("cancel_wuliu_order", api.CancelWuliuOrder)
		user1.GET("shishi_kuaidi", api.ShishiKuaidi)
		user1.GET("get_all_dabao", api.GetAllDaBao)
		user1.GET("get_tuiguang_pic", api.GetTuiguang)
		user1.GET("hebing_pic", api.HebingPic)
		user1.GET("create_vip_order", api.CreateVipOrder) //创建会员下单
		user1.GET("load_vip_orders", api.LoadVipOrders)
		user1.GET("change_vip_pay_status", api.ChangeVipPayStatus)
		user1.GET("get_tuiguang", api.GetTuiguang) //获取推广详细
		user1.GET("set_redis", api.SetRedis)
		user1.GET("get_redis", api.Get_redis)
		user1.GET("yanshi_redis", api.Yanshi)
		user1.GET("get_yunfei_tui_info", api.GetYunfeiTuiInfo) //获取运费退还信息
		user1.GET("get_yunfei_yue", api.GetYunFeiYue)
		user1.GET("change_pay_pass", api.ChangePayPass)
		user1.GET("user_tixian_yue", api.UserTixianYue)
		user1.GET("fan_kui", api.FanKui)
		user1.GET("change_phone", api.ChangePhone)
		user1.GET("check_tel", api.Checktel)
		user1.GET("setting_info", api.SettingInfo)
		user1.GET("test_saoma", api.TestSaoma)
		user1.GET("get_jiaocheng", api.GetJiaoCheng)
		user1.GET("get_sucai", api.GetSucai)
		user1.GET("get_my_sucai", api.GetMySucai)
		user1.GET("pub_sucai", api.PubSucai)
		user1.GET("set_sucai", api.SetSucai)
		user1.GET("sucai_pay", api.SucaiPay)
		user1.GET("my_buy_sucai", api.GetMyBuySucai)
		user1.GET("get_baoguo_list", api.GetBaoGuoList)
		user1.GET("test_mysql", api.TestMysql)

	}
	//测试链接
	{
		test.GET("/add_one", test1.MongoAddOne)
		test.GET("/edit_one", test1.EditOne)
		test.GET("/test_login", test1.Login)
		test.GET("/test_pars_login", test1.ParsLogin)
	}
	return r

}
