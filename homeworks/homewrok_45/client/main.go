package main

import (
	"context"
	"fmt"
	pb "homeworks/homewrok_45/genproto/generator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		panic(err)
	}
	defer con.Close()

	gen := pb.NewLibraryServiceClient(con)
	addReq := &pb.AddBookRequest{
		Title: "nmadur",
		Author: "Robert Kiosaki",
		YearPublished: 2005,
	}
	
	addRes, err := gen.AddBook(context.Background(), addReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(addRes)
	
	serchReq := pb.SearchBookRequest{Query: "nmadur"} 
	serRes, err := gen.SearchBook(context.Background(), &serchReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(serRes)
}