package model


type SignOrders struct {
	Useruuid string `column:"useruuid" json:"useruuid"`
	Apikey string `column:"apikey" json:"apikey"`
	Secretkey string `column:"secretkey" json:"secretkey"`
	Symbol string `column:"symbol" json:"symbol"`
	Qty string `column:"qty" json:"qty"`
	Price string `column:"price" json:"price"`
	Side string `column:"side" json:"side"`
	Positionside string `column:"positionside" json:"positionside"`
	Orderuid string `column:"orderuid" json:"orderuid"`
	Time int64 `column:"time" json:"time"`
	Status int `column:"status" json:"status"`
	Ordertype int `column:"ordertype" json:"ordertype"`
	Ty string `column:"ty" json:"ty"`
	Signid int `column:"signid" json:"signid"`
}

type NsqResMsg struct {
	Ty string `json:"ty"`
	Status int `json:"status"`
	Uid string `json:"uid"`
	Data SignOrders `json:"data"`
	MsgData BiPrice `json:"msg_data"`
}

type NsqMsg struct {
	Ty string `json:"ty"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

type BiPrice struct {
	Symbol string `column:"symbol" json:"symbol"`
	Price float64 `column:"price" json:"price"`
	Time int64 `column:"time" json:"time"`
}
