package model

type UsersInfo struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Tel    string `gorm:"column:tel" json:"tel"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Status int    `gorm:"column:status" json:"status"`
	Time   string `gorm:"time" json:"time"`
	Age    string `gorm:"column:age" json:"age"`
	Sex    string `gorm:"column:sex" json:"sex"`
	Ty     int    `gorm:"column:ty" json:"ty"`
	//Level        int    `gorm:"column:level" json:"level"`
	Head     string `gorm:"column:head" json:"head"`
	Pass     string `gorm:"column:pass" json:"pass"`
	Email    string `gorm:"column:email" json:"email"`
	IsOnline string `gorm:"column:is_online" json:"is_online"`
	Vip      int    `gorm:"column:vip" json:"vip"`
	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
	//Openid       string `gorm:"column:openid" json:"openid"`
	Code         string      `gorm:"column:code" json:"code"`
	Vip_order    interface{} `gorm:"column:vip_order"  json:"vip_order"`
	Dabao_order  interface{} `gorm:"column:dabao_order"  json:"dabao_order"`
	PayPass      string      `gorm:"column:pay_pass"  json:"pay_pass"`
	QQ           string      `gorm:"column:qq"  json:"qq"`
	Weixin       string      `gorm:"column:weixin"  json:"weixin"`
	WeixinCode   string      `gorm:"column:weixin_code"  json:"weixin_code"`
	ShoukuanCode string      `gorm:"column:shoukuan_code"  json:"shoukuan_code"`
	Tixian_ty    string      `gorm:"column:tixian_ty"  json:"tixian_ty"`
	//Model        interface{} `json:",omitempty"`
	//Saoma []Saoma `gorm:"FOREIGNKEY:Father_id; ASSOCIATION_FOREIGNKEY:uid"`
}

type Users struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Tel    string `gorm:"column:tel" json:"tel"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Status int    `gorm:"column:status" json:"status"`
	Time   string `gorm:"time" json:"time"`
	Age    string `gorm:"column:age" json:"age"`
	Sex    string `gorm:"column:sex" json:"sex"`
	Ty     int    `gorm:"column:ty" json:"ty"`
	//Level        int    `gorm:"column:level" json:"level"`
	Head     string `gorm:"column:head" json:"head"`
	Pass     string `gorm:"column:pass" json:"pass"`
	Email    string `gorm:"column:email" json:"email"`
	IsOnline string `gorm:"column:is_online" json:"is_online"`
	Vip      int    `gorm:"column:vip" json:"vip"`
	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
	//Openid       string `gorm:"column:openid" json:"openid"`
	Code         string      `gorm:"column:code" json:"code"`
	Vip_order    interface{} `gorm:"column:vip_order"  json:"vip_order"`
	Dabao_order  interface{} `gorm:"column:dabao_order"  json:"dabao_order"`
	PayPass      string      `gorm:"column:pay_pass"  json:"pay_pass"`
	QQ           string      `gorm:"column:qq"  json:"qq"`
	Weixin       string      `gorm:"column:weixin"  json:"weixin"`
	WeixinCode   string      `gorm:"column:weixin_code"  json:"weixin_code"`
	ShoukuanCode string      `gorm:"column:shoukuan_code"  json:"shoukuan_code"`
	//Model        interface{} `json:",omitempty"`
	Saoma []Saoma `gorm:"FOREIGNKEY:Father_id; ASSOCIATION_FOREIGNKEY:Uid"`
}

