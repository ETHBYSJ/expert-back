package vo

type SearchVO struct {
	Keyword string 		`json:"keyword"`
	Labels 	[]string 	`json:"labels"`
}

type SearchResultVO struct {
	Labels 			[]string 	`json:"labels"`
	Name 			string 		`json:"name"`
	Photo			string 		`json:"photo"`
	Intro			string 		`json:"intro"`
}