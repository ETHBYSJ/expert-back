package vo

// 工作单位
type RecommendCompanyVO struct {
	Name 		string	`json:"name" bson:"name"`			// 填报单位
	Director 	string	`json:"director" bson:"director"`	// 单位负责人
	Agent 		string	`json:"agent" bson:"agent"`			// 经办人
	Phone 		string 	`json:"phone" bson:"phone"`			// 联系电话
}

// 专家推荐
type RecommendExpertVO struct {
	Name 			string 	`json:"name" bson:"name"`					// 姓名
	Sex 			string 	`json:"sex" bson:"sex"`						// 性别
	Age 			int		`json:"age" bson:"age"`						// 年龄
	Qualification	string 	`json:"qualification" bson:"qualification"`	// 学历
	Title 			string 	`json:"title" bson:"title"`					// 职称
	Major 			string 	`json:"major" bson:"major"`					// 专业或学科
	Company 		string 	`json:"company" bson:"company"`				// 工作单位
	Duty 			string 	`json:"duty" bson:"duty"`					// 行政职务
	Phone 			string 	`json:"phone" bson:"phone"`					// 办公电话
	Mobile 			string  `json:"mobile" bson:"mobile"`				// 手机
	Email 			string 	`json:"email" bson:"email"`					// 电子邮箱
}

// 推荐申请
type RecommendVO struct {
	// UserID  string 				`json:"id"`
	RecommendCompanyVO			`json:"company"`
	List []RecommendExpertVO	`json:"list"`
}

/*
type RecommendIDVO struct {
	UserID 	string		`form:"id"`
}
*/
// 解析后的结果
type RecommendParseVO struct {
	List 	[]RecommendExpertVO	`json:"list"`
}