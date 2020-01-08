package data

import (
	"time"
	"fmt"
	"github.com/kapakos/iducate-services/model"
)

var currentId int
var Users model.Users

func init() {
	RepoCreateUser(
		model.User{
			FirstName: "Petros",
			LastName: "Kapakos",
			BirthDate: time.Date(1977, time.January, 16, 0, 0, 0, 0, time.UTC),
			Description: "An Experienced Web Engineer",
		})
	RepoCreateUser(
		model.User{
			FirstName: "Elias",
			LastName: "Kapakos",
			BirthDate: time.Date(2011, time.January, 11, 0, 0, 0, 0, time.UTC),
			Description: "An Experienced sweats eater",
		})
	RepoCreateUser(
		model.User{
			FirstName: "Leos",
			LastName: "Kapakos",
			BirthDate: time.Date(2017, time.February, 9, 0, 0, 0, 0, time.UTC),
			Description: "An Experienced Pooper",
		})
}

func RepoCreateUser(u model.User) model.User {
	currentId += 1
	u.Id = currentId
	u.Created = time.Now()
	Users = append(Users, u)
	return u
}

func RepoFindUser(id int) model.User {
	for _, user := range Users {
		if user.Id == id {
			return user
		}
	}
	return model.User{}
}

func RepoDestroyUser(id int) error {
	for i, user := range Users {
		if user.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}