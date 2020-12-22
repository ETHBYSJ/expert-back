package vo

// 一条记录
type CommonRecordVO struct {
	Title     string `json:"title" bson:"title"`
	Status    string `json:"status" bson:"status"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

type ReviewRecommendVO struct {
	SubmitID	string 	`form:"submitID"`
	Status 		int 	`form:"status"`
}

type ReviewApplyVO struct {
	UserID 		string 	`form:"userID"`
	Status 		int 	`form:"status"`
}
