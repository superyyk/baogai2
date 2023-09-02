package tool

import "encoding/json"

type RequestBody struct {
	req string
}
//json转map
func (r *RequestBody)Json2map()(s map[string]interface{},err error)  {
	var result map[string]interface{}
	if err:=json.Unmarshal([]byte(r.req),&result);err!=nil{
		return nil,err
	}
	return result,nil
}
