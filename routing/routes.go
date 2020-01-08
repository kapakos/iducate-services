package routing

import (
	"net/http"
)

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandlerFunc 	http.HandlerFunc
}

type Routes []Route

var api = "/api"
var version = "/v1"

var routes = Routes {
	Route {
		"Index",
		"GET",
		api + version + "/",
		Index,
	},
	Route {
		"UserIndex",
		"GET",
		api + version + "/user",
		UserIndex,
	},
	Route {
		"UserShow",
		"GET",
		api + version + "/user/{userId}",
		UserShow,
	},
	Route {
		"UserCreate",
		"Post",
		api + version + "/user",
		UserCreate,
	},
	Route{
		"Status",
		"GET",
		api + version + "/status",
		Status,
	},
	Route {
		"Udacity",
		"Get",
		api + version + "/udacity-courses",
		UdacityCourses,
	},
	Route {
		"Coursera",
		"Get",
		api + version + "/coursera-courses",
		CourseraCourses,
	},
}
