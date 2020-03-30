package domain

import "github.com/gobuffalo/pop/nulls"

type Person struct {
	ID             int
	FirstName      string      `json:"name"`
	LastName       string      `json:"lastname"`
	FirstParentID  nulls.Int64 `json:"firstparent, string"`
	SecondParentID nulls.Int64 `json:"secondparent, string"`
}

type PersonStore interface {
	Create(Person) (Person, error)
	//Update(string, Person) error
	//Delete(int) error
	//GetById(int) (Person, error)
	GetAll() ([]Person, error)
}
