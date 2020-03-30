package main

import (
	"log"
	"net/http"

	util "github.com/krishnakumarkp/simple/apputil"
	"github.com/krishnakumarkp/simple/router"
	"github.com/spf13/viper"
)

func main() {

	router := router.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))

}

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	util.AppConfig.DBHost = viper.GetString("mysql.Host")
	util.AppConfig.DBPort = viper.GetString("mysql.Port")
	util.AppConfig.DBUser = viper.GetString("mysql.User")
	util.AppConfig.DBPassword = viper.GetString("mysql.Password")
	util.AppConfig.Database = viper.GetString("mysql.Database")
}
