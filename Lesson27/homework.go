package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

func main() {
// db.Query Goning ma'lumotlar bazasi/sql paketida satrlarni qaytaruvchi so'rovni bajarish uchun ishlatiladi,
// odatda SELECT iborasi ma'lumotlar bazasidan ma'lumotlarni olish uchun ishlatiladi.
// *sql.Rows ob'ektini qaytaradi, undan natijalar to'plamini takrorlash uchun foydalanish mumkin.

    db, err := sql.Open("postgres", "user=macbookpro dbname=first sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    query := "SELECT id, name, age FROM student WHERE age >= $1"

    rows, err := db.Query(query, 19)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id string
        var name string
        var age int
        err = rows.Scan(&id, &name, &age)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %s, Name: %s, Age: %d\n", id, name, age)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}

