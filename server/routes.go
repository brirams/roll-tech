package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes(repo AlumniRepo) Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"AlumnCreate",
			"POST",
			"/alumni",
			AlumnCreate(repo),
		},
	}

}

