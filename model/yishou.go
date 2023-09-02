package model

type Nav struct {
	Id     string `gorm:"column:id",json:"id"`
	Name   string `gorm:"column:name",json:"name"`
	Status string `gorm:"column:status",json:"status"`
}

type UserInfos struct {
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
	//Weixinopenid string `gorm:"column:weixinopenid" json:"weixinopenid"`
	//Openid       string `gorm:"column:openid" json:"openid"`
}
