package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/redis"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := c.Request.Header.Get("token")

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
			c.Abort()
			//tool.Failed(c,"token已过期"+err.Error(),402)
			c.JSON(400, gin.H{
				"code": 0,
				"msg":  "token已过期/无效",
				"data": 400,
			})
		} else {
			fmt.Println("token 正确", claims)
		}
		c.Next()
	}

}

func AdminAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
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
			c.Abort()
			//tool.Failed(c,"token已过期"+err.Error(),402)
			c.JSON(400, gin.H{
				"code": 0,
				"msg":  "token已过期/无效",
				"data": 400,
			})
		} else {
			fmt.Println("token 正确", claims)
		}
		c.Next()
	}

}

// 限制访问
func RateMiddleware(limiter *redis.Limiter, num int64, t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果ip请求连接数在1秒内超过100次，返回429并抛出error
		if !limiter.Allow(c.ClientIP(), num, t*time.Second) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			//log.Println("too many requests")
			return
		}
		c.Next()
	}
}
