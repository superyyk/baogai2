package tool

import (
	"encoding/json"
	"log"

	"github.com/superyyk/yishougai/model"
)

type Req struct {
	Req interface{}
}

func Struct2Byte(msg *model.AllMsg) []byte {
	var res []byte

	if jsonStr, err := json.Marshal(msg); err != nil {
		log.Print(err.Error())

	} else {
		res = jsonStr
	}
	return res
}

func Struct2Byte_2(msg *model.SongelMsg) []byte {
	var res []byte

	if jsonStr, err := json.Marshal(msg); err != nil {
		log.Print(err.Error())

	} else {
		res = jsonStr
	}
	return res
}
