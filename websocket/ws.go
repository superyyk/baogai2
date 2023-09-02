package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"poker/model"
	"poker/redis"
	"poker/tool"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Manager struct {
	Group                   map[string]map[string]*Client
	GroupMenber             map[string]interface{}
	groupCount, clientCount uint //uint 正整数
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	BroadCastMessage        chan *BroadCastMessageData
}

// client 单个websocket 信息
type Client struct {
	Id, Group string
	Socket    *websocket.Conn
	Message   chan []byte
}

// 单个发送数据信息 messageData
type MessageData struct {
	Id, Group string
	Message   []byte
}

// 组广播数据信息 GroupMessageData
type GroupMessageData struct {
	Group   string
	Message []byte
}

// 广播数据信息 BroadCastMessageData
type BroadCastMessageData struct {
	Message []byte
}

// 从websocket连接信息中读取数据
func (c *Client) Read() {
	defer func() {
		WebSocketManager.UnRegister <- c
		log.Printf("客户端：%s，断开连接", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("客户端：%s 意外断开连接:%s", c.Id, err)
		}
	}()
	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.Printf("客户端：%s 接收消息：%s", c.Id, string(message))
		//if jsonStr,err:=json.Marshal(message);err!=nil{
		//	log.Print(err.Error())
		//	log.Print(reflect.TypeOf(jsonStr))
		//	WebSocketManager.SendGroup("FME1",message)
		//	//WebSocketManager.Send("00000","FME",message)
		//	WebSocketManager.SendAll(message)
		//}
		c.Message <- message
	}
}

// 写信息，从channel变量send中读取数据写入websocket连接
func (c *Client) Write() {
	defer func() {
		log.Printf("客户端：%s 读信息断开。。。", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("客户端：%s 断开错误：%s", c.Id, err)
		}
	}()
	for {
		select {
		case msg, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("客户端:%s,发送消息：%s", c.Id, string(msg))
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("客户端:%s,发送消息失败:%s", c.Id, err.Error())
			}

		}
	}
}

var WebSocketManager = Manager{
	Group:            make(map[string]map[string]*Client),
	GroupMenber:      map[string]interface{}{},
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	GroupMessage:     make(chan *GroupMessageData, 128),
	Message:          make(chan *MessageData, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	groupCount:       0,
	clientCount:      0,
}

// 启动websocket管理器
func (manager *Manager) Start() {
	log.Print("websocket启动....")
	//tmp:=make(map[string]interface{})
	arr := []string{}
	for {
		select {
		//注册
		case client := <-manager.Register:
			//log.Printf("客户端：%s，连接注册",client.Id)
			log.Printf("把客户端:%s,注册到组:%s", client.Id, client.Group)
			manager.Lock.Lock()
			if manager.Group[client.Group] == nil { //判断当前客户端的组名是否存在
				manager.Group[client.Group] = make(map[string]*Client)
				manager.groupCount += 1
			}
			if manager.Group[client.Group][client.Id] == nil { //判断当前ID有没有在组中

				manager.clientCount += 1
				//log.Print(manager.Group[client.Group])

			}
			manager.Group[client.Group][client.Id] = client
			is_exit := tool.IsInArray(client.Id, arr)
			if !is_exit {
				arr = append(arr, client.Id)
			}
			manager.GroupMenber[client.Group] = arr
			log.Print(manager.Group[client.Group])

			manager.Lock.Unlock()
		//注销
		case client := <-manager.UnRegister:
			log.Printf("注销用户:%s，从组:%s", client.Id, client.Group)
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Group][client.Id]; ok {
				close(client.Message)
				delete(manager.Group[client.Group], client.Id)
				manager.clientCount -= 1
				is_exit := tool.IsInArray(client.Id, arr)
				if is_exit {
					arr = tool.Slice(arr, client.Id)
				}
				manager.GroupMenber[client.Group] = arr
				if len(manager.Group[client.Group]) == 0 {
					delete(manager.Group, client.Group)
					manager.groupCount -= 1

				}

			}
			manager.Lock.Unlock()

			// 发送广播数据到某个组的 channel 变量 Send 中
			//case data := <-manager.boardCast:
			//	if groupMap, ok := manager.wsGroup[data.GroupId]; ok {
			//		for _, conn := range groupMap {
			//			conn.Send <- data.Data
			//		}
			//	}
		}
	}
}

func (manager *Manager) WsClient(ctx *gin.Context) {
	var Rdb = redis.Rdb

	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("连接错误:%s", ctx.Query("channel"))
		return
	}
	fmt.Print(ctx.Query("channel"))
	fmt.Print(ctx.Query("user_uuid"))
	client := &Client{
		Id:      ctx.Query("user_uuid"),
		Group:   ctx.Query("channel"),
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	Rdb.LPush(ctx.Query("channel"), ctx.Query("user_uuid"))
	manager.RegisterClient(client)
	go client.Read()
	go client.Write()
	//time.Sleep(time.Second*15)
	//测试单个client 发送信息
	//manager.Send(client.Id,client.Group,[]byte("发送消息：---"+time.Now().Format("2016-02-03 15:02:21")))
}

