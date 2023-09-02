package api

import (
	"github.com/superyyk/baogai/model"
	"github.com/superyyk/baogai/tool"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateVipOrder(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	price := c.Query("price")
	id := c.Query("id")
	time_long := c.Query("time_long")
	logo := c.Query("logo")
	t_1 := c.Query("t_1")
	t_2 := c.Query("t_2")
	t := time.Now().Unix()
	//判断是否有未支付的会员订单
	var v []model.VipOrders
	Db.Table("vip_orders").Where("user_uuid=? AND status in (0,1)", user_uuid).Find(&v)
	if len(v) > 0 { //有未支付或正在使用的会员，不能下单
		tool.Success(c, "您有未支付或已开通会员，当前账户无法下单！", 401)
		return
	} else {
		vip := &model.VipOrders{
			User_uuid: user_uuid,
			Price:     price,
			Vip_id:    tool.String_int(id),
			Time_long: tool.String_int64(time_long),
			Logo:      logo,
			Time:      t,
			Name:      c.Query("name"),
			T_1:       tool.String_int(t_1),
			T_2:       tool.String_int(t_2),
		}
		if err := Db.Table("vip_orders").Create(&vip).Error; err != nil {
			tool.Success(c, "下单失败，稍后再试", 401)
		} else {
			tool.Success(c, "下单成功", 200)
		}
	}

}

func LoadVipOrders(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var vip_orders []model.VipOrders
	Db.Table("vip_orders").Where("user_uuid=?", user_uuid).Find(&vip_orders)
	tool.Success(c, "success", vip_orders)
}

func ChangeVipPayStatus(c *gin.Context) {
	order_id := c.Query("order_id")
	trade_no := c.Query("trade_no")
	user_uuid := c.Query("user_uuid")
	ty := c.Query("ty")
	t := time.Now().Unix()
	tx := Db.Begin()

	row := tx.Table("vip_orders").Where("id=?", order_id).Updates(map[string]interface{}{
		"status":   1,
		"trade_no": trade_no,
		"pay_time": t,
	})
	row1 := tx.Table("users").Where("uid=?", user_uuid).Update("vip", ty)

	if row.RowsAffected == 1 && row1.RowsAffected == 1 {
		tx.Commit()
		tool.Success(c, "开通成功！", 200)
		return
	} else {
		tx.Rollback()
		tool.Success(c, "开通失败！", 401)
	}
}

func GetYijiVipMoneyAndCount(user_uuid string) (float64, int, float64, int) {
	var saoma []model.Saoma
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	var m float64 = 0
	var count int = 0
	var mm float64 = 0
	var ccount int = 0
	for _, v := range saoma {
		//查vip订单
		var Vip_order model.VipOrders
		Db.Table("vip_orders").Where("user_uuid=? AND status=?", v.Child_id, 1).First(&Vip_order)
		is_7 := Is7Day(Vip_order.PayTime)
		if is_7 { //大于7天
			mm += tool.String_float64(Vip_order.Price)
			if Vip_order.Status == 1 {
				ccount += 1
			}
		} else { //小于7天的
			m += tool.String_float64(Vip_order.Price)
			if Vip_order.Status == 1 {
				count += 1
			}
		}

	}
	return m, count, mm, ccount
}
