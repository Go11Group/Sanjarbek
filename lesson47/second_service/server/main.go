package main

import (
	"log"
	"net"
	transportS "services/genproto/transportService"
	weatherS "services/genproto/weatherService"
	transport "services/service"
	postgres "services/storage/postgres"

	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	busRepo := postgres.NewTransportRepo(db)
	weatherRepo := postgres.NewWeatherRepo(db)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()

	transportS.RegisterTransportServiceServer(s, &transport.Server{Transport: busRepo})
	weatherS.RegisterWeatherServiceServer(s, &transport.Server{Weather: weatherRepo})

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
