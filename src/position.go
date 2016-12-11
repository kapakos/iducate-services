package main

import "time"

type Position struct {
	Company		string		`json:"company"`
	Title 		string		`json:"title"`
	Location	string		`json:"location"`
	Current 	bool		`json:"current"`
	DateFrom	time.Time	`json:"from"`
	DateTo		time.Time	`json:"to"`
}

type Positions []Position