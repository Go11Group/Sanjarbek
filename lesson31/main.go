package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-faker/faker/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "macbookpro"
	dbname   = "second"
	password = "root"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Failed to open connection:", err)
	}
	defer db.Close()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomCarMake := func() string {
		makes := []string{"Toyota", "Honda", "Ford", "BMW", "Mercedes"}
		return makes[rng.Intn(len(makes))]
	}

	randomCarModel := func() string {
		models := []string{"Corolla", "Civic", "Mustang", "X5", "E-Class"}
		return models[rng.Intn(len(models))]
	}

	randomYear := func() int {
		return rng.Intn(31) + 1994
	}

	randomPrice := func() int {
		return rng.Intn(50000) + 10000
	}

	randomOwner := func() string {
		owners := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
		return owners[rng.Intn(len(owners))]
	}

	query := `insert into car(id, make, model, year, price, owner) values ($1, $2, $3, $4, $5, $6)`
	for i := 0; i < 1000000; i++ {
		id := faker.UUIDHyphenated()
		_, err := db.Exec(query, id, randomCarMake(), randomCarModel(), randomYear(), randomPrice(), randomOwner())
		if err != nil {
			log.Fatal("Error inserting record:", err)
		}
		if i%1000 == 0 {
			fmt.Println("Inserted", i, "records")
		}
	}

	min := math.MaxFloat64
	max := float64(-15152151212)

	for i := 0; i < 1000; i++ {
		rows, err := db.Query(`explain (analyse)
		select * from car
		where price = 999999 and 
		id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab'
		and make = 'Toyota'`)
		if err != nil {
			log.Fatal("Error executing query:", err)
		}
		defer rows.Close()

		row := ""
		for rows.Next() {
			err = rows.Scan(&row)
			if err != nil {
				log.Fatal("Error scanning row:", err)
			}
		}
		timeTook, err := strconv.ParseFloat(row[16:21], 64)
		if err != nil {
			log.Fatal("Error parsing float:", err)
		}
		if min > timeTook {
			min = timeTook
		}
		if max < timeTook {
			max = timeTook
		}
	}
	fmt.Println("Minimum query time:", min)
	fmt.Println("Maximum query time:", max)
}
