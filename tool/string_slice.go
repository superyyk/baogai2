package tool

import "strings"

func String_slice(st string) []string {
	s1 := strings.ReplaceAll(st, "[", "")
	s2 := strings.ReplaceAll(s1, "]", "")

	s3 := strings.Split(s2, ",")
	var m []string
	for _, v := range s3 {
		str := v
		if str[0] == '"' {
			str = str[1:]
		}
		if i := len(str) - 1; str[i] == '"' {
			str = str[:i]
		}
		m = append(m, str)
	}

	return m
}
