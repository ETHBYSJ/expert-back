package vo

type ApplyBaseVO struct {
	Name            string `json:"name" bson:"name"`                       // 姓名
	Sex             string `json:"sex" bson:"sex"`                         // 性别
	Birthday        string `json:"birthday" bson:"birthday"`               // 出生年月
	Nation          string `json:"nation" bson:"nation"`                   // 民族
	Phone           string `json:"phone" bson:"phone"`                     // 办公电话
	Email           string `json:"email" bson:"email"`                     // 邮箱
	Photo           string `json:"photo" bson:"photo"`                     // 证件照片url
	PoliticalStatus string `json:"politicalStatus" bson:"politicalStatus"` // 政治面貌
	Health          string `json:"health" bson:"health"`                   // 健康状况
	HomePhone       string `json:"homePhone" bson:"homePhone"`             // 住宅电话
	Mobile          string `json:"mobile" bson:"mobile"`                   // 手机号码
	WorkAddress     string `json:"workAddress" bson:"workAddress"`         // 单位地址
	WorkPostcode    string `json:"workPostcode" bson:"workPostcode"`       // 单位邮政编码
	HomeAddress     string `json:"homeAddress" bson:"homeAddress"`         // 家庭地址
	HomePostcode    string `json:"homePostcode" bson:"homePostcode"`       // 家庭邮政编码
}

type ApplyMajorVO struct {
	Qualification          string `json:"qualification" bson:"qualification"`                   // 学历程度
	Degree                 string `json:"degree" bson:"degree"`                                 // 最高学位
	Major                  string `json:"major" bson:"major"`                                   // 所学专业
	MajorCategory          string `json:"majorCategory" bson:"majorCategory"`                   // 专业类别
	Workplace              string `json:"workplace" bson:"workplace"`                           // 工作单位
	ProfessionalPosition   string `json:"professionalPosition" bson:"professionalPosition"`     // 专业技术职务
	AdministrativePosition string `json:"administrativePosition" bson:"administrativePosition"` // 现任行政职务
	Authority              string `json:"authority" bson:"authority"`                           // 单位主管部门
	WorkTime               string `json:"workTime" bson:"workTime"`                             // 工作时间
}

type ApplyResearchFieldVO struct {
	MajorLabels        []string `json:"majorLabels" bson:"majorLabels"`               // 从事专业或学科
	ResearchDirections []string `json:"researchDirections" bson:"researchDirections"` // 研究方向或专长
}

type ApplyResumeVO struct {
	WorkExperience string `json:"workExperience" bson:"workExperience"` // 工作经历
	Achievements   string `json:"achievements" bson:"achievements"`     // 工作成绩
}

type ApplyOpinionVO struct {
	UnitOpinion      string `json:"unitOpinion" bson:"unitOpinion"`           // 工作单位意见
	AuthorityOpinion string `json:"authorityOpinion" bson:"authorityOpinion"` // 主管部门意见
}
