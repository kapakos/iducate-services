package routing

import (
	"net/http"
	"../middleware"
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
	Route{
		"SetToken",
		"POST",
		api + version + "/set-token",
		SetToken,
	},
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
		middleware.Validate(UserIndex),
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
		"Logout",
		"Get",
		api + version + "/logout",
		Logout,
	},
	Route {
		"Udacity",
		"Get",
		api + version + "/udacity-courses",
		UdacityCourses,
	},
}
