package main

import "time"

type Course struct {
	Name 		string		`json:"name"`
	Institute	string		`json:"institute"`
	Completed	bool		`json:"completed"`
	Score		float32		`json:"score"`
	GraduationDate	time.Time	`json:"graduationDate"`
}

type Courses []Course