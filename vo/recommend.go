package vo

// 工作单位
type RecommendDepartmentVO struct {
	Name     string `json:"name" bson:"name"`         // 填报单位，用于区分单位
	Director string `json:"director" bson:"director"` // 单位负责人
	Agent    string `json:"agent" bson:"agent"`       // 经办人
	Phone    string `json:"phone" bson:"phone"`       // 联系电话
}

// 专家推荐
type RecommendExpertVO struct {
	Name   string `json:"name" bson:"name"`     // 姓名
	Sex    string `json:"sex" bson:"sex"`       // 性别
	Age    string `json:"age" bson:"age"`       // 年龄
	Edu    string `json:"edu" bson:"edu"`       // 学历
	Title  string `json:"title" bson:"title"`   // 职称
	Major  string `json:"major" bson:"major"`   // 专业或学科
	Dept   string `json:"dept" bson:"dept"`     // 工作单位
	Post   string `json:"post" bson:"post"`     // 行政职务
	Phone  string `json:"phone" bson:"phone"`   // 办公电话
	Mobile string `json:"mobile" bson:"mobile"` // 手机
	Email  string `json:"email" bson:"email"`   // 电子邮箱
}

// 推荐申请
type RecommendVO struct {
	RecommendDepartmentVO `json:"department"`
	List                  []RecommendExpertVO `json:"list"`
	SubmitID              string              `json:"submitID"`
}

// 用于返回给前端
type RecommendRetVO struct {
	RecommendVO
	File 	string 	`json:"file"`
}

// 解析后的结果
type RecommendParseVO struct {
	List []RecommendExpertVO `json:"list"`
}

// 根据提交id获取信息
type RecommendGetSubmitVO struct {
	SubmitID string `form:"submitID"`
}

// 上传文件
type RecommendUploadVO struct {
	SubmitID string `form:"submitID"`
}

// 删除文件
type RecommendDeleteVO struct {
	SubmitID string `form:"submitID"`
}