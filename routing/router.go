package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapakos/iducate-services/middleware"
	"log"
)

func NewRouter() *mux.Router {
	log.Println("I'm in Router")

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println("I'm in loop " + route.Name)

		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
