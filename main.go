package main

import (
	"log"
	"net/http"
	"github.com/kapakos/iducate-services/routing"
	"os"
	"strconv"
)

var (
	repeat int
)

func main() {
	var err error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err = strconv.Atoi(tStr)
	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
	}
	log.Println("I'm in Main")
	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(":" + port, router))
}








