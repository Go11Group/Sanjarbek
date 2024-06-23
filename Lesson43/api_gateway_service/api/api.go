package api

import (
	"gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handle handler.Handler) *gin.Engine {
	r := gin.Default()

	user := r.Group("/users")
	user.POST("/create", handle.CreateUser)
	user.GET("/getAll", handle.GetUsers)
	user.GET("/getById/:id", handle.GetUser)
	user.PUT("/update/:id", handle.UpdateUser)
	user.DELETE("/delete/:id", handle.DeleteUser)



	cards := r.Group("/cards")
	cards.POST("/create", handle.CreateCard)
	cards.GET("/getById/:id", handle.GetCard)
	cards.GET("/getAll", handle.GetCards)
	cards.PUT("/update/:id", handle.UpdateCard)
	cards.DELETE("/delete/:id", handle.DeleteCard)
	


	station := r.Group("/stations")
	station.POST("/create", handle.CreateStation)
	station.GET("/getById/:id", handle.GetStation)
	station.GET("/getAll", handle.GetStations)
	station.PUT("/update/:id", handle.UpdateStation)
	station.DELETE("/delete/:id", handle.DeleteStation)



	terminal := r.Group("/terminals")
	terminal.POST("/", handle.CreateTerminal)
	terminal.GET("/getById/:id", handle.GetTerminal)
	terminal.GET("/getAll", handle.GetTerminals)
	terminal.PUT("/update/:id", handle.UpdateTerminal)
	terminal.DELETE("/delete/:id", handle.DeleteTerminal)



	transaction := r.Group("/transactions")
	transaction.POST("/", handle.CreateTransaction)
	transaction.GET("/getById/:id", handle.GetTransaction)
	transaction.GET("/getAll", handle.GetTransactions)
	transaction.PUT("/update/:id", handle.UpdateTransaction)
	transaction.DELETE("/delete/:id", handle.DeleteTransaction)



	r.GET("/check_balance", handle.CheckBalance)
	r.GET("/balance", handle.GetBalance)

	return r
}