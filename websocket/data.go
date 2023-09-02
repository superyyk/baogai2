package websocket

//文本内容
type TextContent struct {
	Text string `json:"text"`
}
//图片内容
type ImgContent struct {
	Url string `json:"url"`
	W int64 `json:"w"`
	H int64 `json:"h"`
}
//声音内容
type VoiceContent struct {
	Url string `json:"url"`
	Length string `json:"length"`
}
//红包内容
type HongbaoContent struct {
	Text string `json:"blessing"`
	Id string `json:"rid"`
	IsRecevied bool `json:"isReceived"`
}
//视频内容
type VidoContent struct {
	Url string `json:"url"`

}
type UserInfo struct {
	Uid string `json:"uid"`
	//UserUUid string `json:"user_uuid"`
	Username string `json:"username"`
	Face string `json:"face"`
}

//系统消息
type SystemMsg struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Content *TextContent `json:"content"`
}
//文本消息
type TextMsg struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	UserInfo *UserInfo `json:"userinfo"`
	Content *TextContent `json:"content"`
}
//图片消息
type ImgMsg struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	UserInfo *UserInfo `json:"userinfo"`
	Content *ImgContent `json:"content"`
}
//声音消息
type VoiceMsg struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	UserInfo *UserInfo `json:"userinfo"`
	Content *VoiceContent `json:"content"`

}
//红包消息
type HongbaoMsg struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	UserInfo *UserInfo `json:"userinfo"`
	Content *HongbaoContent `json:"content"`
}
//视频消息
type VidoMsg struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	UserInfo *UserInfo `json:"userinfo"`
	Content *VoiceContent `json:"content"`
}

type System struct {
	Type string `json:"type"`
	Msg *SystemMsg `json:"msg"`
}

type UserText struct {
	Type string `json:"type"`
	Msg *TextMsg `json:"msg"`
}

type Chongzhi struct {
	Type int `json:"type"`
	From_id string `json:"from_id"`
	To_id string `json:"to_id"`
	Order_id string `json:"order_id"`

}

type UserImg struct {
	Type string `json:"type"`
	Msg *ImgMsg `json:"msg"`
}

type UserVoice struct {
	Type string `json:"type"`
	Msg *VoiceMsg `json:"msg"`
}

type UserHongbao struct {
	Type string `json:"type"`
	Msg *HongbaoMsg `json:"msg"`
}

type UserVido struct {
	Type string `json:"type"`
	Msg *VidoMsg `json:"msg"`
}

type Data struct {
	Ip string `gorm:"fID",json:"ip"`
	Type string `gorm:"fType",json:"type"`
	FromID string `gorm:"fFrom_id",json:"from_id"`
	ToID string `gorm:"fTo_id",json:"to_id"`
	Content string `gorm:"fContent",json:"content"`
	UserList []string `gorm:"fUser_list",json:"user_list"`
}






