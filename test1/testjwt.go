package test1

import (
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	tel := c.Query("tel")

	token, _ := utils.GenerateToken(tel)

	tool.Success(c, "success", token)

}
func ParsLogin(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		tool.Fail(c, "token 无效", 400)
		return
	}
	tel, _ := utils.ParseToken(token)
	tool.Success(c, "success", tel)
}
