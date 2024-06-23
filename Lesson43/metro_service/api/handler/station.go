package handler

import (
	"fmt"
	"metro_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	stn := models.CreateStation{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetStationByID(ctx *gin.Context) {
	StationId := ctx.Param("id")

	station, err := h.Station.GetById(StationId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Station not found: %s", StationId)})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *handler) GetAllStations(c *gin.Context) {
	stations, err := h.Station.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, stations)
}

func (h *handler) UpdateStation(ctx *gin.Context) {
	stationId := ctx.Param("id")

	if stationId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid station ID"})
		return
	}

	var station models.Station
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateStation, err := h.Station.UpdateStation(station)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateStation)
}

func (h *handler) DeleteStation(ctx *gin.Context) {
	stationId := ctx.Param("id")

	err := h.Station.DeleteStation(stationId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete station: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "Station deleted succesfully"})
}
