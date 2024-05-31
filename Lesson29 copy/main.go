package main

import (
	"fmt"
	"module/model"
	"module/postgres"
)



func main() {
	users := model.User{}

	db := postgres.ConnectGORM()

	us := postgres.NewUserRepo(db)

	// us.CreateTable(users)

	db.Create(&model.User{FirstName: "Sanjarbek", LastName: "Abduraxmonov", Email: "sanjarbek@gmail.com", Password: "Sanjarabd", Age: 17, Field: "GOLANG", Gender: "Male", IsEmployee: false})

	new := us.Read()
	fmt.Println(new)

	us.Delete(users)
}