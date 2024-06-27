package postgres

import (
	"database/sql"
	"log"
	pb "services/genproto/transportService"
)

type TransportRepo struct {
	db *sql.DB
}

func NewTransportRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{db: db}
}

func (t *TransportRepo) GetBusSchedule(tr *pb.Transport) (*pb.Schedule, error){
	sch := pb.Schedule{}

	query := `select`
}
