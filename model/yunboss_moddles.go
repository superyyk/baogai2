package model

import "github.com/jinzhu/gorm"

type UserData struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Tel    string `gorm:"column:tel" json:"tel"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Status int    `gorm:"column:status" json:"status"`
	Time   string `gorm:"time" json:"time"`
	Age    string `gorm:"column:age" json:"age"`
	Sex    string `gorm:"column:sex" json:"sex"`
	//Ty           int    `gorm:"column:ty" json:"ty"`
	//Level        int    `gorm:"column:level" json:"level"`
	Head  string `gorm:"column:head" json:"head"`
	Pass  string `gorm:"column:pass" json:"pass"`
	Email string `gorm:"column:email" json:"email"`
	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
	//Openid       string `gorm:"column:openid" json:"openid"`
}

type Codes struct {
	Tel    string `gorm:"column:tel" json:"tel"`
	Code   int    `gorm:"column:code" json:"code"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Time   int64  `gorm:"column:time" json:"time"`
	Email  string `gorm:"column:email" json:"email"`
	Status int    `gorm:"column:status" json:"status"`
}
type EmailCodes struct {
	//Tel    string `gorm:"column:tel" json:"tel"`
	Code  int    `gorm:"column:code" json:"code"`
	Uid   string `gorm:"column:uid" json:"uid"`
	Time  int64  `gorm:"column:time" json:"time"`
	Email string `gorm:"column:email" json:"email"`
	//Status int    `gorm:"column:status" json:"status"`
}

type ManageData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Img    string `json:"img"`
	Status int    `json:"status"`
}

type Manage struct {
	gorm.Model
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Uid    int    `gorm:"column:uid" json:"uid"`
	Status int    `gorm:"column:status" json:"status"`
	Owner  string `gorm:"column:owner" json:"owner"`
	Ind    int    `gorm:"column:ind" json:"ind"`
	Img    string `gorm:"column:img" json:"img"`
}

type Product struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Uid       string     `gorm:"column:uid" json:"uid"`
	Status    int        `gorm:"column:status" json:"status"`
	Puber     string     `gorm:"column:puber" json:"puber"`
	Imgs      string     `gorm:"column:imgs" json:"imgs"`
	Guige     string     `gorm:"column:guige" json:"guige"`
	Suply     int        `gorm:"column:suply" json:"suply"`
	Pertime   int        `gorm:"column:pertime" json:"pertime"`
	Perhaocai int        `gorm:"column:perhaocai" json:"perhaocai"`
	Sort      int        `gorm:"column:sort" json:"sort"`
	Time      int        `gorm:"column:time" json:"time"`
	Price     string     `gorm:"column:price" json:"price"`
	Tip       string     `gorm:"column:tip" json:"tip"`
	Unit      string     `gorm:"cloumn:unit" json:"unit"`
	Liucheng  []Liucheng `json:"liucheng"`
}

type Kehu struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Age    int    `gorm:"column:age" json:"age"`
	Status int    `gorm:"column:status" json:"status"`
	Puber  string `gorm:"column:puber" json:"puber"`

	Uid   string `gorm:"column:uid" json:"uid"`
	Level int    `gorm:"column:level" json:"level"`
	Tel   string `gorm:"column:tel" json:"tel"`
}

type Worker struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Status int    `gorm:"column:status" json:"status"`
	Puber  string `gorm:"column:puber" json:"puber"`
	//Imgs      interface{} `gorm:"column:imgs" json:"imgs"`
	Tel    string `gorm:"column:tel" json:"tel"`
	Ty     int    `gorm:"column:ty" json:"ty"`
	Salary string `gorm:"column:salary" json:"salary"`
	Age    int    `gorm:"column:age" json:"age"`
	//Sort      int         `gorm:"column:sort" json:"sort"`
	Time int `gorm:"column:time" json:"time"`
	//Price     string      `gorm:"column:price" json:"price"`
	//Tip       string      `gorm:"column:tip" json:"tip"`

}

type Liucheng struct { //生产流程
	Id        int    `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	ProductId string `gorm:"column:productid" json:"productid"`
	Maker     string `gorm:"column:maker" json:"maker"`
	Num       int    `gorm:"column:num" json:"num"`

	Maketime int64    `gorm:"column:maketime" json:"maketime"`
	Bad      int      `gorm:"column:bad" json:"bad"`
	Time     int64    `gorm:"column:time" json:"time"`
	Worker   []Worker `json:"worker"`
}

type Images struct { //图片
	Id     int    `gorm:"column:id" json:"id"`
	Src    string `gorm:"column:src" json:"src"`
	Userid string `gorm:"column:userid" json:"userid"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Time   int64  `gorm:"column:time" json:"time"`
}
type Yuyin struct { //图片
	Id     int    `gorm:"column:id" json:"id"`
	Src    string `gorm:"column:src" json:"src"`
	Userid string `gorm:"column:userid" json:"userid"`
	Uid    string `gorm:"column:uid" json:"uid"`
	Second string `gorm:"column:second" json:"second"`
	Time   int64  `gorm:"column:time" json:"time"`
}

type Unit struct { //计量单位
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Status int    `gorm:"column:status" json:"status"`
	Sort   int    `gorm:"column:sort" json:"sort"`
}

type Haocai struct { //耗材
	Id     int         `gorm:"column:id" json:"id"`
	Name   string      `gorm:"column:name" json:"name"`
	Images interface{} `gorm:"column:images" json:"images"`
	Status int         `gorm:"column:status" json:"status"`
	Puber  string      `gorm:"column:puber" json:"puber"`

	Uid   string `gorm:"column:uid" json:"uid"`
	Tip   string `gorm:"column:tip" json:"tip"`
	Total int    `gorm:"column:total" json:"total"`
	Meter string `gorm:"column:meter" json:"meter"`
	Time  int64  `gorm:"column:time" json:"time"`
}
