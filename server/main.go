package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := NewDb("mysql", "username:password@tcp(127.0.0.1:3306)/your_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}


// This is what google app engine uses. should DRY this up
func init() {
	db, err := NewDb("mysql", "username:password@tcp(127.0.0.1:3306)/your_db")

	if err != nil {
		log.Fatal(err)
	}

	// running in google app engine complains that the db gets closed
	// by the time we get requests. removing that here but that can't be
	// right...

	router := NewRouter(db)

	http.Handle("/", router)
}
