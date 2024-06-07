package postgres

import (
	"database/sql"
	"module/model"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (b *BookRepo) Get(id string) (model.Book, error) {
	var book model.Book
	err := b.db.QueryRow("SELECT * FROM book WHERE id = $1", id).Scan(
		&book.Id, &book.Name, &book.Page, &book.AuthorName)

	return book, err
}
