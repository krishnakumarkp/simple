package router

import (
	"github.com/gorilla/mux"
	"github.com/krishnakumarkp/simple/auth"
	"github.com/krishnakumarkp/simple/controller"
	"github.com/krishnakumarkp/simple/mysqlstore"
)

func SetPersonRoutes(router *mux.Router, store mysqlstore.DataStore) *mux.Router {

	personStore := mysqlstore.PersonStore{
		store,
	}

	personController := controller.PersonController{
		Store: personStore,
	}

	personRouter := mux.NewRouter()

	personRouter.Handle("/person", controller.ResponseHandler(personController.ListAll)).Methods("GET")
	personRouter.Handle("/person", controller.ResponseHandler(personController.Add)).Methods("POST")

	router.PathPrefix("/person").Handler(auth.AuthorizeRequest(personRouter))
	return router
}