type Xiangyan struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Img  string `gorm:"column:img" json:"img"`
	Code string `gorm:"column:code" json:"code"`
	//Code_big   string `gorm:"column:code_big" json:"code_big"`
	Price string `gorm:"price" json:"price"`
	//Price_big    string `gorm:"column:price_big" json:"price_big"`
	Status int `gorm:"column:status" json:"status"`
	//All_status   int    `gorm:"column:all_status" json:"all_status"`
	//Status_small int    `gorm:"column:status_small" json:"status_small"`
	Status_vip int `gorm:"column:status_vip" json:"status_vip"`
	//Max_small    int    `gorm:"column:max_small" json:"max_small"`
	//Max_big      int    `gorm:"column:max_big" json:"max_big"`
	Sort      int     `gorm:"column:sort" json:"sort"`
	Logo      string  `gorm:"column:logo" json:"logo"`
	Ty        int     `gorm:"column:ty" json:"ty"`
	ChushouId int     `json:"chushouId"`
	Chushou   Chushou `json:"chushou"`
	//IsOnline string `gorm:"column:is_online" json:"is_online"`
	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
	//Openid       string `gorm:"column:openid" json:"openid"`
}

// type Xiangyan_big struct {
// 	Id           int    `gorm:"column:id" json:"id"`
// 	Name         string `gorm:"column:name" json:"name"`
// 	Img          string `gorm:"column:img" json:"img"`
// 	Code_small   string `gorm:"column:code_small" json:"code_small"`
// 	Code_big     string `gorm:"column:code_big" json:"code_big"`
// 	Price_small  string `gorm:"price_small" json:"price_small"`
// 	Price_big    string `gorm:"column:price_big" json:"price_big"`
// 	Status       int    `gorm:"column:status" json:"status"`
// 	All_status   int    `gorm:"column:all_status" json:"all_status"`
// 	Status_small int    `gorm:"column:status_small" json:"status_small"`
// 	Status_big   int    `gorm:"column:status_big" json:"status_big"`
// 	Max_small    int    `gorm:"column:max_small" json:"max_small"`
// 	Max_big      int    `gorm:"column:max_big" json:"max_big"`
// 	Sort         int    `gorm:"column:sort" json:"sort"`
// 	Ty           int    `gorm:"column:ty" json:"ty"`
// 	//IsOnline string `gorm:"column:is_online" json:"is_online"`
// 	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
// 	//Openid       string `gorm:"column:openid" json:"openid"`
// }

type ChuShou struct {
	Id      int    `gorm:"column:id" json:"id"`
	YanName string `gorm:"column:yan_name" json:"yan_name"`
	Status  string `gorm:"column:status" json:"status"`
	Num     string `gorm:"column:num" json:"num"`
	//Num_big     string      `gorm:"column:num_big" json:"num_big"`
	Time      int64  `gorm:"column:time" json:"time"`
	User_uuid string `gorm:"column:user_uuid" json:"user_uuid"`
	YanId     int    `gorm:"column:yan_id" json:"yan_id"`
	YanImg    string `gorm:"column:yan_img" json:"yan_img"`
	OkNum     int    `gorm:"column:ok_num" json:"ok_num"`
	//OkBig       int         `gorm:"column:ok_big" json:"ok_big"`
	//NoSmall     int         `gorm:"column:no_small" json:"no_small"`
	//NoBig       int         `gorm:"column:no_big" json:"no_big"`
	FahuoTime   int64  `gorm:"column:fahuo_time" json:"fahuo_time"`
	ShouHuoTime int64  `gorm:"column:shouhuo_time" json:"shouhuo_time"`
	YanshouTime int64  `gorm:"column:yanshou_time" json:"yanshou_time"`
	Price       string `gorm:"column:price" json:"price"`
	//PriceBig    string      `gorm:"column:price_big" json:"price_big"`
	YundanCode string      `gorm:"column:yundan_code" json:"yundan_code"`
	BaoguoId   string      `gorm:"column:baoguo_id" json:"baoguo_id"`
	Order_id   string      `gorm:"column:order_id" json:"order_id"`
	Order_ty   int         `gorm:"column:order_ty" json:"order_ty"`
	Xiangyan   interface{} `gorm:"column:xiangyan" json:"xiangyan"`
	M_1        int         `gorm:"column:m_1" json:"m_1"`
	M_2        int         `gorm:"column:m_2" json:"m_2"`
}

