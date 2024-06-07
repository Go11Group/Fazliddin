package hendler

import (
	"net/http"
)


func NewHandler(handler Handler) *http.Server {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/user/add", handler.CreateUser)
	mux.HandleFunc("/user/get", handler.Get)

	return &http.Server{Addr: ":8080", Handler: mux,}
}