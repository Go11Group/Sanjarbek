package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "macbookpro"
	password = "1111"
	dbname   = "lesson46"
)

func Connection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}