package response

import (
	"expert-back/pkg/e"
)

const (
	CODE = 1
	DATA = 2
	MSG = 3
)

type Response struct {
	Code 	int 		`json:"code"`
	Data 	interface{} `json:"data,omitempty"`
	Msg 	string 		`json:"msg"`
}

// 构造返回体
func BuildResponse(m map[int]interface{}) Response {
	var code int
	var msg string
	if _, ok := m[CODE]; ok {
		code = m[CODE].(int)
	} else {
		code = e.SUCCESS
	}
	if _, ok := m[MSG]; ok {
		msg = m[MSG].(string)
	} else {
		msg = e.GetMsg(code)
	}
	if _, ok := m[DATA]; ok {
		return Response{
			Code: code,
			Data: m[DATA],
			Msg: msg,
		}
	} else {
		return Response{
			Code: code,
			Msg: msg,
		}
	}
}


