package tool

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/superyyk/baogai/config"
	"github.com/superyyk/baogai/db"
	"github.com/superyyk/baogai/model"

	"github.com/gin-gonic/gin"
	//imgtype "github.com/shamsher31/goimgtype"
)

var Db = db.Db
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var nums = []rune("1234567890")

// 根据时间存储
func Upload(c *gin.Context) {
	//1、获取上传的文件
	f, err1 := c.FormFile("file")
	user_uuid := c.Query("user_uuid")
	uid := c.Query("uid")

	if err1 == nil {

		fileDot := strings.Index(f.Filename, ".")
		fileType := f.Filename[fileDot:]
		// ok := isImgFormat(fileType)
		// if !ok {
		// 	Fail(c, "上传文件类型错误", 301)
		// 	return
		// }
		//3、创建图片保存目录,linux下需要设置权限（0666可读可写） static/upload/20200623
		day := time.Now().Format("20060102")
		dir := "./static/fankui/" + day
		if err := os.MkdirAll(dir, 0666); err != nil {
			//c.String(http.StatusBadRequest, "MkdirAll失败")
			Fail(c, "MkdirAll失败", 400)
			return
		}
		//4、生成文件名称 144325235235.png
		fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
		//5、上传文件 static/upload/20200623/144325235235.png
		saveDir := path.Join(dir, fileUnixName+fileType)
		c.SaveUploadedFile(f, saveDir)
		serverDir := config.Base.BaseUrl + Int_string(config.Base.Port) + "/" + saveDir
		img := &model.Images{
			Src:    serverDir,
			Userid: user_uuid,
			Uid:    uid,
			Time:   time.Now().Unix(),
		}
		//time.Sleep(time.Second)
		if err := Db.Table("images").Create(&img).Error; err != nil {
			fmt.Println("save err")
			c.JSON(400, gin.H{
				"code":      400,
				"msg":       "上传失败",
				"data":      "",
				"serverDir": "",
				"uid":       "",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":      0,
			"msg":       "上传成功",
			"data":      saveDir,
			"serverDir": serverDir,
			"uid":       uid,
		})

	}
}

func UpYuyin(c *gin.Context) {
	//1、获取上传的文件
	f, _ := c.FormFile("file")
	user_uuid := c.Query("user_uuid")
	second := c.Query("second")
	uid := RandNum(12)
	fileDot := strings.Index(f.Filename, ".")
	fileType := f.Filename[fileDot:]
	day := time.Now().Format("20060102")
	dir := "./static/yuyin/" + day
	if err := os.MkdirAll(dir, 0666); err != nil {
		//c.String(http.StatusBadRequest, "MkdirAll失败")
		Fail(c, "MkdirAll失败", 400)
		return
	}
	//4、生成文件名称 144325235235.png
	fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
	//5、上传文件 static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+fileType)
	c.SaveUploadedFile(f, saveDir)
	serverDir := config.Base.BaseUrl + Int_string(config.Base.Port) + "/" + saveDir

	yuyin := &model.Yuyin{
		Src:    serverDir,
		Userid: user_uuid,
		Uid:    uid,
		Second: second,
		Time:   time.Now().Unix(),
	}
	//time.Sleep(time.Second)
	if err := Db.Table("yuyin").Create(&yuyin).Error; err != nil {
		fmt.Println("save err")
		c.JSON(400, gin.H{
			"code":      400,
			"msg":       "上传失败",
			"data":      "",
			"serverDir": "",
			"uid":       "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"msg":       "上传成功",
		"data":      saveDir,
		"serverDir": serverDir,
		"uid":       uid,
		"second":    second,
	})

}

func RandNum(n int) string { //大小写字母随机
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	return string(b)
}

// CheckImageExt：检查图片后缀
func CheckImageExt(fileName string) (bool, string) {
	ext := GetExt(fileName)
	for _, allowExt := range []string{".jpg", ".jpeg", ".png", ".gif", ".JPG", ".JPEG", ".PNG"} {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true, ext
		}
	}

	return false, ""
}

// CheckImageSize：检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		log.Println(err)
		//logging.Warn(err)
		return false
	}

	return size <= 5
}

// GetSize：获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 是否是文件格式
func isImgFormat(filetype string) bool {
	format := `^(\.(gif|png|jpg|jpeg|webp|svg|psd|bmp|tif|jfif|PNG|JPG|JPEG))$`
	ok, _ := regexp.MatchString(format, filetype)
	return ok
}
