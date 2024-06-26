package main

import (
	"context"
	"fmt"
	"net"

	pb "library_service/getproto/libraryService"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
}

var titles = []string{"Interesting", "Scientific", "Unreal"}
var authors = []string{"Qodiriy", "Jekson", "Navoiy"}
var YearPublishes = []int{2000, 2010, 1400}

var BookIds = []string{"dsagds1", "gdfigbfi2", "dsghuerhgh3"}

func (s *server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	d := 0
	title := ""
	author := ""
	yearPublished := 0
	for i, v := range titles {
		if v == in.Title {
			title = v
			author = authors[i]
			yearPublished = YearPublishes[i]
			d = i
		}
	}
	fmt.Println(title, author, yearPublished)
	return &pb.AddBookResponse{BookId: BookIds[d]}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := server{}
	grpc := grpc.NewServer()

	pb.RegisterLibraryServiceServer(grpc, &s)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}
