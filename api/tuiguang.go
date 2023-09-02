package api

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"poker/config"
	"poker/model"
	"poker/tool"
	"poker/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"github.com/golang/freetype"
)

func GetTuiguang(c *gin.Context) {
	user_uuid := c.Query("user_uuid")
	var saoma []model.Saoma
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&saoma)
	for k, v := range saoma {
		saoma[k].Info = GetUserDetail(v.Child_id)
		saoma[k].Erji = GetErji(v.Child_id)
	}
	tool.Success(c, "success", saoma)
}

func GetUserDetail(user_uuid string) model.UsersInfo {
	var info model.UsersInfo
	Db.Table("users").Where("uid=?", user_uuid).First(&info)
	orders := GetOrders(user_uuid)
	info.Vip_order = orders
	info.Dabao_order = GetDabaoOrders(user_uuid)
	return info
}

func GetDabaoOrders(user_uuid string) []model.DaBao {
	var dabao []model.DaBao
	Db.Table("dabao").Where("user_uuid=? AND status!=?", user_uuid, 6).Find(&dabao)
	for k, v := range dabao {
		ids := v.Orders_id
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(ids, -1)
		var list []model.ChuShou
		for _, vv := range nums {
			var chushou model.ChuShou
			Db.Table("chushou").Where("order_id=?", vv).First(&chushou)
			list = append(list, chushou)
		}
		l, _ := json.Marshal(list)
		dabao[k].Jiesuan_content = string(l)

	}
	return dabao
}

func GetErji(user_uuid string) []model.Saoma {
	var erji []model.Saoma
	Db.Table("saoma").Where("father_id=?", user_uuid).Find(&erji)
	for k, v := range erji {
		info := GetUserDetail(v.Child_id)
		erji[k].Info = info
		//erji[k].Erji = GetErji(v.Child_id)
	}
	return erji
}

func GetOrders(user_uuid string) []model.VipOrders {
	var orders []model.VipOrders
	Db.Table("vip_orders").Where("user_uuid=?", user_uuid).Find(&orders)
	return orders
}

func MergeImg(img map[int]int) string {
	var centerImg4 image.Image

	//背景图
	backgroundImgFile, _ := os.Open("./static/1.png")
	backgroundImg, _ := jpeg.Decode(backgroundImgFile)
	defer backgroundImgFile.Close()
	backgroundBound := backgroundImg.Bounds()
	//x轴坐标总数
	//backgroundX := backgroundBound.Size().X
	//y轴坐标总数
	//backgroundY := backgroundBound.Size().Y

	//添加图，第一张图片，如果是windows 换成c:/1.jpg
	// if pid1, ok := img[1]; ok {
	// 	centerImgFile1, _ := os.Open(fmt.Sprintf("./static/merge/%d-m.jpg", pid1))
	// 	centerImg1, _ = jpeg.Decode(centerImgFile1)
	// 	defer centerImgFile1.Close()
	// } else {
	// 	centerImgFile1, _ := os.Open("./static/image/white.jpg")
	// 	centerImg1, _ = jpeg.Decode(centerImgFile1)
	// 	defer centerImgFile1.Close()
	// }
	//centerBound := centerImg.Bounds()
	//x轴坐标总数
	//centerX := centerBound.Size().X
	//y轴坐标总数
	//centerY := centerBound.Size().Y

	//坐标偏差，x轴y轴 计算
	//newImgX := (backgroundX - centerX) / 2
	//newImgY := (backgroundY - centerY) / 2

	//第二张图片
	// if pid2, ok := img[2]; ok {
	// 	centerImgFile2, _ := os.Open(fmt.Sprintf("./static/merge/%d-m.jpg", pid2))
	// 	centerImg2, _ = jpeg.Decode(centerImgFile2)
	// 	defer centerImgFile2.Close()
	// } else {
	// 	centerImgFile2, _ := os.Open("./static/image/white.jpg")
	// 	centerImg2, _ = jpeg.Decode(centerImgFile2)
	// 	defer centerImgFile2.Close()
	// }

	//第三张图片
	// if pid3, ok := img[3]; ok {
	// 	centerImgFile3, _ := os.Open(fmt.Sprintf("./static/merge/%d-m.jpg", pid3))
	// 	centerImg3, _ = jpeg.Decode(centerImgFile3)
	// 	defer centerImgFile3.Close()
	// } else {
	// 	centerImgFile3, _ := os.Open("./static/image/white.jpg")
	// 	centerImg3, _ = jpeg.Decode(centerImgFile3)
	// 	defer centerImgFile3.Close()
	// }

	//第四张图片
	if pid4, ok := img[4]; ok {
		centerImgFile4, _ := os.Open(fmt.Sprintf("./static/merge/%d-m.png", pid4))
		centerImg4, _ = jpeg.Decode(centerImgFile4)
		defer centerImgFile4.Close()
	} else {
		centerImgFile4, _ := os.Open("./static/image/white.jpg")
		centerImg4, _ = jpeg.Decode(centerImgFile4)
		defer centerImgFile4.Close()
	}
	//offset1 := image.Pt(177, 172)
	//offset2 := image.Pt(1655, 172)
	//offset3 := image.Pt(177, 2108)
	offset4 := image.Pt(1655, 2108)
	//x轴坐标总数
	m := image.NewRGBA(backgroundBound)
	draw.Draw(m, backgroundBound, backgroundImg, image.Point{}, draw.Src)
	//draw.Draw(m, centerImg1.Bounds().Add(offset1), centerImg1, image.Point{}, draw.Over)
	//draw.Draw(m, centerImg2.Bounds().Add(offset2), centerImg2, image.Point{}, draw.Over)
	//draw.Draw(m, centerImg3.Bounds().Add(offset3), centerImg3, image.Point{}, draw.Over)
	draw.Draw(m, centerImg4.Bounds().Add(offset4), centerImg4, image.Point{}, draw.Over)
	url := "./static/print/1.jpg"
	imgw, _ := os.Create(url)
	err := jpeg.Encode(imgw, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Error("jpeg解码失败", err)

	}
	defer imgw.Close()
	return url
}

