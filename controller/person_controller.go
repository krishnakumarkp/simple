package controller

import (
	"encoding/json"
	"net/http"

	"github.com/krishnakumarkp/simple/domain"
	"github.com/pkg/errors"
)

type PersonController struct {
	Store domain.PersonStore
}

func (pc PersonController) ListAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	people, err := pc.Store.GetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Unable to fetch persons")
	}
	_, jerr := json.Marshal(people)
	if jerr != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(jerr, "Unable to json encode Person")
	}
	return people, http.StatusOK, nil
}

func (pc PersonController) Add(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	person := domain.Person{}

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(err, "Unable to decode JSON request body")
	}

	newPerson, err := pc.Store.Create(person)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Error on inserting Person")
	}
	return newPerson, http.StatusCreated, nil
}
