package main

import (
	"homeworks/homework_46/generate/transport"
	service "homeworks/homework_46/server_2/metods"
	"homeworks/homework_46/storage"
	postgers "homeworks/homework_46/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {

	db, err := postgers.Conn()
	if err != nil{
		panic(err)
	}

	t := storage.NewStorageRepo(db)
	s := service.TranpostRepo{St: t}

	listner, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}

	grpc := grpc.NewServer()
	transport.RegisterTranportServiceServer(grpc, s)

	err = grpc.Serve(listner)
	if err != nil {
		panic(err)
	}
}
