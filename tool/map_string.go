package tool

func MapToSlice(m map[string]interface{}) []interface{} {
	s := make([]interface{}, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}


