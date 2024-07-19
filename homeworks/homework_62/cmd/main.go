package main

import (
	"homework_62/api"
	"homework_62/api/handler"
	"homework_62/service"
	"homework_62/storage/redis_1"
	"log"
)

func main() {
	client := redis_1.NewRedisClient()
	redis := redis_1.NewBookRepo(client)
	serv := &service.Service{Dtb: redis}

	hd := handler.Handler{Service: serv}

	router := api.Router(hd)

	log.Fatalln(router.Run(":8080"))
}
