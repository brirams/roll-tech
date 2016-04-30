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
