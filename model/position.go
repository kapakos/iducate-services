package model

import "time"

type Position struct {
	Company		string		`json:"company"`
	Title 		string		`json:"title"`
	Location	string		`json:"location"`
	Current 	bool		`json:"current"`
	From		time.Time	`json:"from"`
	To		time.Time	`json:"to"`
}

type Positions []Position