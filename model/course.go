package model

type Course struct {
	Title 		string		`json:"title"`
	Subtitle	string		`json:"subtitle"`
	Homepage	string		`json:"homepage"`
	Summary		string		`json:"summary"`
	ShortSummary	string		`json:"short_summary"`
	Key		string		`json:"key"`
}

type Courses []Course