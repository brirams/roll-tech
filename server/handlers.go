package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func AlumniIndex(w http.ResponseWriter, r *http.Request) {
	alumni := Alumni{
		Alumn{Name: "Brian Ramos"},
		Alumn{Name: "Michael Parfomak"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(alumni); err != nil {
		panic(err)
	}
}

func AlumnShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alumnId := vars["alumnId"]
	fmt.Fprintf(w, "AlumnShow: %s\n", alumnId)
}