type Chushou struct {
	Id      int    `gorm:"column:id" json:"id"`
	YanName string `gorm:"column:yan_name" json:"yan_name"`
	Status  string `gorm:"column:status" json:"status"`
	Num     string `gorm:"column:num" json:"num"`
	//Num_big     string      `gorm:"column:num_big" json:"num_big"`
	Time      int64  `gorm:"column:time" json:"time"`
	User_uuid string `gorm:"column:user_uuid" json:"user_uuid"`
	YanId     int    `gorm:"column:yan_id" json:"yan_id"`
	YanImg    string `gorm:"column:yan_img" json:"yan_img"`
	OkNum     int    `gorm:"column:ok_num" json:"ok_num"`
	//OkBig       int         `gorm:"column:ok_big" json:"ok_big"`
	//NoSmall     int         `gorm:"column:no_small" json:"no_small"`
	//NoBig       int         `gorm:"column:no_big" json:"no_big"`
	FahuoTime   int64  `gorm:"column:fahuo_time" json:"fahuo_time"`
	ShouHuoTime int64  `gorm:"column:shouhuo_time" json:"shouhuo_time"`
	YanshouTime int64  `gorm:"column:yanshou_time" json:"yanshou_time"`
	Price       string `gorm:"column:price" json:"price"`
	//PriceBig    string      `gorm:"column:price_big" json:"price_big"`
	YundanCode string      `gorm:"column:yundan_code" json:"yundan_code"`
	BaoguoId   string      `gorm:"column:baoguo_id" json:"baoguo_id"`
	Order_id   string      `gorm:"column:order_id" json:"order_id"`
	Order_ty   int         `gorm:"column:order_ty" json:"order_ty"`
	Xiangyan   interface{} `gorm:"column:xiangyan" json:"xiangyan"`
	M_1        int         `gorm:"column:m_1" json:"m_1"`
	M_2        int         `gorm:"column:m_2" json:"m_2"`
}

