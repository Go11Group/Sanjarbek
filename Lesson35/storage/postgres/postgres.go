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
    dbname   = "lessonn35"
    password = "1111"
)

func ConnectDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
        host, port, user, dbname, password)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        db.Close() // Close the connection if there's an error
        return nil, err
    }

    return db, nil
}