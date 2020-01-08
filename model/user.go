package model

import "time"

type User struct {
	Id		int 		`json:"id"`
	FirstName	string		`json:"firstName"`
	MiddleName 	string		`json:"middleName"`
	LastName	string		`json:"lastName"`
	BirthDate	time.Time	`json:"birthDate"`
	Description	string		`json:"description"`
	Created		time.Time	`json:"created"`
	EducationList	EducationList	`json:"educationList"`
	SkillSet	SkillSet	`json:"skillSet"`
	Positions	Positions	`json:"positions"`
}

type Users []User