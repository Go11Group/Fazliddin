package main

import (
	"context"
	"net"
	"strings"

	pb "homeworks/homework_44/genproto/getLastName"
	"google.golang.org/grpc"
)

type nameserver struct {
	pb.UnimplementedGeneratorServer
}

func main() {
	listen, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}

	s := nameserver{}

	grpcServer := grpc.NewServer()
	pb.RegisterGeneratorServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}

func (s *nameserver) GetLastName(ctx context.Context, rq *pb.Request) (*pb.Response, error) {
	names := []string{
		"Abbos Qambarov",
		"Azizbek Qobulov",
		"Bekzod Qo'chqarov",
		"Diyorbek Nematov Dadajon o'g'li",
		"Faxriddin Raximberdiyev Farxodjon o'g'li",
		"Fazliddin Xayrullayev",
		"Hamidjon Nuriddinov",
		"Hamidulloh Hamidullayev",
		"Ibrohim Umarov Fazliddin o'g'li",
		"Jamshidbek Hatamov Erkin o'g'li",
		"Javohir Abdusamatov",
		"Muhammadaziz Yoqubov",
		"Muhammadjon Ko'palov",
		"Nurmuhammad",
		"Ozodjon A'zamjonov",
		"Sanjarbek Abduraxmonov",
		"Yusupov Bobur",
		"Firdavs",
		"Ozodbek",
		"Abdulaziz Xoliqulov",
		"Intizor opa",
	}

	var lastName string
	for _, name := range names {
		if strings.Contains(name, rq.FirstName) {
			lastName = name
			break
		}
	}
	if lastName == "" {
		return nil, status.Errorf(codes.NotFound, "last name not found for first name: %s", rq.FirstName)
	}
	return &pb.Response{LastName: lastName}, nil
}
