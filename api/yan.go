package api

import (
	"encoding/json"
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func Chushou(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	num := tool.String_int(c.Query("num"))

	yan_id := c.Query("yan_id")
	yan_img := c.Query("yan_img")
	yan_name := c.Query("yan_name")
	price := c.Query("price")
	m_1 := c.Query("m_1")
	m_2 := c.Query("m_2")

	ty := tool.String_int(c.Query("ty"))

	//看看是否有同款烟入库
	var chushou []model.ChuShou
	Db.Table("chushou").Where("yan_id=? AND status=? AND user_uuid=?", yan_id, 0, user_uuid).Find(&chushou)
	if len(chushou) > 0 {
		//有，累加数量，添加入库信息

		n := tool.String_int(chushou[0].Num) + num
		row := Db.Table("chushou").Where("yan_id=? AND status=? AND user_uuid=?", yan_id, 0, user_uuid).Update("num", n)

		if row.RowsAffected == 1 {
			ruku := &model.RuKuInfo{
				YanId: yan_id,
				Name:  yan_name,
				Time:  time.Now().Unix(),
				Num:   num,
			}
			Db.Table("ruku_info").Create(&ruku)
			tool.Success(c, "入库成功", 200)
		} else {
			tool.Success(c, "系统繁忙，稍后再试", 401)
		}

	} else {

		//没有，则新增

		data := &model.ChuShou{
			User_uuid: user_uuid,
			Num:       tool.Int_string(num),
			//Num_big:    tool.Int_string(num_big),
			YanId:   tool.String_int(yan_id),
			YanName: yan_name,
			YanImg:  yan_img,
			Price:   price,
			//PriceBig:   price_big,
			Time:     time.Now().Unix(),
			Order_id: utils.GetOrderUuid(),
			Order_ty: ty,
			M_1:      tool.String_int(m_1),
			M_2:      tool.String_int(m_2),
		}
		if err := Db.Table("chushou").Create(&data).Error; err != nil {
			tool.Success(c, "系统繁忙，稍后再试", 401)
			return

		} else {
			ruku := &model.RuKuInfo{
				YanId: yan_id,
				Name:  yan_name,
				Time:  time.Now().Unix(),
				Num:   num,
			}
			Db.Table("ruku_info").Create(&ruku)
			tool.Success(c, "入库成功", 200)
		}
	}

}

func GetMaichu(c *gin.Context) {
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	user_uuid := c.Query("user_uuid")
	status := c.Query("status")

	s := tool.String_int(status)
	var maichu []model.ChuShou
	var baoguo []model.DaBao
	if s == 0 {
		offset := (pageNum - 1) * pagesize
		Db.Table("chushou").Where("user_uuid=? AND status=?", user_uuid, s).Offset(offset).Limit(pagesize).Order("id desc").Find(&maichu)
		for k, v := range maichu {
			maichu[k].Xiangyan = GetXiangyanById(v.YanId, v.Order_ty)
		}
		tool.Success(c, "success", maichu)
	} else {
		if s == 1 {
			offset := (pageNum - 1) * pagesize
			Db.Table("dabao").Where("user_uuid=? AND status=? OR status=?", user_uuid, s, 0).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
			tool.Success(c, "success", baoguo)
		} else {
			offset := (pageNum - 1) * pagesize
			Db.Table("dabao").Where("user_uuid=? AND status=?", user_uuid, status).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
			tool.Success(c, "success", baoguo)
		}

	}

}

func CancelMaichu(c *gin.Context) {
	id := c.Query("id")
	row := Db.Table("chushou").Where("id=?", id).Update("status", 4)
	if row.RowsAffected == 1 {
		tool.Success(c, "取消成功", 200)
		return
	} else {
		tool.Success(c, "取消失败！", 400)
	}
}

func Dabao(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	content := c.Query("content")
	orders_id := c.Query("orders_id")
	total_num := tool.String_int(c.Query("total_num"))
	total_money := c.Query("total_money")
	baoguo_id := utils.GetOrderUuid()
	t := time.Now().Unix()
	var dabao []model.DaBao
	Db.Table("dabao").Where("user_uuid=? AND status in (0,1)", user_uuid).Find(&dabao)
	if len(dabao) > 0 { //有未支付订单 不能打包
		tool.Success(c, "您有未支付包裹订单，打包失败", 401)
	} else {
		data := &model.DaBao{
			User_uuid:   user_uuid,
			Content:     content,
			Orders_id:   orders_id,
			Total_num:   total_num,
			Total_money: total_money,
			BaoguoId:    baoguo_id,
			Time:        t,
		}
		if err := Db.Table("dabao").Create(&data).Error; err != nil {
			//tx.Rollback()
			tool.Success(c, "系统繁忙，打包失败", 401)
			return
		} else {
			//打包成功
			ids := tool.String_slice(orders_id)
			var count int = 0
			var m = make([]string, len(ids))
			for _, v := range ids {

				row := Db.Table("chushou").Where("order_id=?", v).Update("status", 1)
				if row.RowsAffected == 1 {
					count += 1
					//tx.Commit()
				} else {
					//tx.Rollback()
					m = append(m, v)
				}
			}
			if count != len(ids) {
				tool.Success(c, "打包成功！", m)
				return
			}
			//tx.Commit()

			tool.Success(c, "打包成功！", baoguo_id)
		}
	}

	//tx := Db.Begin()

}

func GetBaoGuoById(c *gin.Context) {
	id := c.Query("id")
	user_uuid := c.Query("user_uuid")
	var dabao model.DaBao
	var yunfei_tui []model.YunfeiFanhuanInfo
	var list []model.ChuShou
	Db.Table("yunfei_fanhuan_info").Where("user_uuid=? AND ty=?", user_uuid, 2).Find(&yunfei_tui)
	Db.Table("dabao").Where("baoguo_id=?", id).First(&dabao)
	//获取运费余额
	yunfei_yue := GetYunfei_yue(user_uuid)

	res := make(map[string]interface{})
	if dabao.Pay_status == 0 {
		ids := dabao.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)

		for _, vv := range nums {
			chushou := GetChushouInfo(vv)
			list = append(list, chushou)
		}
		str, _ := json.Marshal(list)
		dabao.Content = string(str)
		res["dabao"] = dabao
		res["yunfei_tui"] = yunfei_tui
		res["yunfei_yue"] = yunfei_yue
		tool.Success(c, "sucess", res)
	} else {
		ids := dabao.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)
		var list []model.ChuShou
		for _, vv := range nums {
			var chu model.ChuShou
			Db.Table("chushou").Where("order_id=?", vv).First(&chu)
			list = append(list, chu)
		}
		str, _ := json.Marshal(list)
		dabao.Jiesuan_content = string(str)
		res["dabao"] = dabao
		res["yunfei_tui"] = yunfei_tui
		res["yunfei_yue"] = yunfei_yue
		tool.Success(c, "sucess", res)
	}

}

