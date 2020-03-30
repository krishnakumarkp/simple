package apputil

import "net/http"

type Configuration struct {
	DBHost, DBPort, DBUser, DBPassword, Database string
}

var AppConfig Configuration

func init() {
	AppConfig = Configuration{}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
