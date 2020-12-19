package e

const (
	// http返回码
	HttpBadRequest = 400
	// 通用状态码
	Success = 10000
	Error = 10001
	// 具体错误码：文件操作
	ErrorDownload = 20001
	ErrorUpload   = 20002
	// 具体错误码：专家推荐
	ErrorRecommend = 30001
	ErrorParse     = 30002
	ErrorGet       = 30003
	// 其他错误码
	ErrorGetAccountProfile = 40001
)
