package utils

import (
	"fmt"
	"github.com/superyyk/yishougai/model"
	"github.com/superyyk/yishougai/tool"
	"net/http"
	"os"
	"path"
	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var DB = db.Db

// 单文件上传
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		dst := path.Join("./static", file.Filename)
		saveErr := c.SaveUploadedFile(file, dst)
		if saveErr == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": dst,
			})
		}
	}
}

// 多文件上传
func UploadFiles(c *gin.Context) {
	var urls []string
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		dst := path.Join("./static", file.Filename)
		urls = append(urls, dst)
		c.SaveUploadedFile(file, dst)
	}
	fmt.Println(urls)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": urls,
	})
}

// 根据时间存储
func Upload(c *gin.Context) {
	//1、获取上传的文件
	file, err1 := c.FormFile("file")
	user_uuid := c.Query("user_uuid")
	uid := RandNum(12)
	if err1 == nil {
		//2、获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.String(http.StatusBadRequest, "文件类型不合法")
			return
		}
		//3、创建图片保存目录,linux下需要设置权限（0666可读可写） static/upload/20200623
		day := time.Now().Format("20060102")
		dir := "./static/" + day
		if err := os.MkdirAll(dir, 0666); err != nil {
			c.String(http.StatusBadRequest, "MkdirAll失败")
			return
		}
		//4、生成文件名称 144325235235.png
		fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
		//5、上传文件 static/upload/20200623/144325235235.png
		saveDir := path.Join(dir, fileUnixName+extName)
		c.SaveUploadedFile(file, saveDir)
		serverDir := config.Base.BaseUrl + tool.Int_string(config.Base.Port) + "/" + saveDir
		img := &model.Images{
			Src:    serverDir,
			Userid: user_uuid,
			Uid:    uid,
			Time:   time.Now().Unix(),
		}
		time.Sleep(time.Second)
		if err := DB.Table("images").Create(&img).Error; err != nil {
			fmt.Println("save err")
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
