package utils

import "strings"

// 贴一下我自己手写的解析url参数
// url参数转换为map
func UrlParamToMap(param string) map[string]string {
	split := strings.Split(param, "&")
	maps := make(map[string]string)
	for _, v := range split {
		i := strings.Split(v, "=")
		maps[i[0]] = i[1]
	}
	return maps
}
