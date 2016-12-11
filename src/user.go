package main

import "time"

type User struct {
	FirstName	string		`json:"firstName"`
	MiddleName 	string		`json:"middleName"`
	LastName	string		`json:"lastName"`
	BirthDate	time.Time	`json:"birthDate"`
	Description	string		`json:"description"`
}