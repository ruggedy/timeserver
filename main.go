package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	http.ListenAndServe(WEBSERVERPORT, nil)
}
