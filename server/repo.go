package main

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	_ "google.golang.org/appengine/cloudsql"
)

type Db struct {
	*sql.DB
}

type AlumniRepo interface {
	GetAlumn(int) (Alumn, error)
	InsertAlumn(Alumn) (Alumn, error)
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

func (d Db) InsertAlumn(alumn Alumn) (inserted Alumn, err error) {
	stmt, err := d.Prepare("INSERT INTO alumni (name, year, occupation, phone, email, location, hobbies, talents, interests) values (?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return
	}
	res, err := stmt.Exec(alumn.Name, alumn.Year, alumn.Occupation, alumn.Phone, alumn.Email, alumn.Location, alumn.Hobbies, alumn.Talents, alumn.Interests)
	if err != nil {
		return
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return
	}

	inserted = Alumn{int(lastId), alumn.Name, alumn.Year, alumn.Occupation, alumn.Phone, alumn.Email, alumn.Location, alumn.Hobbies, alumn.Talents, alumn.Interests}

	return
}
