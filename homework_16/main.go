package main

import (
	"fmt"
	"model/postgers"
	"model/problems"
	"model/server"
	"model/user"
)

func main() {
	db, err := postgers.Conn()
	if err != nil{
		fmt.Println(err)
		return
	}

	u := user.NewUserRepo(db)
	p := problems.NewProblemRepo(db)

	hendler := server.HandlerRepo{u, p}

	server := server.Handler(hendler)
	
	err = server.ListenAndServe()
	if err != nil{
		fmt.Println("Error 2: ", err)
		return
	}
}
