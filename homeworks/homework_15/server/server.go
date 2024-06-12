package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Handler(handler HandlerRepo) *http.Server{

	m := mux.NewRouter()

	user := m.PathPrefix("/user").Subrouter()
	user.HandleFunc("/create", handler.UserCreate)
	user.HandleFunc("", handler.UserGetByID).Methods("GET")
	user.HandleFunc("/get", handler.UserGet)
	user.HandleFunc("/update", handler.UserUpdate)
	user.HandleFunc("/delete/{id:[0-9]+}", handler.UserDelete)

	problems := m.PathPrefix("/problems").Subrouter()
	problems.HandleFunc("/create", handler.ProblemCreate)
	problems.HandleFunc("/{id:[0-9]+}", handler.ProblemGetByID).Methods("GET")
	problems.HandleFunc("/get", handler.ProblemGet)
	problems.HandleFunc("/update", handler.ProblemUpdate)
	problems.HandleFunc("/delete/{id:[0-9]+}", handler.ProblemDelete)

	

	return &http.Server{Addr: ":8080", Handler: m}
}