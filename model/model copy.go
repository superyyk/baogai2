package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	ID   string `gorm:"column:fID"`
	Name string `gorm:"column:fUser_name"`
	Age  int    `gorm:"column:fAge"`
}

type ConnInfo struct {
	Username string
	Password string
	Host     string
	Db       string
	Port     int
}
type Users2 struct {
	ID   string `gorm:"column:fID"`
	Name string `gorm:"column:fUser_name"`
}

type User struct {
	Username string `json:"username"`
	UUID     string `json:"id"`
}

type Row struct {
	ID   string
	Name string
	Age  int
}

// 图片上传
type Img struct {
	Uid     string `gorm:"column:Uid"`
	Path    string `gorm:"column:Path"`
	UserUid string `gorm:"column:UserId"`
	Time    int    `gorm:"column:Time"`
}

// 获取post参数模型
type PostParams struct {
	User_uuid string `form:user_uuid`
	Uid       string `form:uid`
}

//日志结构体

type Log struct {
	Ip      string        `gorm:"column:Ip",json:"ip"`
	Method  string        `gorm:"column:Method",json:"method"`
	Path    string        `gorm:"column:Path",json:"path"`
	Time    int64         `gorm:"column:start",json:"start"`
	Haoshi  time.Duration `gorm:"column:Haoshi",json:"haoshi"`
	Headers string        `gorm:"column:headers",json:"headers"`
}
type Cars struct {
	Name   string `gorm:"column:name",json:"name"`
	Type   int    `gorm:"column:type",json:"type"`
	Uid    string `gorm:"column:uid",json:"uid"`
	Status int    `gorm:"column:status",json:"status"`
	Img    string `gorm:"column:img",json:"img"`
	Tip    string `gorm:"column:tip",json:"tip"`
}

type Shouru struct {
	Money    float64 `gorm:"column:money",json:"money"`
	Type     int     `gorm:"column:type",json:"type"`
	Uid      string  `gorm:"column:uid",json:"uid"`
	Status   int     `gorm:"column:status",json:"status"`
	From_id  string  `gorm:"column:from_id",json:"from_id"`
	Time     int64   `gorm:"column:time",json:"time"`
	UserUuid string  `gorm:"column:useruuid",json:"user_uuid"`
	//CreatedAt       time.Time
	//UpdatedAt       time.Time
}

type Order struct {
	gorm.Model
	Uid       string  `gorm:"column:uid";json:"uid"`
	Price     string  `gorm:"column:price";json:"price"`
	Bi_id     string  `gorm:"column:bi_id";json:"bi_id"`
	Shoushu   string  `gorm:"column:shoushu";json:"shoushu"`
	Ganggan   string  `gorm:"column:ganggan";json:"ganggan"`
	Zhisun    string  `gorm:"column:zhishunlv";json:"zhisun"`
	Total     float64 `gorm:"column:total";json:"total"`
	Time      int64   `gorm:"column:time";json:"time"`
	Status    int     `gorm:"column:status";json:"status"`
	Type      string  `gorm:"column:type";json:"type"`
	User_uuid string  `gorm:"column:useruuid";json:"user_uuid"`
	Now_price float64 `gorm:"column:now_price";json:"now_price"`
}

type Order_info struct {
	gorm.Model
	Uid       string  `gorm:"column:uid";json:"uid"`
	Price     string  `gorm:"column:price";json:"price"`
	Bi_id     string  `gorm:"column:bi_id";json:"bi_id"`
	Shoushu   string  `gorm:"column:shoushu";json:"shoushu"`
	Ganggan   string  `gorm:"column:ganggan";json:"ganggan"`
	Zhisun    string  `gorm:"column:zhishunlv";json:"zhisun"`
	Total     float64 `gorm:"column:total";json:"total"`
	Time      int64   `gorm:"column:time";json:"time"`
	Status    int     `gorm:"column:status";json:"status"`
	Type      string  `gorm:"column:type";json:"type"`
	User_uuid string  `gorm:"column:useruuid";json:"user_uuid"`
	Now_price float64 `gorm:"column:now_price";json:"now_price"`
	Mc        interface{}
	Bi        interface{}
}

