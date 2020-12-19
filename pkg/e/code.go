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
	ErrorRecommendParse     = 30002
	ErrorRecommendGet       = 30003
	// 具体错误码：专家申请
	ErrorApplyCreate = 40001
	ErrorApplyUpdate = 40002
	// 其他错误码
	ErrorGetAccountProfile = 50001
)
