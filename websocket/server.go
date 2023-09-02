package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Message struct {
	//Type int64 //0.消息 1.订单，2。互转 3.提现审核
	Data []byte
	Roomid string
	conn *Connection
	Ip string
}

type Ws struct {
	Conn *websocket.Conn
}



type hub struct {
	rooms map[string]map[*Connection]bool
	broadcast chan Message
	broadcastss chan Message
	warnings chan Message
	register chan Message
	unregister chan Message
	kickoutroom chan Message
	warmsg chan Message
}

var H=hub{
	rooms:make(map[string]map[*Connection]bool),
	broadcast:make(chan Message),
	broadcastss:make(chan Message),
	warnings:make(chan Message),
	warmsg:make(chan Message),
	register:make(chan Message),
	unregister:make(chan Message),
	kickoutroom:make(chan Message),

}

func (h *hub)Run()  {
	msg:=&System{
		Type:"system",
		Msg:&SystemMsg{
            Id:1,
            Type:"text",
            Content:&TextContent{
            	Text:"欢迎进入房间",
			},
		},
	}

	for{
		select {
		case m:=<-h.register:
			conns:=h.rooms[m.Roomid]

			if conns==nil{
				conns=make(map[*Connection]bool)
				
				h.rooms[m.Roomid]=conns
				fmt.Print("在线人数：==",len(conns))
				fmt.Print("房间:=",h.rooms)
			}
			h.rooms[m.Roomid][m.conn]=true
			fmt.Println("房间：",h.rooms)
			fmt.Println("在线人数：==",len(conns))
			for con :=range conns{
				ms,err:=json.Marshal(msg)
				if err!=nil{
					fmt.Print(err)
				}

				delmsg:=[]byte(ms)

				data:=[]byte(delmsg)
				select {
				case con.send<-data:
					fmt.Println("用户进入房间并注册了房间"+m.Ip)
				}
			}

		case m:=<-h.unregister: //断开连接
		conns:=h.rooms[m.Roomid]
		if conns!=nil{
			if _,ok:=conns[m.conn];ok{
				delete(conns,m.conn)
				close(m.conn.send)
				for con :=range conns{
					delmsg:="系统消息有小伙伴离开了房间："+m.Roomid
					data:=[]byte(delmsg)
					select {
					case con.send<-data:
					}
					if len(conns)==0{ //连接都断开
						delete(h.rooms,m.Roomid)
					}
				}
			}
		}

		case m:=<-h.kickoutroom : //3次不合法信息后，被提出
		 conns:=h.rooms[m.Roomid]
		 notice:="由于您多次发送不合法信息，已被提出群聊"
			select {
		     case m.conn.send<-[]byte(notice):

		 }
		 if conns!=nil{
		 	if _,ok:=conns[m.conn];ok{
		 		delete(conns,m.conn)
		 		close(m.conn.send)
		 		if len(conns)==0{
		 			delete(h.rooms,m.Roomid)
				}
			}
		 }

		case m:=<-h.warnings: //不合法信息
		conns:=h.rooms[m.Roomid]
		if conns!=nil{
			if _,ok:=conns[m.conn];ok{
				notice:="警告：您发布了不合法的信息，将禁言5分钟"
				select {
				    case m.conn.send<-[]byte(notice):
				    	fmt.Print(notice)

				}
			}
		}

        case m:=<-h.warmsg: //禁言中提示
        	conns:=h.rooms[m.Roomid]
        	if conns !=nil{
        		if _,ok:=conns[m.conn];ok{
        			notice:="您还在禁言中，暂时不能发送消息"
					select {
        			  case m.conn.send<-[]byte(notice):
					}
				}
			}

        case m:=<-h.broadcast: //传输群消息/房间消息
        fmt.Print("有消息！")
        conns:=h.rooms[m.Roomid]
        for con:=range conns{
        	//if con==m.conn{  //自己发送消息，不能再发给自己
        	//	continue
			//}
			select {
        	  case con.send<-m.Data:
        	  fmt.Print("发送成功！")

			default:
				close(con.send)
				delete(conns,con)
				if len(conns)==0{
					delete(h.rooms,m.Roomid)
				}
			}
		}

		}
	}
}


