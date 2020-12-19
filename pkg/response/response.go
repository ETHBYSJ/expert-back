package response

import (
	"expert-back/pkg/e"
)

const (
	Code = 1
	Data = 2
	Msg = 3
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
	if _, ok := m[Code]; ok {
		code = m[Code].(int)
	} else {
		code = e.Success
	}
	if _, ok := m[Msg]; ok {
		msg = m[Msg].(string)
	} else {
		msg = e.GetMsg(code)
	}
	if _, ok := m[Data]; ok {
		return Response{
			Code: code,
			Data: m[Data],
			Msg: msg,
		}
	} else {
		return Response{
			Code: code,
			Msg: msg,
		}
	}
}


