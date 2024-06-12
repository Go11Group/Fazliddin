package main

import (
	"fmt"
	"model/hendler"
	"model/postgres"
	"model/product"
	"model/user"
)

func main() {
	db, err := postgres.Conn()
	if err != nil{
		panic(err)
	}

	u := user.NewUserRepo(db)
	p := product.NewProductRepo(db)

	server := hendler.NewHandler(hendler.Handler{u, p})
	err = server.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}