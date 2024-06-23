package postgres

import (
	"database/sql"
	"metro_service/models"

	"github.com/google/uuid"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (t *TerminalRepo) Create(terminal *models.Terminal) error {
	id := uuid.NewString()
	_, err := t.Db.Exec(`
        INSERT INTO terminal(id, station_id)
        VALUES($1, $2)`,
		id, terminal.StationId)
	return err
}

func (t *TerminalRepo) GetById(id string) (*models.Terminal, error) {
	terminal := &models.Terminal{Id: id}
	err := t.Db.QueryRow(`
        SELECT station_id
        FROM terminal
        WHERE id = $1`, id).Scan(&terminal.StationId)
	if err != nil {
		return nil, err
	}
	return terminal, nil
}

func (t *TerminalRepo) GetAll() ([]models.Terminal, error) {
	rows, err := t.Db.Query(`
        SELECT id, station_id
        FROM terminal`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		terminal := models.Terminal{}
		err := rows.Scan(&terminal.Id, &terminal.StationId)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return terminals, nil
}

func (t *TerminalRepo) Update(terminal models.Terminal) (*models.Terminal, error) {
	id := terminal.Id
	var checkTerminal models.Terminal
	err := t.Db.QueryRow(`
		SELECT id, station_id
		FROM terminal
		WHERE id = $1`, terminal.Id).Scan(&checkTerminal.Id, &checkTerminal.StationId)
	if err != nil {
		return nil, err
	}

	if terminal.Id == "" {
		terminal.Id = checkTerminal.Id
	}
	if terminal.StationId == "" {
		terminal.StationId = checkTerminal.StationId
	}

	_, err = t.Db.Exec(`
        UPDATE terminal
        SET station_id = $1, id = $2
        WHERE id = $3`,
		terminal.Id, terminal.StationId, id)
	if err != nil {
		return nil, err
	}
	
	return &terminal, nil
}

func (t *TerminalRepo) Delete(id string) error {
	_, err := t.Db.Exec(`DELETE FROM terminal WHERE id = $1`, id)
	return err
}
