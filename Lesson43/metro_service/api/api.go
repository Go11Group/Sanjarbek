package api

import (
	"database/sql"
	"metro_service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(db)

	station := router.Group("/station")
	station.POST("/create", h.CreateStation)
	station.GET("/getById/:id", h.GetStationByID)
	station.GET("/GetAll", h.GetAllStations)
	station.PUT("/update/:id", h.UpdateStation)
	station.DELETE("/delete/:id", h.DeleteStation)

	card := router.Group("/card")
	card.POST("/create", h.CreateCard)
	card.GET("/getById/:id", h.GetCardByID)
	card.GET("/GetAll", h.GetAllCards)
	card.PUT("/update/:id", h.UpdateCard)
	card.DELETE("/delete/:id", h.DeleteCard)

	terminal := router.Group("/terminal")
	terminal.POST("/create", h.CreateTerminal)
	terminal.GET("/getById/:id", h.GetTerminalByID)
	terminal.GET("/GetAll", h.GetAllTerminals)
	terminal.PUT("/update/:id", h.UpdateTerminal)
	terminal.DELETE("/delete/:id", h.DeleteTerminal)

	transaction := router.Group("/transaction")
	transaction.POST("/create", h.CreateTransaction)
	transaction.GET("/getById/:id", h.GetTransactionByID)
	transaction.GET("/GetAll", h.GetAllTransactions)
	transaction.PUT("/update/:id", h.UpdateTransaction)
	transaction.DELETE("/delete/:id", h.DeleteTransaction)

	return router
}
