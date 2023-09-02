package tool

import (
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

func GenerrateQrCode(c *gin.Context, fileName, content string, w, h int) {
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		Fail(c, "error", err.Error())
	}
	qrCoder, _ := barcode.Scale(qrCode, w, h)
	//os.Open("./kkmc_upload/")
	time_str := time.Now().Format("2006-01-02") + "-" + Int64_string(time.Now().Unix()) + "-"
	file, _ := os.Create("./static/qrcode/" + time_str + fileName + ".png")

	defer file.Close()

	png.Encode(file, qrCoder)
	//base64
	ff, _ := ioutil.ReadFile("./static/qrcode/" + time_str + fileName + ".png")
	//buf:=make([]byte,13000)
	res := make(map[string]interface{})
	base64.StdEncoding.EncodeToString(ff)
	res["time"] = time.Now().Unix()

	res["href"] = "https://www.kissnet.cn:39700/static/qrcode/" + time_str + fileName + ".png"
	Success(c, "success", res)

	fmt.Print(time.Now().Format("2006-01-02"))
}

func GenerrateQrCode_2(fileName, content string, w, h int) (error, string) {
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return err, ""
	}
	qrCoder, _ := barcode.Scale(qrCode, w, h)
	//os.Open("./kkmc_upload/")
	time_str := time.Now().Format("2006-01-02") + "-" + Int64_string(time.Now().Unix()) + "-"
	file, _ := os.Create("./static/qrcode/" + time_str + fileName + ".png")

	defer file.Close()

	png.Encode(file, qrCoder)
	//base64
	ff, _ := ioutil.ReadFile("./static/qrcode/" + time_str + fileName + ".png")
	//buf:=make([]byte,13000)
	//res := make(map[string]interface{})
	base64.StdEncoding.EncodeToString(ff)
	//res["time"] = time.Now().Unix()

	path := "./static/qrcode/" + time_str + fileName + ".png"

	//fmt.Print(time.Now().Format("2006-01-02"))
	return nil, path
}
