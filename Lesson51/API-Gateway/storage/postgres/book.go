package postgres

import (
	"api-service/models"
	"database/sql"
)

type BookRepo struct {
	DB *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{DB: db}
}



func (b *BookRepo) CreateBook(req models.Book) error {
	query := "INSERT INTO books (book_id, name, author, published, pages) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id string
	err := b.DB.QueryRow(query, req.BookID, req.Name, req.Author, req.Published, req.Pages).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) GetBook(req models.Book) error {
	query := "Select book_id, name, author, published, pages from books  WHERE book_id = $1"
	var id string
	err := b.DB.QueryRow(query, req.BookID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}