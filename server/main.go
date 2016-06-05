package main

import (
	"log"
	"net/http"
	"os"

	_ "google.golang.org/appengine/cloudsql"
)

func main() {
	connectionString := os.Getenv("MYSQL_CONNECTION")

	db, err := NewDb("mymysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// This is what google app engine uses. should DRY this up
func init() {
	connectionString := os.Getenv("CLOUDSQL_CONNECTION")
	db, err := NewDb("mymysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter(db)

	http.Handle("/", router)
}