// 处理广播数据
func (manager *Manager) SendAllService() {
	for {
		select {
		case data := <-manager.BroadCastMessage:
			for _, v := range manager.Group {
				for _, conn := range v {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理组group广播数据
func (manager *Manager) SendGroupService() {
	for {
		select {
		case data := <-manager.GroupMessage:
			if groupMap, ok := manager.Group[data.Group]; ok {
				for _, conn := range groupMap {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理单个客户端client发送数据
func (manager *Manager) SendService() {
	for {
		select {
		case data := <-manager.Message:
			if groupMap, ok := manager.Group[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
					log.Print("点对点发送信息")
				}
			}

		}
	}
}

// 向指定的client发送消息
func (manager *Manager) Send(id string, group string, message []byte) error {
	data := &MessageData{
		Id:      id,
		Group:   group,
		Message: message,
	}
	manager.Message <- data

	return nil
}

// 注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client

	fmt.Printf("第：%s个用户注册进来了", manager.clientCount)

}

// 注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// 向指定的Group广播
func (manager *Manager) SendGroup(group string, message []byte) error {
	data := &GroupMessageData{
		Group:   group,
		Message: message,
	}
	manager.GroupMessage <- data
	return nil
}

// 向所有人广播
func (manager *Manager) SendAll(message []byte) {
	data := &BroadCastMessageData{
		Message: message,
	}
	manager.BroadCastMessage <- data
}

// 当前组个数
func (manager *Manager) LenGroup() uint {
	return manager.groupCount
}

// 当前连接个数
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// 获取websocketManager 管理器信息
func (manager *Manager) Info() map[string]interface{} {
	jsonStr, err := json.Marshal(manager.Group)
	if err != nil {
		log.Print(err.Error())
	}
	//var client *Client
	managerInfo := make(map[string]interface{})
	managerInfo["groupLen"] = manager.LenGroup()
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnRegisterLen"] = len(manager.UnRegister)
	managerInfo["chanMessageLen"] = len(manager.Message)
	managerInfo["chanGroupmessageLen"] = len(manager.GroupMessage)
	managerInfo["chanBroadCastMessageLen"] = len(manager.BroadCastMessage)
	managerInfo["group_info"] = jsonStr
	managerInfo["md5"] = tool.Md5_2("123")
	managerInfo["is_inarray"] = tool.IsInArray("FME", []string{"FM", "TTk", "yyk"})
	managerInfo["group_menber"] = manager.GroupMenber
	//managerInfo["FME_detail"]=jsonStr

	return managerInfo
}

// 单个发送
func SendToUser(c *gin.Context) {
	ty := c.Query("type")
	user_uuid := c.Query("user_uuid")
	group := c.Query("group")
	data := c.Query("data")
	msg := &model.SongelMsg{
		Ty:    ty,
		Uuid:  user_uuid,
		Group: group,
		Data:  data,
	}
	err := WebSocketManager.Send(user_uuid, group, tool.Struct2Byte_2(msg))
	if err != nil {
		log.Fatalln(err)
	}
	tool.Success(c, "success", 200)
}

// 测试组广播
func SendToGroup(c *gin.Context) {
	//var client *Client
	group := c.Query("group")
	msg := c.Query("data")
	ty := c.Query("type")
	user_id := c.Query("user_id")
	ms := &model.AllMsg{
		Ty:     ty,
		Userid: user_id,
		Group:  group,
		Data:   msg,
		Info:   WebSocketManager.Info(),
	}

	err := WebSocketManager.SendGroup(group, tool.Struct2Byte(ms))
	if err != nil {
		log.Fatalln(err)
	}
	tool.Success(c, "success", 200)
}

// 测试广播
func SendAll(c *gin.Context) {
	//for{
	//	time.Sleep(time.Second*20)
	//} //每个20秒执行一次
	log.Print(c.Query("data"))
	ty := c.Query("type")

	//client :=&Client{}
	data := c.Query("data")
	msg := &model.AllMsg{
		Ty:    ty,
		Group: "all",
		Data:  data,
		Info:  WebSocketManager.Info(),
	}
	WebSocketManager.SendAll(tool.Struct2Byte(msg))

}

func MainSendAll(data interface{}, ty string) {
	msg := &model.AllMsg{
		Ty:    ty,
		Group: "all",
		Data:  data,
		Info:  WebSocketManager.Info(),
	}
	WebSocketManager.SendAll(tool.Struct2Byte(msg))
}

func GetConnectInfo(c *gin.Context) {

	tool.Success(c, "info", WebSocketManager.Info())
}

func OutRoom(c *gin.Context) {
	// user_uuid := c.Query("user_uuid")
	// room_id:=c.Query("room_id")

}
