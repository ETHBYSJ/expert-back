package e

var MsgMap = map[int]string{
	// http返回码
	HTTP_BADREQUEST: "参数错误，请检查",
	// 通用状态码
	SUCCESS: "成功",
	ERROR: "失败",
	// 具体错误码：文件操作
	ERROR_FILE_DOWNLOAD_INVALID_TYPE: "下载文件类型错误",
	ERROR_DOWNLOAD: "下载失败",
	// 具体错误码：专家推荐
	ERROR_RECOMMEND: "提交推荐信息错误",
}

// 获得返回码对应的错误信息
func GetMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}
	return MsgMap[ERROR]
}
