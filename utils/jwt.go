package utils

import (
	"fmt"
	"os"
	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetJWT(long int, uid, name string) string {

	l := 24
	if long != 0 {
		l = long
	}
	OneDayOfHours := 60 * 60 * l
	expiresTime := time.Now().Unix() + int64(OneDayOfHours)
	//OneDayOfHours用来设置过期时间，这里设置的时一天
	//可以使用int  60*60*24

	user := &model.User{
		UUID:     uid,
		Username: name,
	}
	claims := jwt.StandardClaims{
		Audience:  user.Username,     // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        string(user.UUID), // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "admin0001",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	//通过 StandardClaims 生成标准的载体，也就是上文提到的七个字段，其中 编号设定为 用户 id。

	var jwtSecret = []byte(config.JWTSecret)
	//jwtSecret 是我们设定的密钥
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过 HS256 算法生成 tokenClaims ,这就是我们的 HEADER 部分和 PAYLOAD。

	token, err := tokenClaims.SignedString(jwtSecret)
	//生成token

	if err != nil {
		fmt.Println("err:", err.Error())
		//tool.Failed(c,"err:",err.Error())
	}

	token = "Bearer " + token
	//将 token 和 Bearer 拼接在一起，同时中间用空格隔开。
	//生成 Bearer Token
	//res:=make(map[string]interface{})
	//res["res"]=200
	//res["token"]=token
	//tool.Success(c,"success",token)

	return token

}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userID string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1440 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// DecodeToken handles decoding a jwt token
func ParseToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
