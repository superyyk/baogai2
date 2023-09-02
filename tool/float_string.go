package tool

import "strconv"

func Float_String(num float64) string  {
	str:=strconv.FormatFloat(num, 'G', 30, 32)
	return str
}
