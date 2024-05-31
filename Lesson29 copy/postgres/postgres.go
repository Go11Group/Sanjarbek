package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM() (*gorm.DB) {
	db, _ := gorm.Open(postgres.Open("postgres://macbookpro:pass@localhost:5432/second?sslmode=disable"))

	return db
}