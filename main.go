package main

import (
	"log"
	"net/http"
	"github.com/kapakos/iducate-services/routing"
	"os"
	"strconv"
	"fmt"
)

var (
	repeat int
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("I'm in Main")
	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(addr, router))
}








