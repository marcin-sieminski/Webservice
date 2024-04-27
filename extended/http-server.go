package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "5000"
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

func main() {
	http.HandleFunc("/", basicAuth(createResponse, "Please enter your username and password"))
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}

func basicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}

func createResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Extended microservice")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login")
}
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout")
}
