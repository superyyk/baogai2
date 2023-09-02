package api

import (
	"poker/model"
	"poker/tool"

	"github.com/gin-gonic/gin"
)

var nav []model.Nav

func GetNav(c *gin.Context) {
	Db.Table("goods_typs").Where("status=?", 0).Find(&nav)
	if len(nav) > 0 {
		tool.Success(c, "success", nav)
		return
	}
	tool.Fail(c, "no nav", 400)
}