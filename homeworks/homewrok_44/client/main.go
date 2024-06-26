package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Request struct {
	FirstName string
}

type Response struct{
	LastName string
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := pb.NewGeneratorClient(conn)
	req := Request{
		FirstName: "Fazliddin",
	}

	res, err := gen.GetLastName(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
