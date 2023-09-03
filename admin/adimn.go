package admin

import (
	"github.com/superyyk/baogai/config"
	"github.com/superyyk/baogai/tool"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.JWTSecret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

func GetHuishouFenlei(c *gin.Context) {
	auth := c.Request.Header.Get("Token")

	if len(auth) == 0 {
		c.Abort()
		//tool.Failed(c,"have no authorization!",401)
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "没有权限!",
			"data": 400,
		})
	}

	auth = strings.Fields(auth)[1]
	//校验token
	claims, err := parseToken(auth)
	if err != nil {
		tool.Fail(c, "token有误", 400)
	}
	tool.Success(c, "success", claims.Id)

}