func GetYunfei_yue(user_uuid string) float64 {
	var tui_result []float64
	var kou_result []float64
	var yunfei_tui float64
	var yunfei_kou float64

	Db.Table("yunfei_fanhuan_info").Where("user_uuid=? AND ty=?", user_uuid, 2).Pluck("yunfei", &tui_result)
	Db.Table("yunfei_kouchu").Where("user_uuid=?", user_uuid).Pluck("price", &kou_result)
	for _, v := range tui_result {
		yunfei_tui += v
	}
	for _, v := range kou_result {
		yunfei_kou += v
	}
	return yunfei_tui - yunfei_kou

}
func GetChushouInfo(order_id string) (chushou model.ChuShou) {
	var c model.ChuShou
	//var xiangyan model.Xiangyan
	Db.Table("chushou").Where("order_id=?", order_id).First(&c)
	xiangyan := GetXiangyanById(c.YanId, c.Order_ty)
	c.Xiangyan = xiangyan
	return c
}

func GetXiangyanById(id, ty int) (yan model.Xiangyan) {
	//var info []model.Xiangyan
	Db.Table("xiangyan").Where("id=? AND ty=?", id, ty).Order("id desc").First(&yan)
	return
}

func GetBaoGuo(c *gin.Context) {
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	user_uuid := c.Query("user_uuid")
	status := c.Query("status")
	var baoguo []model.DaBao
	offset := (pageNum - 1) * pagesize
	s := tool.String_int(status)
	if s == 1 {
		Db.Table("dabao").Where("user_uuid=? AND status in (0,1)", user_uuid).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
	} else {
		Db.Table("dabao").Where("user_uuid=? AND status=?", user_uuid, s).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
	}

	tool.Success(c, "sucess", baoguo)

}

func GetAllDaBao(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var dabao []model.DaBao
	Db.Table("dabao").Where("user_uuid=?", user_uuid).Order("id desc").Find(&dabao)

	tool.Success(c, "success", dabao)
}

func GetBaoguoByWuliuStatus(c *gin.Context) {
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	user_uuid := c.Query("user_uuid")
	status := c.Query("status")

	s := tool.String_int(status)
	var baoguo []model.DaBao
	offset := (pageNum - 1) * pagesize
	var count int = 0
	Db.Table("dabao").Where("user_uuid=? AND wuliu_status=?", user_uuid, 2).Count(&count)
	Db.Table("dabao").Where("user_uuid=? AND wuliu_status=?", user_uuid, s).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
	res := make(map[string]interface{})
	res["jichu_count"] = count
	res["baoguo"] = baoguo
	tool.Success(c, "sucess", res)
}
func GetBaoguoAll(c *gin.Context) {
	pagesize := tool.String_int(c.Query("pagesize"))
	pageNum := tool.String_int(c.Query("pagenum"))
	user_uuid := c.Query("user_uuid")
	var baoguo []model.DaBao
	offset := (pageNum - 1) * pagesize
	Db.Table("dabao").Where("user_uuid=?", user_uuid).Offset(offset).Limit(pagesize).Order("id desc").Find(&baoguo)
	tool.Success(c, "sucess", baoguo)
}
