DB_URL = postgres://macbookpro:1111@localhost:5432/user_service?sslmode=disable

run:
	go run main.go

migrate_file:
	migrate create -ext sql -dir migrations -seq user_service

migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down