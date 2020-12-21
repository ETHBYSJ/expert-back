package vo

type SearchVO struct {
	Keyword string 		`json:"keyword"`
	Labels 	[]string 	`json:"labels"`
}
