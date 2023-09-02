package model

type Cludes struct {
	ID       int    `gorm:"column:id";json:"id"`
	Uid      string `gorm:"column:uid";json:"uid"`
	Name     string `gorm:"column:name";json:"name"`
	Puber    string `gorm:"column:puber";json:"puber"`
	Head     string `gorm:"column:head";json:"head"`
	Tip      string `gorm:"column:tip";json:"tip"`
	Count    int    `gorm:"column:count";json:"count"`
	Shenqing int    `gorm:"column:shenqing";json:"shenqing"`
	//Total     float64 `gorm:"column:total";json:"total"`
	Time   int64 `gorm:"column:time";json:"time"`
	Status int   `gorm:"column:status";json:"status"`
	//Type      string  `gorm:"column:type";json:"type"`
	//User_uuid string  `gorm:"column:useruuid";json:"user_uuid"`
	//Now_price float64 `gorm:"column:now_price";json:"now_price"`
	Ok      interface{} `gorm:"column:ok";json:"ok"`
	Shen    interface{} `gorm:"column:shen";json:"shen"`
	Lat     string      `gorm:"column:lat";json:"lat"`
	Lgt     string      `gorm:"column:lgt";json:"lgt"`
	Address string      `gorm:"column:address";json:"address"`
}

type JoinCludes struct {
	ID int `gorm:"column:id";json:"id"`
	//Uid   string `gorm:"column:uid";json:"uid"`
	Clube string `gorm:"column:clube_id";json:"clube_id"`
	Ask   string `gorm:"column:asker_id";json:"asker_id"`
	Pass  string `gorm:"column:pass_time";json:"pass_time"`
	Tip   string `gorm:"column:tip";json:"tip"`
	//Count int    `gorm:"column:count";json:"count"`
	//Total     float64 `gorm:"column:total";json:"total"`
	Time   int64 `gorm:"column:time";json:"time"`
	Status int   `gorm:"column:status";json:"status"`
	//Type      string  `gorm:"column:type";json:"type"`
	//User_uuid string  `gorm:"column:useruuid";json:"user_uuid"`
	//Now_price float64 `gorm:"column:now_price";json:"now_price"`
	Info       interface{} `gorm:"column:info";json:"info"`
	CludesInfo interface{} `gorm:"column:clubeinfo";json:"clubeinfo"`
}
