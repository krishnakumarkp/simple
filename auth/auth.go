package auth

import (
	"fmt"
	"net/http"
)

func checkUsernameAndPassword(username, password string) bool {
	return username == "abc" && password == "123"
}

//AuthorizeRequest Middleware validates requests.
func AuthorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)

		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
