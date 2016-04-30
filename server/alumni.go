package main

type Alumn struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Year       int     `json:"year"`
	Occupation string  `json:"occupation"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Location   string  `json:"location"`
	Hobbies    string  `json:"hobbies"`
	Talents    string  `json:"talents"`
	Interests  string  `json:"interests"`
}

type Alumni []Alumn
