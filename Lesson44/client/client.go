package main

import (
	"context"
	"log"
	"time"
	pb "translate_service/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewTranslateServiceClient(conn)

	words := []string{"Salom", "dunyo", "qaleysan", "telephone", "naushnik"}

	req := pb.Massage{Words: words}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GiveTranslation(ctx, &req)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Translated words \n", r)
}