func HebingPic(c *gin.Context) {
	code := c.Query("code")
	user_uuid := c.Query("user_uuid")
	//判断有无图
	var tg []model.TuiguangUrls
	res := make(map[string]interface{})
	Db.Table("tuiguang_urls").Where("user_uuid=?", user_uuid).Find(&tg)
	if len(tg) > 0 { //找到推广图
		res["ty"] = 0
		res["data"] = tg
		tool.Success(c, "success", res)
		return
	} else {

		conent := config.Base.TpUrl + "?code=" + code
		name := code
		w := 400
		h := 400
		_, path := tool.GenerrateQrCode_2(name, conent, w, h) //获取二维码
		//生成文字图片
		text_img := TextToImage(name)
		var cc []string
		bb := []string{"./static/1.jpg", "./static/2.jpg", "./static/3.jpg", "./static/4.jpg", "./static/5.jpg", "./static/6.jpg"}
		for _, v := range bb {
			str := createImg(v, path, code, text_img)

			tg := &model.TuiguangUrls{
				Code:     code,
				UserUuid: user_uuid,
				Url:      str,
			}
			if err := Db.Table("tuiguang_urls").Create(&tg).Error; err != nil {
				fmt.Println(err)
			} else {
				cc = append(cc, str)
				res["ty"] = 1
				res["urls"] = cc

			}
		}

		//tool.Success(c, "success", str)
		tool.Success(c, "success", res)
	}

}

// import (
//     "bytes"
//     "image"
//     "image/color"
//     "image/draw"
//     "image/png"
//     "golang.org/x/image/font"
//     "golang.org/x/image/font/basicfont"
//     "golang.org/x/image/math/fixed"
// )

func TextToImage(textName string) string { //文字生成图片
	imgfile, _ := os.Create("./static/text/" + textName + ".png")
	defer imgfile.Close()
	//创建位图,坐标x,y,长宽x,y
	img := image.NewNRGBA(image.Rect(0, 0, 400, 60))
	/*
		// 画背景,这里可根据喜好画出背景颜色
		for y := 0; y < dy; y++ {
			for x := 0; x < dx; x++ {
				//设置某个点的颜色，依次是 RGBA
				img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
			}
		}
	*/
	//读字体数据
	fontBytes, err := ioutil.ReadFile("./static/font/SFNS.ttf")
	if err != nil {
		//log.Println(err)
		return ""
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		//log.Println(err)
		return ""
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(40)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)
	//设置字体显示位置
	pt := freetype.Pt(5, 20+int(c.PointToFixed(40)>>8))
	_, err = c.DrawString("推荐码："+textName, pt)
	if err != nil {
		//log.Println(err)
		return ""
	}
	//保存图像到文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
	p := "./static/text/" + textName + ".png"
	return p
}

func createImg(bg, code_url, code string, text_img string) string {
	//背景图
	//如果是windows 换成c:/1.jpg
	backgroudImgFile, err := os.Open(bg) //"./static/1.jpg"

	//tool.Success(c, "success", err)
	fmt.Println(err)

	backgroudImg, _ := jpeg.Decode(backgroudImgFile)

	defer backgroudImgFile.Close()
	backgroudBound := backgroudImg.Bounds()
	//x轴坐标总数
	backgroudX := backgroudBound.Size().X
	//y轴坐标总数
	backgroudY := backgroudBound.Size().Y
	//添加图
	//如果是windows 换成c:/1.jpg
	centerImgFile, _ := os.Open(code_url) //"./static/qrcode/2023-08-23-1692737175-wewe.png"
	centerImg, _ := png.Decode(centerImgFile)
	defer centerImgFile.Close()
	centerBound := centerImg.Bounds()
	//x轴坐标总数
	centerX := centerBound.Size().X
	//y轴坐标总数
	centerY := centerBound.Size().Y

	textFile, err := os.Open(text_img)
	textImg, _ := png.Decode(textFile)
	defer textFile.Close()
	textBound := textImg.Bounds()
	textX := textBound.Size().X
	textY := textBound.Size().Y

	//坐标偏差，x轴y轴 计算
	newImgX := (backgroudX - centerX) / 2
	newImgY := (backgroudY - centerY) - 70
	newTextX := (backgroudX - textX) / 2
	newTextY := (backgroudY - textY) + 5
	offset := image.Pt(newImgX, newImgY)
	offset1 := image.Pt(newTextX, newTextY)
	//x轴坐标总数

	m1 := image.NewRGBA(backgroudBound)
	draw.Draw(m1, backgroudBound, backgroudImg, image.ZP, draw.Src)
	draw.Draw(m1, textImg.Bounds().Add(offset1), textImg, image.ZP, draw.Over)

	//m := image.NewRGBA(backgroudBound)
	//draw.Draw(m, backgroudBound, backgroudImg, image.ZP, draw.Src)
	draw.Draw(m1, centerImg.Bounds().Add(offset), centerImg, image.ZP, draw.Over)
	//如果是windows 换成c:/1.jpg
	u := utils.RandUid(5)
	p := "./static/ok/" + code + u + ".jpg"
	p1 := "/static/ok/" + code + u + ".jpg"
	imgw, _ := os.Create(p)
	//jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	jpeg.Encode(imgw, m1, &jpeg.Options{jpeg.DefaultQuality})
	//jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()
	path := config.Base.BaseUrl + tool.Int_string(config.Base.Port) + p1
	return path
}
