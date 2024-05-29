package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Brand struct to represent a brand entity
type Brand struct {
	ID   string
	Name string
	Year int
}

// Car struct to represent a car entity
type Car struct {
	ID      string
	BrandID string
	Name    string
	Year    int
	Price   float64
}

func main() {
	// Database connection string
	connStr := "user=yourusername dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert data into the brand table
	var brand Brand
	insertBrandQuery := `
		INSERT INTO brand (name, year)
		VALUES ($1, $2)
		RETURNING id, name, year;
	`
	err = db.QueryRow(insertBrandQuery, "Toyota", 1937).Scan(&brand.ID, &brand.Name, &brand.Year)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted brand: %+v\n", brand)

	// Insert data into the car table
	var car Car
	insertCarQuery := `
		INSERT INTO car (brand_id, name, year, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, brand_id, name, year, price;
	`
	err = db.QueryRow(insertCarQuery, brand.ID, "Camry", 2020, 24000).Scan(&car.ID, &car.BrandID, &car.Name, &car.Year, &car.Price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted car: %+v\n", car)

	// Fetch all brands using db.Query
	fetchBrandsQuery := `SELECT id, name, year FROM brand`
	rows, err := db.Query(fetchBrandsQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Brands:")
	for rows.Next() {
		var b Brand
		err := rows.Scan(&b.ID, &b.Name, &b.Year)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Name: %s, Year: %d\n", b.ID, b.Name, b.Year)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Fetch a specific car using db.QueryRow
	var fetchedCar Car
	fetchCarQuery := `SELECT id, brand_id, name, year, price FROM car WHERE id = $1`
	err = db.QueryRow(fetchCarQuery, car.ID).Scan(&fetchedCar.ID, &fetchedCar.BrandID, &fetchedCar.Name, &fetchedCar.Year, &fetchedCar.Price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Car Details: %+v\n", fetchedCar)
}

