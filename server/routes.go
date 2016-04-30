package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AlumniIndex",
		"GET",
		"/alumni",
		AlumniIndex,
	},
	Route{
		"AlumnShow",
		"GET",
		"/alumni/{alumnId}",
		AlumnShow,
	},
	Route{
		"AlumnCreate",
		"POST",
		"/alumni",
		AlumnCreate,
	},
}
