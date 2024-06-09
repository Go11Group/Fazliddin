package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Handler(handler HandlerRepo) *http.Server{

	m := mux.NewRouter()

	m.HandleFunc("/user", handler.UserCreate)
	m.HandleFunc("/user/{id:[0-9]+}", handler.UserGetByID)
	m.HandleFunc("/user/get", handler.UserGet)
	m.HandleFunc("/user/update", handler.UserUpdate)
	m.HandleFunc("/user/delete/{id:[0-9]+}", handler.UserDelete)


	return &http.Server{Addr: ":8080", Handler: m}
}