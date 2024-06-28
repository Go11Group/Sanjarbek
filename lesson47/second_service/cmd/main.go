package main

import (
	"services/storage/postgres"
	"log"
	"net"

	t "services/genproto/transportService"

	"google.golang.org/grpc"

	"github.com/Go11Group/at_lesson/lesson47/second_service/pkg/db"
	s "github.com/Go11Group/at_lesson/lesson47/second_service/service"
	_ "github.com/lib/pq"
)

func main() {

	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()

	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewWeatherStorage(db)
	service := s.WeatherService{Storage: *storage}
	tr := s.TransportService{}

	p.RegisterWeatherServiceServer(server, &service)
	t.RegisterTransportServiceServer(server, &tr)

	log.Println("server is running on :7070 ...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}