package utils

import (
	"time"
)

func GetNowTime()string  { //获取当前时间 年-月-日 小时-分钟-秒
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowTimeStamp()int64  {  //获取当前时间戳
	return time.Now().Unix()
}

func ToDate(timestamp int64)  string{ //时间戳转日期

	return time.Unix(timestamp,0).Format("2006-01-02 15:04:05")
}

func ToTimeStamp(date string ) int64{ //日期转时间戳
	formatTime,err:=time.Parse("2006-01-02 15:04:05",date)

	if err==nil{

		//fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC
		return formatTime.Unix()

	}else {
		return formatTime.Unix()
	}
}
