package postgres

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	pb "services/genproto/transportService"
)

type TransportRepo struct {
	db *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{db: db}
}

func (t *TransportRepo) GetBusSchedule(tr *pb.Transport) (*pb.Schedule, error) {
	sch := pb.Schedule{}
	var stations pq.StringArray

	query := `SELECT stations FROM transport WHERE number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&stations)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No bus schedule found")
			return nil, err
		}
		log.Println("Error scanning bus schedule:", err)
		return nil, err
	}

	// Convert pq.StringArray to []string
	sch.Stations = make([]string, len(stations))
	for i, station := range stations {
		sch.Stations[i] = station
	}

	return &sch, nil
}

func (t *TransportRepo) TrackBusLocation(tr *pb.Transport) (*pb.Location, error) {
	loc := pb.Location{}

	query := `SELECT current_station FROM transport WHERE number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&loc.Station)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Bus's location not found")
			return nil, err
		}
		log.Println("Error scanning bus location:", err)
		return nil, err
	}

	return &loc, nil
}

func (t *TransportRepo) ReportTrafficJam(tr *pb.Transport) (*pb.Traffic, error) {
	tra := pb.Traffic{}

	query := `SELECT is_full FROM transport WHERE number = $1`

	err := t.db.QueryRow(query, tr.Number).Scan(&tra.IsFull)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Bus not found")
			return nil, err
		}
		log.Println("Error scanning traffic jam report:", err)
		return nil, err
	}

	return &tra, nil
}
