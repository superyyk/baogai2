package api

import (
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"

	"github.com/gin-gonic/gin"
)

func GetGongGao(c *gin.Context) {
	var gg []model.Gonggao
	Db.Table("gonggao").Where("status=?", 0).Find(&gg)
	tool.Success(c, "success", gg)
}

func SetStatus_1(c *gin.Context) {
	var count int = 0
	row := Db.Table("xiangyan").Where("id>?", 0).Update("status", 1)
	if row.RowsAffected == 1 {
		count += 1
	}
	tool.Success(c, "success", count)
}
