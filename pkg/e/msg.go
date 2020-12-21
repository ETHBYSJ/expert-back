package e

var MsgMap = map[int]string{
	// http返回码
	HttpBadRequest: "参数错误，请检查",
	// 通用状态码
	Success: "成功",
	Error:   "失败",
	// 具体错误码：文件操作
	ErrorDownload: "下载失败",
	ErrorUpload:   "上传失败",
	// 具体错误码：专家推荐
	ErrorRecommend:           "提交推荐信息错误",
	ErrorRecommendParse:      "解析推荐表失败",
	ErrorRecommendGetExperts: "获取推荐专家信息失败",
	ErrorRecommendRecordsGet: "获取推荐记录失败",
	ErrorRecommendSubmitGet:  "获取提交信息失败",
	ErrorRecommendRecordSet: "更新推荐记录失败",
	ErrorRecommendFileRecordSet: "更新推荐文件记录失败",
	ErrorRecommendFileRecordGet: "获取推荐文件记录失败",
	// 具体错误码：专家申请
	ErrorApplyCreate:     "创建专家申请失败",
	ErrorApplyUpdate:     "更新申请信息失败",
	ErrorApplyRecordsGet: "获取申请记录失败",
	ErrorApplyGet: "获取申请信息失败",
	ErrorApplyRecordSet: "更新申请记录失败",
	ErrorApplyFileRecordSet: "更新申请文件记录失败",
	ErrorApplyFileRecordGet: "获取申请文件记录失败",
	ErrorApplySaveInvertedIndex: "保存倒排索引错误",
	// 具体错误码：搜索
	ErrorSearch: "搜索出错",
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
