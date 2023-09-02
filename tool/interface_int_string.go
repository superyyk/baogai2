package tool

func Interface_int(i interface{}) int64  {
	var res int64
	res=i.(int64)
	return res
}
func Interface_string(value interface{}) (string,bool)   {
	switch value.(type) {
	case string:
		op,ok:=value.(string)
		return op,ok

	default:
		return "unkown",false
	}


}


//func test(value interface{}) {
//	switch value.(type) {
//	case string:
//		// 将interface转为string字符串类型
//		op, ok := value.(string)
//		fmt.Println(op, ok)
//	case int32:
//		// 将interface转为int32类型
//		op, ok := value.(int32)
//		fmt.Println(op, ok)
//	case int64:
//		// 将interface转为int64类型
//		op, ok := value.(int64)
//		fmt.Println(op, ok)
//	case User:
//		// 将interface转为User struct类型，并使用其Name对象
//		op, ok := value.(User)
//		fmt.Println(op.Name, ok)
//	case []int:
//		// 将interface转为切片类型
//		op := make([]int, 0)
//		op = value.([]int)
//		fmt.Println(op)
//	default:
//		fmt.Println("unknown")
//	}
//}