type DaBao struct {
	Id              int    `gorm:"column:id" json:"id"`
	Content         string `gorm:"column:content" json:"content"`                 //打包时候的包内容
	Jiesuan_content string `gorm:"column:jiesuan_content" json:"jiesuan_content"` //结算时候的包内容
	Status          string `gorm:"column:status" json:"status"`
	Orders_id       string `gorm:"column:orders_id" json:"orders_id"`
	//Fahuo     string `gorm:"column:num_big" json:"num_big"`
	Time      int64  `gorm:"column:time" json:"time"`
	User_uuid string `gorm:"column:user_uuid" json:"user_uuid"`
	//YanId       int    `gorm:"column:yan_id" json:"yan_id"`
	//YanImg      string `gorm:"column:yan_img" json:"yan_img"`
	//OkSmall     int    `gorm:"column:ok_small" json:"ok_small"`
	//OkBig       int    `gorm:"column:ok_big" json:"ok_big"`
	//NoSmall     int    `gorm:"column:no_small" json:"no_small"`
	//NoBig       int    `gorm:"column:no_big" json:"no_big"`
	FahuoTime   int64  `gorm:"column:fahuo_time" json:"fahuo_time"`
	Total_num   int    `gorm:"column:total_num" json:"total_num"`
	Total_money string `gorm:"column:total_money" json:"total_money"`
	Total_hege  int    `gorm:"column:total_hege" json:"total_hege"`
	Money       string `gorm:"column:money" json:"money"`
	ShouHuoTime int64  `gorm:"column:shouhuo_time" json:"shouhuo_time"`
	//YanshouTime int64  `gorm:"column:yanshou_time" json:"yanshou_time"`
	//PriceSmall  string `gorm:"column:price_small" json:"price_small"`
	//PriceBig    string `gorm:"column:price_big" json:"price_big"`
	Weight   int    `gorm:"column:weight" json:"weight"`
	BaoguoId string `gorm:"column:baoguo_id" json:"baoguo_id"`
	//Order_id    string `gorm:"column:" json:""`
	Wuliu_status      int     `gorm:"column:wuliu_status" json:"wuliu_status"`
	Yunfei            float64 `gorm:"column:yunfei" json:"yunfei"`
	Yunfei_org        float64 `gorm:"column:yunfei_org" json:"yunfei_org"`
	Xiadan_time       int64   `gorm:"column:xiadan_time" json:"xiadan_time"`
	Task_id           string  `gorm:"column:task_id" json:"task_id"`
	Order_id          string  `gorm:"column:order_id" json:"order_id"`
	Yundan_num        string  `gorm:"column:yundan_num" json:"yundan_num"`
	Kuidi_code        string  `gorm:"column:kuaidi_code" json:"kuaidi_code"`
	Small_buhege      int     `gorm:"column:small_buhege" json:"small_buhege"`
	Big_buhege        int     `gorm:"column:big_buhege" json:"big_buhege"`
	Yunfei_chajia     float64 `gorm:"column:yunfei_chajia" json:"yunfei_chajia"`
	Pay_status        int     `gorm:"column:pay_status" json:"pay_status"`
	Pay_time          int64   `gorm:"column:pay_time" json:"pay_time"`
	Pay_chajia        int     `gorm:"column:pay_chajia" json:"pay_chajia"`
	Pay_chajia_time   int64   `gorm:"column:pay_chajia_time" json:"pay_chajia_time"`
	Pay_chajia_status int     `gorm:"column:pay_chajia_status" json:"pay_chajia_status"`
	Trade_no          string  `gorm:"column:trade_no" json:"trade_no"`
	Tuikuan_time      int64   `gorm:"column:tuikuan_time" json:"tuikuan_time"`
	Tuikuan_dao_time  int64   `gorm:"column:tuikuan_dao_time" json:"tuikuan_dao_time"`
}

