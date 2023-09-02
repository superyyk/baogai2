package model

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type Orders struct {
	UserUuid string `gorm:"column:fUser_uuid",json:"user_uuid"`
	Num float64 `gorm:"column:fNum",json:"num"`
	HuiLv float64 `gorm:"column:fHui_lv",json:"hui_lv"`
	Status int `gorm:"column:"fStatus",json:"status"`
	Tshu int `gorm:"column:fT_shu",json:"tshu"`
	Fanlv float64 `gorm:"column:fFan_lv",json:"fanlv"`
	Beishu float64 `gorm:"column:fBei_shu",json:"beishu"`
	Usdt float64 `gorm:"column:fUsdt",json:"usdt"`
}
