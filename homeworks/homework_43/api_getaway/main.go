package main

import "homeworks/homework_43/api_getaway/handler"

func main() {
	panic(handler.Handler().ListenAndServe())
}