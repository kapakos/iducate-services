package model

import "time"

type Education struct {
	School		string		`json:"school"`
	Degree		string		`json:"degree"`
	From		time.Time	`json:"from"`
	To		time.Time	`json:"to"`
	Field		string		`json:"field"`
	Grade		float32		`json:"grade"`
	Description	string		`json:"description"`
}

type EducationList []Education