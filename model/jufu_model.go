package model

type JuFuUserInfo struct {
	UserName string `gorm:"column:fUser_name",json:"user_name"`
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	Uid string `gorm:"column:fUid",json:"uid"`
	Tel string `gorm:"column:fTel",json:"tel"`
	Pack_address string `gorm:"column:fPack_fil_address",json:"pack_address"`
	UserType int `gorm:"column:fUser_Type",json:"user_type"`
	UserHead string `gorm:"column:fUser_head",json:"user_head"`
	Time int64 `gorm:"column:fTime",json:"time"`
	Password string `gorm:"column:fPassword",json:"password"`
	Pay_password int `gorm:"column:fPassword_pay",json:"pay_password"`
}

type PubOrder struct {
	Type int `gorm:"column:type",json:"type"`
	Title string `gorm:"column:title",json:"title"`
	T_shu int `gorm:"column:t_shu",json:"t_shu"`
	Fil_num float64 `gorm:"column:fil_num",json:"fil_num"`
	Gas_num float64 `gorm:"column:gas_num",json:"gas_num"`
	Zhouqi int `gorm:"column:zhouqi",json:"zhouqi"`
	Tian int `gorm:"column:tian",json:"tian"`
	Total float64 `gorm:"column:total",json:"total"`
	Count float64 `gorm:"column:count",json:"count"`
	Time int64 `gorm:"column:time",json:"time"`
	Uid string `gorm:"column:Uid",json:"uid"`
	Limit_time int64 `gorm:"column:limit_time",json:"limit_time"`
	Price float64 `gorm:"columen:price",json:"price"`

}

type FilGold struct {
	UserUuid string `gorm:"column:user_uuid",json:"user_uuid"`
	Type int `gorm:"column:type",json:"type"`
	Time int64 `gorm:"column:time",json:"time"`
	Status int `gorm:"column:status",json:"status"`
	Fil_num float64 `gorm:"column:fil_num",json:"fil_num"`
}

type Gold struct {
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	Uid string `gorm:"column:fUuid",json:"uid"`
	PingZheng string `gorm:"column:fPingzheng_uid",json:"ping_zheng"`
	Num float64 `gorm:"column:fNum",json:"num"`
	Status int `gorm:"column:fStatus",json:"status"`
	Time int64 `gorm:"column:fTime",json:"time"`
	Type int `gorm:"column:fType",json:"type"`
	Recev_address string `gorm:"column:recev_address",json:"recev_address"`
	Pay_address string `gorm:"column:pay_address",json:"pay_address"`

}

type UserOrders struct {
	Status int `gorm:"column:fStatus",json:"status"`
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	OrderUid string `gorm:"column:fOrder_uuid"json:"order_uid"`
	Fil_num float64 `gorm:"column:fFil",json:"fil_num"`
    Pub_order_uid string `gorm:"column:pub_order_uid",json:"pub_order_uid"`
	Time int64 `gorm:"column:time",json:"time"`
	T_shu float64 `gorm:"column:fT_shu",json:"t_shu"`


}

type PubInfo struct {
	Shouxufei float64 `gorm:"column:bi_shouxufei",json:"shouxufei"`
	Bi_min float64 `gorm:"column:bi_min"`
	Fil_address string `gorm:"column:f_fil_address"`
	Version int `gorm:"column:version",json:"version"`
}

type Tibi struct {
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	Fil_num float64 `gorm:"column:fFil_num",json:"fil_num"`
	Time int64 `gorm:"column:fTime",json:"time"`
	Address string `gorm:"column:fPack_address",json:"address"`
	Status int `gorm:"column:fStatus",json:"status"`
	Uid string `gorm:"column:fUid",json:"uid"`
	Type int `gorm:"column:fType",json:"type"`
    Manager_time int64 `gorm:"column:manager_time",json:"manager_time"`
 }