type Maichu struct {
	Order_uid    string  `gorm:"column:order_uid",json:"order_uid"`
	User_uuid    string  `gorm:"column:user_uuid",json:"user_uuid"`
	Price        float64 `gorm:"column:price",json:"price"`
	Shoushu      int     `gorm:"column:shoushu", json:"shoushu"`
	Status       int     `gorm:"column:status",json:"status"`
	Time         int64   `gorm:"column:time",json:"time"`
	Yingkui      string  `gorm:"column:yingkui",`
	Shouxufei    string  `gorm:"column:shouxufei",`
	Yingkui_type int     `gorm:"column:yingkui_type", json:"yingkui_type"`
}

type News struct {
	ID      string `gorm:"column:fID",json:"id"`
	Title   string `gorm:"column:fTitle",json:"title"`
	Content string `gorm:"column:fContent",json:"content"`
	Img     string `gorm:"column:fImg",json:"img"`
	Time    string `gorm:"column:fTime",json:"time"`
	Status  int    `gorm:"column:fStatus",json:"status"`
}
type Bis struct {
	PairId        string  `gorm:"column:pairId",json:"pair_id"`
	Logo          string  `gorm:"column:logo",json:"logo"`
	CName         string  `gorm:"column:cName",json:"c_name"`
	FullName      string  `gorm:"column:fullName",json:"full_name"`
	Shizhi        float64 `gorm:"column:shizhi",json:"shizhi"`
	Price         float64 `gorm:"column:price",json:"price"`
	BuyPrice      int     `gorm:"column:buy_price",json:"buy_price"`
	ChangePercent float64 `gorm:"column:changePercent",json:"change_percent"`
	nowPrice      string  `gorm:"column:now_price",json:"now_price"`
	Zanbi         float64 `gorm:"column:zhanbi",json:"zanbi"`
	Rank          int     `gorm:"column:rk",json:"rank"`
}

type News_detail struct {
	Father_id string `gorm:"column:fFather_id",json:"father_id"`
	Title     string `gorm:"column:fTitle",json:"title"`
	Content   string `gorm:"column:fContent",json:"content"`
	Img       string `gorm:"column:fImg",json:"img"`
	Time      string `gorm:"column:fTime",json:"time"`
}

type MemcacheReq struct {
	Ip       string `json:"ip"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Memcache *Log   `json:"memcache"`
}

type Navtab struct {
	ID    int64  `gorm:"column:fID",json:"id"`
	Title string `gorm:"column:title",json:"title"`
	Uid   string `gorm:"column:uid",json:"uid"`
}

// 消息
type AllMsg struct {
	Ty     string      `json:"ty"`
	Userid string      `json:"user_id"`
	Group  string      `json:"group"`
	Data   interface{} `json:"data"`
	Info   interface{} `json:"info"`
}

// 一对一消息
type SongelMsg struct {
	Ty    string      `json:"ty"`
	Group string      `json:"group"`
	Uuid  string      `json:"uuid"`
	Data  interface{} `json:"data"`
}

type Kaihuhang struct {
	Name string `gorm:"column:khh",json:"name"`
}

type UserMaichu struct {
	gorm.Model
	Useruuid string  `gorm:"column:user_uuid",json:"useruuid"`
	Yingkui  float64 `gorm:"column:yingkui",json:"yingkui"`
}

//type RequestInfo struct {
//	Method string `json:"method"`
//	Path string `json:"path"`
//	Ip string `json:"ip"`
//	Headers interface{} `json:"headers"`
//	Start time.Time `json:"start"`
//	Haoshi time.Duration `json:"haoshi"`
//}

type DxmUsers struct {
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	Uid      string `gorm:"column:fUid",json:"uid"`
}
