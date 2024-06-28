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

func (t *TransportRepo) GetBusSchedule(tr *pb.Transport) (*pb.Schedule, error) {
	sch := pb.Schedule{}

	query := `select stations from transport where number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&sch.Stations)
	if err != nil {
		log.Println("Bus not found")
		return nil, err
	}
	return &sch, nil
}

func (t *TransportRepo) TrackBusLocation(tr *pb.Transport) (*pb.Location, error) {
	loc := pb.Location{}

	query := `select current_station from transport where number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&loc.Station)
	if err != nil {
		log.Println("Bus's location not found")
		return nil, err
	}
	return &loc, nil
}

func (t *TransportRepo) ReportTrafficJam(tr *pb.Transport) (*pb.Traffic, error) {
	tra := pb.Traffic{}

	query := `select is_full from transport where number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&tra.IsFull)
	if err != nil {
		log.Println("Bus not found")
		return nil, err
	}
	return &tra, nil
}
