package hendler

import (
	"net/http"
)


func NewHandler(handler Handler) *http.Server {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/user/add", handler.CreateUser)
	mux.HandleFunc("/user/get", handler.Get)
	mux.HandleFunc("/user/id", handler.GetById)

	return &http.Server{Addr: ":8080", Handler: mux,}
}