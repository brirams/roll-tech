package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Routes that we can handle
func InitRoutes(repo AlumniRepo) Routes {
	return Routes{
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
			AlumnShow(repo),
		},
		Route{
			"AlumnCreate",
			"POST",
			"/alumni",
			AlumnCreate(repo),
		},
	}

}
