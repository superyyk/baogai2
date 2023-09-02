package main

import (
	"github.com/superyyk/yishougai/config"
	"github.com/superyyk/yishougai/route"
	"github.com/superyyk/yishougai/websocket"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}

func main() {
	r := route.NewRoute()
	// 加载静态资源
	r.Static("/static", "./static")
	//r.Static("/static/qrcode", "./static/qrcode")
	r.Static("/api/static", "./api/static")
	//ws2服务
	{
		go websocket.WebSocketManager.Start()
		go websocket.WebSocketManager.SendService()
		go websocket.WebSocketManager.SendGroupService()
		go websocket.WebSocketManager.SendAllService()
		go websocket.WebSocketManager.Start()
		go websocket.WebSocketManager.SendService()
		go websocket.WebSocketManager.SendGroupService()
		go websocket.WebSocketManager.SendAllService()
		//go websocket.TestSendGroup()
		//go websocket.TestSendAll()

		r.GET("/ws2", websocket.WebSocketManager.WsClient)
		r.GET("send_to_all", websocket.SendAll)
		r.GET("send_to_group", websocket.SendToGroup)
		r.GET("send_to_user", websocket.SendToUser)
		r.GET("get_connect_info", websocket.GetConnectInfo)
		r.GET("out_room", websocket.OutRoom)
	}

	port := config.Base.Port
	//r.Run(":" + strconv.Itoa(port))
	r.Use(TlsHandler(port))
	r.RunTLS(":"+strconv.Itoa(port), "./my_ssl/fullchain.pem", "./my_ssl/privkey.key")
}
