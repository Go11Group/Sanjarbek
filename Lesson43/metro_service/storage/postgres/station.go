package postgres

import (
	"database/sql"
	"metro_service/models"

	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.CreateStation) error {

	id := uuid.NewString()
	_, err := s.Db.Exec(`INSERT INTO station(id, name) VALUES($1, $2)`, id, station.Name)

	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}

	err := s.Db.QueryRow(`SELECT name from station WHERE id = $1`, id).Scan(&station.Name)

	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (s *StationRepo) GetAll() ([]models.Station, error) {

	query := `SELECT id, name from station`

	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stations []models.Station

	for rows.Next() {
		station := models.Station{}

		err := rows.Scan(&station.Id, &station.Name)
		if err != nil {
			return nil, err
		}

		stations = append(stations, station)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}

func (s *StationRepo) UpdateStation(station models.Station) (*models.Station, error) {
	id := station.Id
	var checkStation models.Station
	err := s.Db.QueryRow(`
		SELECT id, name
		FROM courses
		WHERE id = $1`, station.Id).Scan(&checkStation.Id, &checkStation.Name)
	if err != nil {
		return nil, err
	}

	if station.Name == "" {
		station.Name = checkStation.Name
	}
	if station.Id == "" {
		station.Id = checkStation.Id
	}

	s.Db.Exec(`UPDATE station SET
			id = $1, name = $2
			where id = $3;`, station.Id, station.Name, id)

	return &station, err
}

func (s *StationRepo) DeleteStation(id string) error {
	query := `DELETE FROM station where id = $1`

	_, err := s.Db.Exec(query, id)
    if err != nil {
        return err
    }

	return nil
}
