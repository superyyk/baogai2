package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"poker/config"
	"poker/model"
	"poker/tool"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var KD100Flags = map[string]string{
	"ane66":          "安能快递",
	"debangwuliu":    "德邦物流",
	"debangkuaidi":   "德邦快递",
	"ems":            "EMS",
	"guotongkuaidi":  "国通快递",
	"huitongkuaidi":  "百世快递",
	"jd":             "京东物流",
	"kuayue":         "跨越速运",
	"pjbest":         "品骏快递",
	"shentong":       "申通快递",
	"shunfeng":       "顺丰速运",
	"suer":           "速尔快递",
	"xinfengwuliu":   "信丰物流",
	"youshuwuliu":    "优速物流",
	"youzhengguonei": "邮政快递包裹",
	"yuantong":       "圆通速递",
	"yuantongguoji":  "圆通国际",
	"yunda":          "韵达快递",
	"zhaijisong":     "宅急送",
	"zhongtong":      "中通快递",
	"ewe":            "EWE全球快递",
	"quanyikuaidi":   "全一快递",
	"tiantian":       "天天快递",
	"sxjdfreight":    "顺心捷达",
	"dhl":            "DHL-中国件",
	"tnt":            "TNT",
	"other":          "其它快递",
}

type DData struct {
	DefFirstPrice string `json:"defFirstPrice"`
	DefOverPrice  string `json:"defOverPrice"`
	DefPrice      string `json:"defPrice"`
	FirstPrice    string `json:"firstPrice"`
	OverPrice     string `json:"overPrice"`
	Price         string `json:"price"` //折后总价，单位：元
	ServiceType   string `json:"serviceType"`
}

type OrderRespData struct {
	TaskId     string `json:"taskId"`
	OrderId    string `json:"orderId"`
	Kuaidinum  string `json:"kuaidinum"`
	EOrder     string `json:"eOrder"`
	Attach     string `json:"attach"`
	KuaidiName string `json:"kuaidiCom"`
}

