package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	user := User {
		FirstName: "Petros",
		LastName: "Kapakos",
		BirthDate: time.Date(1977, time.January, 16,0,0,0,0, time.UTC),
		Description: "An Experienced Web Engineer",
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["userId"]
	fmt.Fprintln(w, "User show:", todoId)
}