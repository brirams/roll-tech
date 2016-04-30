package main

type Alumn struct {
	Name       string  `json:"name"`
	Year       int     `json:"completed"`
	Occupation string  `json:"occupation"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	State      string  `json:"state"`
	Hobbies    string  `json:"hobbies"`
	Talents    string  `json:"talents"`
	Interests  string  `json:"interests"`
}

type Alumni []Alumn
