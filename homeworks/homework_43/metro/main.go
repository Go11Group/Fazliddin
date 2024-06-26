package main

import (
	"homeworks/homework_43/metro/api"
	"homeworks/homework_43/metro/storage/postres"
)

func main() {

	db, err := postres.Conn()
	if err != nil{
		panic(err)
	}

	panic(api.Handler(db).ListenAndServe())
}