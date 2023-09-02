package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func Md5(val string) string  {
	data := []byte(val)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
func Md5_2(s string) string  {
	m:=md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

//MD5加盐加密 比不加盐安全性高，但是也不足够安全
func Md5_salt(password string) string {
	const salt = "yyk*2012"//自定义加盐
	hash := md5.New()
	hash.Write([]byte(password+salt))//密码与盐自定义组合
	res:=hex.EncodeToString(hash.Sum(nil))
	fmt.Println(len(res))
	return res
}

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
//加盐加密
func Sha256_salt(password string) string {
	const salt = "2021/10/21"//自定义加盐
	hash := sha256.New()
	hash.Write([]byte(password+salt))//密码与盐自定义组合
	res:=hex.EncodeToString(hash.Sum(nil))
	fmt.Println(len(res))
	return res
}

//密码编码为哈希值
func HashPassword(password string) (string, error) {
	start:=time.Now()
	fmt.Println("开始时间：",start)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)//bcrypt.DefaultCost默认数值10，编码一次100ms以内，可增大数值，增加破解时间成本，例如设置为14，编码一次1s以上
	time:=time.Since(start)
	log.Printf("加密结束时间:%s",time)
	return string(bytes), err
}
//验证密码，例如实际业务中登录密码与数据库存储的哈希值比较，以此验证是否相等
func MatchPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
