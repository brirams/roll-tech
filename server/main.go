package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := NewDb("mysql", "root:k3nn@1214@tcp(127.0.0.1:3306)/roll_tech_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
