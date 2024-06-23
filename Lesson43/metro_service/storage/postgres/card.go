package postgres

import (
	"database/sql"
	"fmt"
	"metro_service/models"

	"github.com/google/uuid"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (s *CardRepo) Create(card *models.CreateCard) error {

	id := uuid.NewString()
	_, err := s.Db.Exec(`INSERT INTO card(id, number, user_id) VALUES($1, $2)`, id, card.Number, card.UserId)

	return err
}

func (s *CardRepo) GetById(id string) (*models.Card, error) {
	var card = models.Card{Id: id}

	err := s.Db.QueryRow(`SELECT number, user_id from card WHERE id = $1`, id).Scan(&card.Number, &card.UserId)

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (s *CardRepo) GetAll() ([]models.Card, error) {

	query := `SELECT id, number, user_id from card`

	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cards []models.Card

	for rows.Next() {
		card := models.Card{}

		err := rows.Scan(&card.Id, &card.Number, &card.UserId)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (s *CardRepo) UpdateCard(card models.Card) (*models.Card, error) {
	var checkCard models.Card
	err := s.Db.QueryRow(`
		SELECT id, number, user_id
		FROM card
		WHERE id = $1`, card.Id).Scan(&checkCard.Id, &checkCard.Number, &checkCard.UserId)
	fmt.Println(err, checkCard)
	if err != nil {
		return nil, err
	}

	if card.Number == "" {
		card.Number = checkCard.Number
	}
	if card.Id == "" {
		card.Id = checkCard.Id
	}
	if card.UserId == "" {
		card.UserId = checkCard.UserId
	}

	s.Db.Exec(`UPDATE card SET
			id = $1, number = $2, user_id;`, card.Id, card.Number, card.UserId)

	return &card, err
}

func (s *CardRepo) DeleteCard(id string) error {
	query := `DELETE FROM card where id = $1`

	_, err := s.Db.Exec(query, id)
    if err != nil {
        return err
    }

	return nil
}