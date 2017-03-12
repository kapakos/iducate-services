package main

import (
	"log"
	"net/http"
	"../routing"
)

func main() {

	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}








