package postgres

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "macbookpro"
    dbname   = "examine"
    password = "1111"
)

func ConnectDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
        host, port, user, dbname, password)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}