type Address struct {
	Id      int    `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Tel     string `gorm:"column:tel" json:"tel"`
	Address string `gorm:"column:address" json:"address"`

	Tip       string  `gorm:"column:tip" json:"tip"`
	Ty        int     `gorm:"column:ty" json:"ty"`
	User_uuid string  `gorm:"column:user_uuid" json:"user_uuid"`
	Status    int     `gorm:"column:status" json:"status"`
	Used      int     `gorm:"column:used" json:"used"`
	Lat       float64 `gorm:"column:lat" json:"lat"`
	Lgt       float64 `gorm:"column:lgt" json:"lgt"`
}

type WuliuComp struct {
	Id     int    `gorm:"column:id" json:"id"`
	Logo   string `gorm:"column:logo" json:"logo"`
	Name   string `gorm:"column:name" json:"name"`
	Code   string `gorm:"column:code" json:"code"`
	Price  string `gorm:"column:price" json:"price"`
	JiaJia string `gorm:"column:jiajia" json:"jiajia"`
	Status int    `gorm:"column:status" json:"status"`
	Sort   int    `gorm:"column:sort" json:"sort"`
}

type Vip struct {
	Id       int    `gorm:"column:id" json:"id"`
	Logo     string `gorm:"column:logo" json:"logo"`
	Name     string `gorm:"column:name" json:"name"`
	Code     string `gorm:"column:code" json:"code"`
	Price    string `gorm:"column:price" json:"price"`
	Quanxian string `gorm:"column:quanxian" json:"quanxian"`
	Status   int    `gorm:"column:status" json:"status"`
	Tip      string `gorm:"column:tip" json:"tip"`
	TimeLong int    `gorm:"column:time_long" json:"time_long"`
}

type Saoma struct {
	Id        int         `gorm:"column:id" json:"id"`
	Father    string      `gorm:"column:father" json:"father"`
	Father_id string      `gorm:"column:father_id" json:"father_id"`
	FatherId  string      `gorm:"column:father_id" json:"father_id"`
	Child     string      `gorm:"column:child" json:"child"`
	Child_id  string      `gorm:"column:child_id" json:"child_id"`
	ChildId   string      `gorm:"column:child_id" json:"child_id"`
	Time      int64       `gorm:"column:time" json:"time"`
	Status    int         `gorm:"column:status" json:"status"`
	Erji      interface{} `gorm:"column:erji" json:"erji"`
	Info      interface{} `gorm:"column:info" json:"info"`
	Is_show   int         `gorm:"column:is_show" json:"is_show"`
	Users     Users       `gorm:"FOREIGNKEY:Uid; ASSOCIATION_FOREIGNKEY:FatherId"`
}
type Version struct {
	Id        int    `gorm:"column:id" json:"id"`
	Status    int    `gorm:"column:status" json:"status"`
	Version   int    `gorm:"column:version" json:"version"`
	Tanchuang string `gorm:"column:tanchuang" json:"tanchuang"`
	Is_code   int    `gorm:"column:is_code" json:"is_code"`
	T_1       int    `gorm:"column:t_1" json:"t_1"`
	T_2       int    `gorm:"column:t_2" json:"t_2"`
	M_1       int    `gorm:"column:m_1" json:"m_1"`
	M_2       int    `gorm:"column:m_2" json:"m_2"`
}

type RuKuInfo struct {
	Id    int    `gorm:"column:id" json:"id"`
	YanId string `gorm:"column:yan_id" json:"yan_id"`
	Time  int64  `gorm:"column:time" json:"time"`
	Name  string `gorm:"column:name" json:"name"`
	Num   int    `gorm:"column:num" json:"num"`
}

type YunfeiFanhuanInfo struct {
	Id          int         `gorm:"column:id" json:"id"`
	Status      int         `gorm:"column:status" json:"status"`
	Time        int64       `gorm:"column:time" json:"time"`
	User_uuid   string      `gorm:"column:user_uuid" json:"user_uuid"`
	YunFei      float64     `gorm:"column:yunfei" json:"yunfei"`
	BaoguoId    string      `gorm:"column:baoguo_id" json:"baoguo_id"`
	Yundan_num  string      `gorm:"column:yundan_num" json:"yundan_num"`
	Kuaidi_code string      `gorm:"column:kuaidi_code" json:"kuaidi_code"`
	Ty          string      `gorm:"column:ty" json:"ty"`
	Baoguo_info interface{} `gorm:"column:baoguo_info" json:"baoguo_info"`
}

type PriceZoushi struct {
	Id    int    `gorm:"column:id" json:"id"`
	Time  int64  `gorm:"column:time" json:"time"`
	Price string `gorm:"column:price" json:"price"`
	YanId int    `gorm:"column:yan_id" json:"yan_id"`
}
type Pub struct {
	Id        int    `gorm:"column:id" json:"id"`
	Version   int    `gorm:"column:version" json:"version"`
	Status    int    `gorm:"column:status" json:"status"`
	Tanchuang string `gorm:"column:tanchuang" json:"tanchuang"`
	Is_code   int    `gorm:"column:is_code" json:"is_code"`
	T_1       int    `gorm:"column:t_1" json:"t_1"`
	T_2       int    `gorm:"column:t_2" json:"t_2"`
	M_1       int    `gorm:"column:m_1" json:"m_1"`
	M_2       int    `gorm:"column:m_2" json:"m_2"`
}

type TuiguangUrls struct {
	Id       int    `gorm:"column:id" json:"id"`
	Code     string `gorm:"column:code" json:"code"`
	Status   int    `gorm:"column:status" json:"status"`
	Url      string `gorm:"column:url" json:"url"`
	UserUuid string `gorm:"column:user_uuid" json:"user_uuid"`
}

type VipOrders struct {
	Id        int    `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Status    int    `gorm:"column:status" json:"status"`
	Vip_id    int    `gorm:"column:vip_id" json:"vip_id"`
	Price     string `gorm:"column:price" json:"price"`
	Time_long int64  `gorm:"column:time_long" json:"time_long"`
	Logo      string `gorm:"column:logo" json:"logo"`
	Time      int64  `gorm:"column:time" json:"time"`
	PayTime   int64  `gorm:"column:pay_time" json:"pay_time"`
	DaoqiTime int64  `gorm:"column:daoqi_time" json:"daoqi_time"`
	User_uuid string `gorm:"column:user_uuid" json:"user_uuid"`
	Trade_no  string `gorm:"column:trade_no" json:"trade_no"`
	T_1       int    `gorm:"column:t_1" json:"t_1"`
	T_2       int    `gorm:"column:t_2" json:"t_2"`
}

