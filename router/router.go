package router

import (
	"github.com/gorilla/mux"

	util "github.com/krishnakumarkp/simple/apputil"
	"github.com/krishnakumarkp/simple/mysqlstore"
)

func InitRoutes() *mux.Router {

	config := mysqlstore.Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}
	// Creates a Mysql DB instance
	dataStore, err := mysqlstore.New(config)

	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router = SetPersonRoutes(router, dataStore)
	return router
}
