package mysqlstore

import (
	"github.com/krishnakumarkp/simple/domain"
	"github.com/pkg/errors"
)

type PersonStore struct {
	Store DataStore
}

func (ps PersonStore) GetAll() ([]domain.Person, error) {
	var people []domain.Person
	var person domain.Person

	sqlStatement := `SELECT id, firstname, lastname, firstparent_id, secondparent_id FROM person ORDER BY id DESC;`
	rows, err := ps.Store.Db.Query(sqlStatement)
	if err != nil {
		return people, errors.New("could not retrive persons")
	}

	for rows.Next() {
		person = domain.Person{}
		err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.FirstParentID, &person.SecondParentID)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

func (ps PersonStore) Create(p domain.Person) (domain.Person, error) {
	sqlStatement := "insert into person(firstname, lastname, firstparent_id, secondparent_id) values (?,?,?,?)"
	insStatment, err := ps.Store.Db.Prepare(sqlStatement)
	if err != nil {
		return p, err
	}
	result, ierr := insStatment.Exec(p.FirstName, p.LastName, p.FirstParentID, p.SecondParentID)
	if ierr != nil {
		return p, errors.Wrap(ierr, "Error while inserting person")
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return p, nil
}