type YunfeiKouchu struct {
	Id       int     `gorm:"column:id" json:"id"`
	Time     int64   `gorm:"column:time" json:"time"`
	BaoguoId string  `gorm:"column:baoguo_id" json:"baoguo_id"`
	Price    float64 `gorm:"column:price" json:"price"`
	UserUuid string  `gorm:"column:user_uuid" json:"user_uuid"`
}

type Tixian struct {
	Id        int     `gorm:"column:id" json:"id"`
	Time      int64   `gorm:"column:time" json:"time"`
	Ty        int     `gorm:"column:ty" json:"ty"`
	Status    int     `gorm:"column:status" json:"status"`
	PassTime  int64   `gorm:"column:pass_time" json:"pass_time"`
	Price     float64 `gorm:"column:price" json:"price"`
	User_uuid string  `gorm:"column:user_uuid" json:"user_uuid"`
}

type Gonggao struct {
	Id      int    `gorm:"column:id" json:"id"`
	Time    int64  `gorm:"column:time" json:"time"`
	Title   string `gorm:"column:title" json:"title"`
	Status  int    `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
	Puber   string `gorm:"column:puber" json:"puber"`
}

type FanKui struct {
	Id        int    `gorm:"column:id" json:"id"`
	Time      int64  `gorm:"column:time" json:"time"`
	Tel       string `gorm:"column:tel" json:"tel"`
	Status    int    `gorm:"column:status" json:"status"`
	Content   string `gorm:"column:content" json:"content"`
	User_uuid string `gorm:"column:user_uuid" json:"user_uuid"`
	Uid       string `gorm:"column:uid" json:"uid"`
}
type JiaoCheng struct {
	Id      int    `gorm:"column:id" json:"id"`
	Time    int64  `gorm:"column:time" json:"time"`
	Title   string `gorm:"column:title" json:"title"`
	Status  int    `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
	Puber   string `gorm:"column:puber" json:"puber"`
}
type Sucai struct {
	Id        int     `gorm:"column:id" json:"id"`
	Content   string  `gorm:"column:content" json:"content"`
	Time      int64   `gorm:"column:time" json:"time"`
	Status    int     `gorm:"column:status" json:"status"`
	User_uuid string  `gorm:"column:user_uuid" json:"user_uuid"`
	Price     float64 `gorm:"column:price" json:"price"`
	Ty        int     `gorm:"column:ty" json:"ty"`
	Urls      string  `gorm:"column:urls" json:"urls"`
	Is_buy    int     `gorm:"column:is_buy" json:"is_buy"`
}

type BySucai struct {
	Id        int         `gorm:"column:id" json:"id"`
	Time      int64       `gorm:"column:time" json:"time"`
	Status    int         `gorm:"column:status" json:"status"`
	User_uuid string      `gorm:"column:user_uuid" json:"user_uuid"`
	Money     float64     `gorm:"column:money" json:"money"`
	Sucai_id  int         `gorm:"column:sucai_id" json:"sucai_id"`
	Info      interface{} `gorm:"column:info" json:"info"`
}
