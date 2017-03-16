package main

import (
	"log"
	"net/http"
	"github.com/kapakos/iducate-services/routing"
)

func main() {

	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}








