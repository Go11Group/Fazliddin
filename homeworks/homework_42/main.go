package main

import (
	"fmt"
	"homeworks/homework_42/handler"
)

func main() {
	server := handler.Handler()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