type resp struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
	Data       DData  `json:"data"`
}
type cresp struct {
	Result     bool        `json:"result"`
	ReturnCode string      `json:"returnCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type OrderResp struct {
	Result     bool          `json:"result"`
	ReturnCode string        `json:"returnCode"`
	Message    string        `json:"message"`
	Data       OrderRespData `json:"data"`
}
type FeiDetail struct {
	feeType   string `json:"feeType"`
	feeDesc   string `json:"feeDesc"`
	amount    string `json:"amount"`
	payStatus string `json:"payStatus"`
}

type SearchRespData struct {
	cargo           string      `json:"cargo"`
	comment         string      `json:"comment"`
	courierMobile   string      `json:"courierMobile"`
	courierName     string      `json:"courierName"`
	createTime      string      `json:"createTime"`
	dayType         string      `json:"dayType"`
	freight         string      `json:"freight"`
	feeDetails      []FeiDetail `json:"feeDetails"`
	kuaidiCom       string      `json:"kuaidiCom"`
	kuaidiNum       string      `json:"kuaidiNum"`
	lastWeight      string      `json:"lastWeight"`
	orderId         string      `json:"orderId"`
	payStatus       string      `json:"payStatus"`
	payment         string      `json:"payment"`
	pickupEndTime   string      `json:"pickupEndTime"`
	pickupStartTime string      `json:"pickupStartTime"`
	preWeight       string      `json:"preWeight"`
	defPrice        string      `json:"defPrice"`
	recAddr         string      `json:"recAddr"`
	recCity         string      `json:"recCity"`
	recDistrict     string      `json:"recDistrict"`
	recMobile       string      `json:"recMobile"`
	recName         string      `json:"recName"`
	recProvince     string      `json:"recProvince"`
	sendAddr        string      `json:"sendAddr"`
	sendCity        string      `json:"sendCity"`
	sendDistrict    string      `json:"sendDistrict"`
	sendMobile      string      `json:"sendMobile"`
	sendName        string      `json:"sendName"`
	sendProvince    string      `json:"sendProvince"`
	serviceType     string      `json:"serviceType"`
	status          string      `json:"status"`
	taskId          string      `json:"taskId"`
	valins          string      `json:"valins"`
}

type searchResp struct {
	Result     bool           `json:"result"`
	ReturnCode string         `json:"returnCode"`
	Message    string         `json:"message"`
	Data       SearchRespData `json:"data"`
}

func GetYunFei(c *gin.Context) {
	ad1 := c.Query("ad1")
	ad2 := c.Query("ad2")
	weight := c.Query("weight")
	kuaidi_code := c.Query("kuaidi_code")
	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/order/borderapi.do"
	paramData := make(map[string]string)
	paramData["kuaidiCom"] = kuaidi_code
	paramData["sendManPrintAddr"] = ad1
	paramData["recManPrintAddr"] = ad2
	paramData["weight"] = weight
	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000
	method := c.Query("method")
	// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
	sign := strings.ToUpper(GetMD5Encode(paramjson + tool.Int64_string(t) + key + secrect))
	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"method": {method}, "key": {key}, "sign": {sign}, "t": {tool.Int64_string(t)}, "param": {paramjson}})
	if postErr != nil {
		//fmt.Println("查询失败" + postErr.Error())
		tool.Success(c, "查询失败", 401)
		return
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		tool.Success(c, "查询失败,请至快递公司官网自行查询", 401)
		return
	}
	var resdata *resp
	if err := json.Unmarshal([]byte(postBody), &resdata); err != nil {
		//fmt.Println("Unmarshal resdata fail,", err)
		//httpext.Error(ctx, e.ERROR)
		tool.Success(c, "获取失败", 401)
		return
	}
	r := make(map[string]interface{})
	r["http_resp"] = resdata
	r["t"] = t
	r["paramjson"] = paramjson
	r["res"] = string(postBody)
	// r["sign"] = sign
	// r["key"] = key
	// r["secrect"] = secrect
	tool.Success(c, "success", r)

}
func JiJian(c *gin.Context) { //寄件
	user_uuid := c.Query("user_uuid")
	ad1 := c.Query("ad1")
	name1 := c.Query("name1")
	tel1 := c.Query("tel1")
	baoguo_id := c.Query("baoguo_id")
	ad2 := c.Query("ad2")
	name2 := c.Query("name2")
	tel2 := c.Query("tel2")
	weight := c.Query("weight")
	kuaidi_code := c.Query("kuaidi_code")
	kuai_price := tool.String_float64(c.Query("kuaidi_price"))
	yunfei_org := c.Query("yunfei_org")
	wupin_name := c.Query("wupin_name")
	yuyue_time := c.Query("yuyue_time")
	//yunfei_yue := tool.String_float64(c.Query("yunfei_yue"))
	method := c.Query("method")
	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/order/borderapi.do"
	paramData := make(map[string]string)
	paramData["kuaidicom"] = kuaidi_code
	paramData["recManName"] = name2
	paramData["recManMobile"] = tel2
	paramData["recManPrintAddr"] = ad2
	paramData["sendManName"] = name1
	paramData["sendManMobile"] = tel1
	paramData["sendManPrintAddr"] = ad1
	paramData["callBackUrl"] = config.Base.CallBackUrl
	paramData["cargo"] = wupin_name   //物品名称,例：文件。当kuaidicom=jd，yuantong时，必填
	paramData["dayType"] = yuyue_time //预约日期，例如：今天/明天/后天

	paramData["weight"] = weight
	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000
	var need_price float64 = 0  //须补交运费
	if config.Base.Kuaidi_Bug { //正式上线可下单

		yunfei_yue := GetYunfei_yue(user_uuid)

		if yunfei_yue >= kuai_price { //运费余额支付
			err := YunfeiKoufei(user_uuid, baoguo_id, kuai_price)
			if err != nil {
				tool.Success(c, "下单失败", 401)
			}
		} else { //运费余额不足,全部扣除
			if yunfei_yue >= 0 && yunfei_yue < kuai_price {
				err := YunfeiKoufei(user_uuid, baoguo_id, yunfei_yue)
				if err != nil {
					tool.Success(c, "下单失败", 401)
				}
				need_price = kuai_price - yunfei_yue
			}

		}

		//快递公司接口下单
		// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
		sign := strings.ToUpper(GetMD5Encode(paramjson + tool.Int64_string(t) + key + secrect))
		//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
		postRes, postErr := http.PostForm(posturl, url.Values{"method": {method}, "key": {key}, "sign": {sign}, "t": {tool.Int64_string(t)}, "param": {paramjson}})
		if postErr != nil {
			//fmt.Println("查询失败" + postErr.Error())
			tool.Success(c, "下单", 401)
			return
		}
		postBody, err := ioutil.ReadAll(postRes.Body)
		if err != nil {
			tool.Success(c, "下单请求失败", 401)
			return
		}
		var orderresp *OrderResp
		if err := json.Unmarshal([]byte(postBody), &orderresp); err != nil {
			//fmt.Println("Unmarshal resdata fail,", err)
			//httpext.Error(ctx, e.ERROR)
			tool.Success(c, "下单失败", 401)
			return
		}
		//更改包裹状态为下单成功
		r := make(map[string]interface{})
		if orderresp.ReturnCode == "200" {
			row := Db.Table("dabao").Where("baoguo_id=?", baoguo_id).Updates(map[string]interface{}{
				"wuliu_status": 1,
				"status":       1,
				"xiadan_time":  time.Now().Unix(),
				"kuaidi_code":  kuaidi_code,
				"yunfei":       need_price,
				"yunfei_org":   yunfei_org,
				"task_id":      orderresp.Data.TaskId,
				"order_id":     orderresp.Data.OrderId,
				"yundan_num":   orderresp.Data.Kuaidinum,
				"weight":       weight,
			})

			if row.RowsAffected == 1 {

				r["http_resp"] = orderresp
				r["t"] = t
				r["paramjson"] = paramjson
				r["res"] = string(postBody)
				tool.Success(c, "success", r)
			} else {

				tool.Success(c, "下单成功，订单更新失败", 401)
			}
		} else {
			r["http_resp"] = orderresp
			r["t"] = t
			r["paramjson"] = paramjson
			r["res"] = string(postBody)
			tool.Success(c, "下单失败", r)
		}

	} else { //模拟下单成功
		da := &OrderRespData{
			TaskId:  "12121212",
			OrderId: "232323",
		}
		cancelresp := &OrderResp{
			Result:     true,
			ReturnCode: "200",
			Message:    "下单成功",
			Data:       *da,
		}
		row := Db.Table("dabao").Where("baoguo_id=?", baoguo_id).Updates(map[string]interface{}{
			"wuliu_status": 1,
			"status":       1,
			"xiadan_time":  time.Now().Unix(),
			"kuaidi_code":  kuaidi_code,
			"yunfei":       kuai_price,
			"task_id":      "orderresp.Data.TaskId",
			"order_id":     "orderresp.Data.OrderId",
			"yundan_num":   "orderresp.Data.Kuaidinum",
		})
		r := make(map[string]interface{})
		if row.RowsAffected == 1 {
			r["http_resp"] = cancelresp
			r["paramjson"] = paramjson
			tool.Success(c, "success", r)
		} else {

			tool.Success(c, "下单失败", 401)
		}

	}
}

func YunfeiKoufei(user_uuid, baoguo_id string, price float64) error {
	yunfei_kouchu := &model.YunfeiKouchu{
		Time:     time.Now().Unix(),
		BaoguoId: baoguo_id,
		Price:    price,
		UserUuid: user_uuid,
	}
	if err := Db.Table("yunfei_kouchu").Create(&yunfei_kouchu).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func CancelOrder(c *gin.Context) { //取消物流下单
	task_id := c.Query("task_id")
	order_id := c.Query("order_id")
	cancel_message := c.Query("cancel_message")

	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/order/borderapi.do"
	paramData := make(map[string]string)
	paramData["taskId"] = task_id
	paramData["orderId"] = order_id
	paramData["cancelMsg"] = cancel_message
	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000
	method := c.Query("method")

	// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
	sign := strings.ToUpper(GetMD5Encode(paramjson + tool.Int64_string(t) + key + secrect))
	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"method": {method}, "key": {key}, "sign": {sign}, "t": {tool.Int64_string(t)}, "param": {paramjson}})
	if postErr != nil {
		//fmt.Println("查询失败" + postErr.Error())
		tool.Success(c, "取消失败", 401)
		return
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		tool.Success(c, "取消请求失败", 401)
		return
	}
	var cancelresp *cresp
	if err := json.Unmarshal([]byte(postBody), &cancelresp); err != nil {
		//fmt.Println("Unmarshal resdata fail,", err)
		//httpext.Error(ctx, e.ERROR)
		tool.Success(c, "取消失败", 401)
		return
	}
	r := make(map[string]interface{})
	r["http_resp"] = cancelresp
	r["t"] = t
	r["paramjson"] = paramjson
	r["res"] = string(postBody)
	tool.Success(c, "success", r)

}

func CancelWuliuOrder(c *gin.Context) {

	baoguo_id := c.Query("baoguo_id")
	ty := c.Query("ty")
	//取消物流订单
	var baoguo model.DaBao

	Db.Table("dabao").Where("baoguo_id=?", baoguo_id).First(&baoguo)
	err := CancelWuliu(baoguo.Task_id, baoguo.Order_id)
	if err != nil {
		tool.Success(c, "快递/物流订单取消失败", 401)
	} else {
		row := Db.Table("dabao").Where("baoguo_id=?", baoguo_id).Updates(map[string]interface{}{
			"status":       6,
			"pay_status":   3,
			"tuikuan_time": time.Now().Unix(),
		})
		if row.RowsAffected == 1 {
			yunfei_fanhuan := &model.YunfeiFanhuanInfo{
				Time:        time.Now().Unix(),
				User_uuid:   baoguo.User_uuid,
				YunFei:      baoguo.Yunfei,
				BaoguoId:    baoguo_id,
				Yundan_num:  baoguo.Yundan_num,
				Kuaidi_code: baoguo.Kuidi_code,
				Ty:          ty,
			}
			if err := Db.Table("yunfei_fanhuan_info").Create(&yunfei_fanhuan).Error; err != nil {
				tool.Success(c, "申请成功，返还运费失败", 401)
				return
			} else {
				tool.Success(c, "申请成功", 200)
			}

		} else {
			tool.Success(c, "系统繁忙，稍后再试", 401)
		}
	}

}

func CancelWuliu(task_id_in, order_id_in string) error {
	task_id := task_id_in
	order_id := order_id_in
	cancel_message := "暂时不寄了"

	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/order/borderapi.do"
	paramData := make(map[string]string)
	paramData["taskId"] = task_id
	paramData["orderId"] = order_id
	paramData["cancelMsg"] = cancel_message
	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000
	method := "cancel"

	// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
	sign := strings.ToUpper(GetMD5Encode(paramjson + tool.Int64_string(t) + key + secrect))
	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"method": {method}, "key": {key}, "sign": {sign}, "t": {tool.Int64_string(t)}, "param": {paramjson}})
	if postErr != nil {
		//fmt.Println("查询失败" + postErr.Error())

		return postErr
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {

		return err
	}
	var cancelresp *cresp
	if err := json.Unmarshal([]byte(postBody), &cancelresp); err != nil {
		//fmt.Println("Unmarshal resdata fail,", err)
		//httpext.Error(ctx, e.ERROR)

		return err
	}
	if cancelresp.ReturnCode == "200" {
		return nil
	} else {
		return err
	}

	// r := make(map[string]interface{})
	// r["http_resp"] = cancelresp
	// r["t"] = t
	// r["paramjson"] = paramjson
	// r["res"] = string(postBody)
	// tool.Success(c, "success", r)
}

func SearchOrder(c *gin.Context) {
	method := c.Query("method")
	task_id := c.Query("task_id")
	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/order/borderapi.do"
	paramData := make(map[string]string)
	paramData["taskId"] = task_id

	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000

	// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
	sign := strings.ToUpper(GetMD5Encode(paramjson + tool.Int64_string(t) + key + secrect))
	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"method": {method}, "key": {key}, "sign": {sign}, "t": {tool.Int64_string(t)}, "param": {paramjson}})
	if postErr != nil {
		//fmt.Println("查询失败" + postErr.Error())
		tool.Success(c, "查询失败", 401)
		return
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		tool.Success(c, "查询请求失败", 401)
		return
	}
	var searresp *searchResp
	if err := json.Unmarshal([]byte(postBody), &searresp); err != nil {
		//fmt.Println("Unmarshal resdata fail,", err)
		//httpext.Error(ctx, e.ERROR)
		tool.Success(c, "查询失败", 401)
		return
	}
	r := make(map[string]interface{})
	r["http_resp"] = searresp
	r["t"] = t
	r["paramjson"] = paramjson
	r["res"] = string(postBody)
	//r["sign"] = sign
	//r["key"] = key
	//r["secrect"] = secrect
	tool.Success(c, "success", r)

}

//实时快递

func ShishiKuaidi(c *gin.Context) {
	//method := c.Query("method")
	com := c.Query("com") //快递编码
	num := c.Query("num") //快递单号
	key := config.Base.Key100
	secrect := config.Base.Secret100
	posturl := "https://poll.kuaidi100.com/poll/query.do"
	paramData := make(map[string]string)
	paramData["com"] = com
	paramData["num"] = num

	paramDataSlice, _ := json.Marshal(paramData)
	paramjson := string(paramDataSlice)
	t := time.Now().Unix() * 1000

	// sign := strings.ToUpper(crypto.GetMd5String(paramjson + key + customer))
	sign := strings.ToUpper(GetMD5Encode(paramjson + key + secrect))
	//POST请求需要三个参数，分别为customer(CustomerId)和sign(签名)和param(参数)
	postRes, postErr := http.PostForm(posturl, url.Values{"customer": {secrect}, "sign": {sign}, "param": {paramjson}})
	if postErr != nil {
		//fmt.Println("查询失败" + postErr.Error())
		tool.Success(c, "查询失败", 401)
		return
	}
	postBody, err := ioutil.ReadAll(postRes.Body)
	if err != nil {
		tool.Success(c, "查询请求失败", 401)
		return
	}
	var searresp *searchResp
	if err := json.Unmarshal([]byte(postBody), &searresp); err != nil {
		//fmt.Println("Unmarshal resdata fail,", err)
		//httpext.Error(ctx, e.ERROR)
		tool.Success(c, "查询失败", 401)
		return
	}
	r := make(map[string]interface{})
	r["http_resp"] = searresp
	r["t"] = t
	r["paramjson"] = paramjson
	r["res"] = string(postBody)
	if searresp.ReturnCode == "200" {
		tool.Success(c, "success", r)
	} else {
		tool.Success(c, "查询失败", searresp)
	}

}

//变更物流状态

func ChangeWuliuStatus(c *gin.Context) {
	baoguo_id := c.Query("baoguo_id")

	status := c.Query("status")
	row := Db.Table("dabao").Where("baoguo_id=?", baoguo_id).Updates(map[string]interface{}{
		"status":       status,
		"wuliu_status": status,

		"fahuo_time": time.Now().Unix(),
	})
	if row.RowsAffected == 1 {
		tool.Success(c, "发货成功！", 200)

	} else {
		tool.Success(c, "系统繁忙，稍后再试！", 401)
	}

}

func GetMD5(str string) string {
	m := GetMD5Encode(str)
	m = strings.ToUpper(m)
	return m
}

// 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
