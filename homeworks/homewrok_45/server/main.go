package main

import (
	"context"
	"fmt"
	pb "homeworks/homewrok_45/genproto/generator"
	"net"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
}

var Books []pb.Book

func main() {
	lisetner, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := server{}

	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, &s)
	err = grpc.Serve(lisetner)
	if err != nil{
		panic(err)
	}
}

func (s *server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error){
	id := uuid.New().String()
	book := pb.Book{
		BookId: id,
		Title: in.Title,
		Author: in.Author,
		YearPublished: in.YearPublished,
	}
	Books = append(Books, book)
	fmt.Println(Books)
	return &pb.AddBookResponse{BookId: id}, nil
}

func (s *server) SearchBook(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error){
	bs := []*pb.Book{}
	strings.Contains(in.Query, "name")
	for i := 0; i < len(Books); i++{
		if Books[i].BookId == in.Query || Books[i].Title == in.Query || Books[i].Author == in.Query{
			bs = append(bs, &Books[i])
		}
	}
	return &pb.SearchBookResponse{Books: bs}, nil
}