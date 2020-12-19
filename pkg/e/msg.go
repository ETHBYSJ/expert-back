package e

var MsgMap = map[int]string{
	// http返回码
	HttpBadRequest: "参数错误，请检查",
	// 通用状态码
	Success: "成功",
	Error: "失败",
	// 具体错误码：文件操作
	ErrorDownload: "下载失败",
	ErrorUpload:   "上传失败",
	// 具体错误码：专家推荐
	ErrorRecommend: "提交推荐信息错误",
	ErrorParse:     "解析推荐表失败",
	ErrorGet:       "获取推荐信息失败",
	// 其他错误码
	ErrorGetAccountProfile: "获取用户信息失败",
}

// 获得返回码对应的错误信息
func GetMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}
	return MsgMap[Error]
}
