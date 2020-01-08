package main

import (
	"log"
	"net/http"
	"github.com/kapakos/iducate-services/routing"
	"os"
	"fmt"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		// return "", fmt.Errorf("$PORT not set")
		return ":8080", nil
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	router := routing.NewRouter()

	fmt.Println("listening on Port: " + addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	} 
}








