package vo

type ApplyBaseVO struct {
	Name            string `json:"name" bson:"name"`                       	// 姓名
	Sex             string `json:"sex" bson:"sex"`                         	// 性别
	Born        	string `json:"born" bson:"born"`               			// 出生年月
	Nation          string `json:"nation" bson:"nation"`                   	// 民族
	Phone           string `json:"phone" bson:"phone"`                    	// 办公电话
	Email           string `json:"email" bson:"email"`                     	// 邮箱
	Photo           string `json:"photo" bson:"photo"`                     	// 证件照片url
	Politics 		string `json:"politics" bson:"politics"` 				// 政治面貌
	Health          string `json:"health" bson:"health"`                   	// 健康状况
	Homecall       	string `json:"homecall" bson:"homecall"`             	// 住宅电话
	Mobile          string `json:"mobile" bson:"mobile"`                   	// 手机号码
	DeptAddr     	string `json:"deptAddr" bson:"deptAddr"`         		// 单位地址
	DeptCode    	string `json:"deptCode" bson:"deptCode"`       			// 单位邮政编码
	HomeAddr     	string `json:"homeAddr" bson:"homeAddr"`         		// 家庭地址
	HomeCode    	string `json:"homeCode" bson:"homeCode"`       			// 家庭邮政编码
}

type ApplyMajorVO struct {
	Edu          			string `json:"edu" bson:"edu"`                   	// 学历程度
	Degree                 	string `json:"degree" bson:"degree"`             	// 最高学位
	Major                  	string `json:"major" bson:"major"`              	// 所学专业
	Category          		string `json:"category" bson:"category"`        	// 专业类别
	Dept              		string `json:"dept" bson:"dept"`                 	// 工作单位
	TechPost   				string `json:"techPost" bson:"techPost"`     		// 专业技术职务
	AdminPost 				string `json:"adminPost" bson:"adminPost"` 			// 现任行政职务
	AuthDept              	string `json:"authDept" bson:"authDept"`            // 单位主管部门
	WorkTime               	string `json:"workTime" bson:"workTime"`            // 工作时间
}

type ApplyResearchFieldVO struct {
	MajorLabels        	[]string `json:"majorLabels" bson:"majorLabels"`               	// 从事专业或学科
	ResearchLabels		[]string `json:"researchLabels" bson:"researchLabels"` 			// 研究方向或专长
}

type ApplyResumeVO struct {
	WorkExperience string `json:"workExperience" bson:"workExperience"` // 工作经历
	Achievements   string `json:"achievements" bson:"achievements"`     // 工作成绩
}

type ApplyOpinionVO struct {
	DeptOpinion      	string `json:"deptOpinion" bson:"deptOpinion"`           	// 工作单位意见
	AuthDeptOpinion 	string `json:"authDeptOpinion" bson:"authDeptOpinion"` 		// 主管部门意见
}
