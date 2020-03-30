package controller

import (
	"encoding/json"
	"fmt"
	util "github.com/krishnakumarkp/simple/apputil"
	"net/http"
)

type response struct {
	Data interface{} `json:"data"`
}

func ResponseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)

		if err != nil {
			data = err.Error()
		}

		w.Header().Set("Content-Type", "application/json")
		util.EnableCors(&w)
		w.WriteHeader(status)

		if data != nil {
			err = json.NewEncoder(w).Encode(response{Data: data})
			if err != nil {
				fmt.Printf("Error from Handler: %s\n", err.Error())
			}
		}
	})
}
