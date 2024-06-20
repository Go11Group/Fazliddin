package main

import (
	"fmt"
	"homework_41/server_client/server/handler"
	"homework_41/server_client/server/postgres"
)

func main() {
	db, err := postgres.Conn()
	if err != nil{
		fmt.Println("ERROR in conn", err)
	}

	u := postgres.NewUserRepo(db)

	h := handler.HandlerRepo{u}

	server := handler.Server(h)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}