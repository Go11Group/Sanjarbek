package main

import (
	"fmt"
	"module/model"
	postgres "module/storage"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Users
	users := []model.Users{
		{Id: 1, UserName: "johndoe", Email: "johndoe@example.com", Password:"johnspassword"},
		{Id: 2, UserName: "janedoe", Email: "janedoe@example.com", Password: "janespassword"},
		{Id: 3, UserName: "alice", Email: "alice@example.com", Password: "alicepassword"},
		{Id: 4, UserName: "bobsmith", Email: "bobsmith@example.com", Password: "bobspassword"},
		{Id: 5, UserName: "charliebrown", Email: "charliebrown@example.com", Password: "charliespassword"},
	}
	// Products
	products := []model.Products{
		{Id: 1, Name: "Laptop", Description: "A high-performance laptop for everyday use", Price: 999.99, Stock_quantity: 50},
		{Id: 2, Name: "Smartphone", Description: "A latest model smartphone with advanced features", Price: 699.99, Stock_quantity: 200},
		{Id: 3, Name: "Headphones", Description: "Noise-cancelling over-ear headphones", Price: 199.99, Stock_quantity: 150},
		{Id: 4, Name: "Smartwatch", Description: "A smartwatch with fitness tracking capabilities", Price: 249.99, Stock_quantity: 120},
		{Id: 5, Name: "Tablet", Description: "A tablet with a high-resolution display and large storage", Price: 399.99, Stock_quantity: 80},
	}

	// UserProducts
	userProducts := []model.UserProduct{
		{Id: 1, UserId: 1, ProductId: 1},
		{Id: 2, UserId: 2, ProductId: 2},
		{Id: 3, UserId: 3, ProductId: 3},
		{Id: 4, UserId: 4, ProductId: 4},
		{Id: 5, UserId: 5, ProductId: 5},
		{Id: 6, UserId: 1, ProductId: 3},
		{Id: 7, UserId: 2, ProductId: 5},
		{Id: 8, UserId: 3, ProductId: 4},
		{Id: 9, UserId: 4, ProductId: 1},
		{Id: 10, UserId: 5, ProductId: 2},
		
	}

	userRepo := postgres.NewUserRepo(db)
	productRepo := postgres.NewProductRepo(db)
	userProductRepo := postgres.NewUserProductRepo(db)

	for _, v := range users {
		err = userRepo.CreateUser(v)
		if err != nil {
			panic(err)
		}
	}

	
	for _, v := range products {
		err = productRepo.CreateProduct(v)
		if err != nil {
			panic(err)
		}
	}

	for _, v := range userProducts {
		err = userProductRepo.CreateUserProduct(v)
		if err != nil {
			panic(err)
		}
	}
	
	user := users[0]
	product := products[0]
	userProduct := userProducts[0]


	users1, err := userRepo.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Users: %+v\n", users)

	user.Id = users1[0].Id
	user.UserName = "NewUser"
	err = userRepo.UpdateUser(user)
	if err != nil {
		panic(err)
	}

	products1, err := productRepo.GetProducts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Products: %+v\n", products1)

	product.Id = products[0].Id
	product.Name = "NewProduct"
	err = productRepo.UpdateProduct(product)
	if err != nil {
		panic(err)
	}

	userProducts1, err := userProductRepo.GetUserProducts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("UserProducts: %+v\n", userProducts1)

	userProduct.Id = userProducts[0].Id
	userProduct.ProductId = product.Id
	err = userProductRepo.UpdateUserProduct(userProduct)
	if err != nil {
		panic(err)
	}

	err = userProductRepo.DeleteUserProduct(userProduct.Id)
	if err != nil {
		panic(err)
	}

	err = productRepo.DeleteProduct(product.Id)
	if err != nil {
		panic(err)
	}

	err = userRepo.DeleteUser(user.Id)
	if err != nil {
		panic(err)
	}
}