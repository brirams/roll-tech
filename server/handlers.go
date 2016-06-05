package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func AlumniIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(alumni); err != nil {
		panic(err)
	}
}

func AlumnShow(repo AlumniRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var alumnId int
		var alumn Alumn
		var err error

		if alumnId, err = strconv.Atoi(vars["alumnId"]); err != nil {
			panic(err)
		}

		alumn, err = repo.GetAlumn(alumnId)

		if err != nil {
			if err == sql.ErrNoRows {
				// If we didn't find it, 404
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNotFound)
				if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
					panic(err)
				}
				return
			} else {
				panic(err)
			}
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(alumn); err != nil {
				panic(err)
			}
			return
		}
	}
}

func AlumnCreate(repo AlumniRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var alumn Alumn
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			panic(err)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body, &alumn); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		var inserted Alumn
		inserted, err = repo.InsertAlumn(alumn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// cuz i'm a tug like that
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(inserted); err != nil {
			panic(err)
		}
	}
}
