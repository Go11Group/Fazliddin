package main

import (
	"homeworks/homework_43/user/api"
	p "homeworks/homework_43/user/storage/postgres"
)


func main() {
	db, err := p.Conn()
	if err != nil{
		panic(err)
	}

	server := api.Handler(db)
	panic(server.ListenAndServe())
}