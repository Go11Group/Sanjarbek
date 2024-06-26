package main

import (
	"context"
	"fmt"
	pb "library_service/getproto/libraryService"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	gen := pb.NewLibraryServiceClient(conn)
	req := pb.AddBookRequest{
		Title: "Intersting",
		Author: "Qodiriy",
		YearPublished: 2000,
	}

	resp, err := gen.AddBook(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}