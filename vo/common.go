package vo

// 一条记录
type CommonRecordVO struct {
	Title     string `json:"title" bson:"title"`
	Status    string `json:"status" bson:"status"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
