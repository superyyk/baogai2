package tool

import "sort"

func IsInArray(target interface{}, str_array interface{}) bool { //100万数据 要查50万次得出结果
	switch key:=target.(type) {
	case string:
		for _,item:=range str_array.([]string) {
			if item==key{
				return true
			}
		}
	case int:
		for _,item:=range str_array.([]int){
			if item==key{
				return true
			}
		}

	case int64:
		for _,item:=range str_array.([]int64){
			if item==key{
				return true
			}
		}

	default:
		return false

	}

	return false

}

func IsInArrayFast(target string, str_array []string) bool {   //百万数据20次得出结果，二分法，

	sort.Strings(str_array)

	index := sort.SearchStrings(str_array, target)

	if index < len(str_array) && str_array[index] == target {

		return true

	}

	return false

}

func Slice(arr []string,s string) []string  {
	k:=0
	for key,item:=range arr {
		//fmt.Print(key)
		//fmt.Print(item)
		if item==s{
			k=key
		}
	}
	arr=append(arr[:k],arr[(k+1):]...)
	return arr
}
