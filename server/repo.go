package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	*sql.DB
}

func NewDb(driverName string, creds string) (Db, error) {
	db, err := sql.Open(driverName, creds)
	return Db{db}, err
}

func (d Db) GetAlumn(id int) (a Alumn, err error) {
	var query = "SELECT a_id, name, year, occupation, phone, email, location, hobbies, talents, interests FROM alumni WHERE a_id = ?;"

	err = d.QueryRow(query, id).Scan(&a.Id, &a.Name, &a.Year, &a.Occupation, &a.Phone, &a.Email, &a.Location, &a.Hobbies, &a.Talents, &a.Interests)
	if err != nil {
		return
	}
	return
}


var currentId int

var alumni Alumni

// Give us some seed data
func init() {
	RepoCreateAlumn(Alumn{Name: "Dibo"})
	RepoCreateAlumn(Alumn{Name: "davis"})
}

func RepoFindAlumn(id int) Alumn {
	for _, a := range alumni {
		if a.Id == id {
			return a
		}
	}
	// return empty Alumni if not found
	return Alumn{}
}

// this is definitely not thread safe
func RepoCreateAlumn(t Alumn) Alumn {
	currentId += 1
	t.Id = currentId
	alumni = append(alumni, t)
	return t
}

func RepoDestroyAlumn(id int) error {
	for i, t := range alumni {
		if t.Id == id {
			alumni = append(alumni[:i], alumni[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Alumn with id of %d to delete", id)
}
