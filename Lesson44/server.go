package main

import (
	"context"
	"log"
	pb "translate_service/server"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTranslateServiceServer
}

var allwords = map[string]string{
	"salom":"Hello",
	"dunyo":"World",
	"qaleysan":"What's up",
	"kompyuter":"Laptop",
	"telephone":"cellphone",
	"naushnik":"earphone"}

func (s *server) GiveTranslation(ctx context.Context, req *pb.Massage) (*pb.Answer, error) {
	res := req.Words
	for i, v := range res {
		res[i] = allwords[strings.ToLower(v)]
	}
	return &pb.Answer{Words: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslateServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}