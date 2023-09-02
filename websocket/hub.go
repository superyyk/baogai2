package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"

)

const(
	writeWait=10*time.Second
	pongWait=60*time.Second
	pingPeriod=(pongWait*9)/10
	maxMessageSize=5120
)

type Connection struct {
	ws *websocket.Conn
	send chan []byte
	numberv int
	forbiddenword bool
	timelog int64
}

var (upgrader=websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
})

func (m Message)ReadPump()  {
	c:=m.conn

	msg:=&UserText{}
	msg1:=&Chongzhi{}

	defer func() {
		H.unregister<-m
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for  {
		err:=c.ws.ReadJSON(&msg)
		if err!=nil{
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway){
				fmt.Print("error:%v",err)
			}
			break
		}

		//fmt.Print("收到消息:%v",msg.Type)
		jsonStr,err:=json.Marshal(msg);
		if err!=nil{
			fmt.Print(err)

		}
        //c.send<-[]byte(jsonStr)

        fmt.Print(string(jsonStr))

		//m.Kickout(jsonStr)
		jsonStr1,err:=json.Marshal(msg1)
		if(err!=nil){
			fmt.Print(err.Error())
		}
		m.Fenlei(jsonStr1)

	}

}
func (m Message)Fenlei(msg []byte)  {
	c:=m.conn
	jsonStr:=&Chongzhi{}

	err:=json.Unmarshal(msg,jsonStr)
	if(err!=nil){
		fmt.Print(err.Error())
	}
	mm:=Message{msg,m.Roomid,c,c.ws.RemoteAddr().String()}
	H.broadcast<-mm

	switch jsonStr.Type {
	     case 0:
			 m1:=Message{msg,m.Roomid,c,c.ws.RemoteAddr().String()}
			 H.broadcast<-m1
	     	break
	default:
		break

	}

}

//信息处理，不合法言论，禁言警告，超过3次，提出群聊
func (m Message)Kickout(msg []byte)  {
	c:=m.conn
	//判断是否有禁言时间，并超过10秒禁言时间，没有超过进入禁言提醒
	nowT:=int64(time.Now().Unix())
	if nowT-c.timelog<10{
		H.warmsg<-m
	}
	//不合法信息3次，判断是否有不合法信息，没有进行消息提示
	if c.numberv<3{
		//basestr:="死妈儿孙子爸父B逼狗日FW废"
		basestr:=""
		teststr:=string(msg[:])
		for _,ev:=range teststr{
			//判断字符串中是否含有某个自负，true/false
			result:=strings.Contains(basestr,string(ev))
			if result==true{
				c.numberv+=1
				c.forbiddenword=true//禁言为真
				//记录禁言开始时间，禁言时间内任何信息不能发送
				c.timelog=int64(time.Now().Unix())
				H.warnings<-m
				break
			}
		}
		//不禁言，消息合法 可以发送
		if c.forbiddenword!=true{
			// 设置广播消息, 所有房间内都可以收到信息;给广播消息开头加一个特定字符串为标识，当然也有其他方法;
			// 此例 设置以开头0为标识, 之后去掉0 ;
			if msg[0]==48{
				head:=string("所有玩家请注意：")
				data:=head+string(msg[1:])
				m:=Message{[]byte(data),m.Roomid,c,c.ws.RemoteAddr().String()}
				H.broadcastss<-m
			}else if msg[0]!=48{ //不是0，就是普通消息
				m:=Message{msg,m.Roomid,c,c.ws.RemoteAddr().String()}
				H.broadcast<-m
			}

		}

	}else {
		H.kickoutroom<-m
		log.Println("您要被提出群聊了...")
		c.ws.Close() //此处关闭了提出的连接，也可以不关闭做其他处理
	}


}

func (c *Connection)write(mt int,payload []byte) error  {
c.ws.SetWriteDeadline(time.Now().Add(writeWait))
return c.ws.WriteMessage(mt,payload)
}

func (s Message)writePump()  {
	c:=s.conn
	ticker:=time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for{
		select {
		case message,ok:=<-c.send:
			if !ok{
				c.write(websocket.CloseMessage,[]byte{})
				return
			}


			if err:=c.write(websocket.TextMessage,message);err!=nil{
				return
			}
			case <-ticker.C:
				if err:=c.write(websocket.PingMessage,[]byte{});err!=nil{
					return
				}



		}
	}
}


func ServerWs(c *gin.Context)  {

      roomid:=c.Query("roomid")

      ws,err:=upgrader.Upgrade(c.Writer,c.Request,nil)
      if err!=nil{
      	log.Print(err)
		  return
	  }
      cc:=&Connection{send:make(chan []byte),ws:ws}
      m:=Message{nil,roomid,cc,ws.RemoteAddr().String()}
      H.register<-m
      go m.writePump()
      go m.ReadPump()

}
