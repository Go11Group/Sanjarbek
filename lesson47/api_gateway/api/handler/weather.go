package handler

import (
	"api_gateway/genproto/weatherService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	location := ctx.Param("location")
	req := &weatherService.Place{Name: location}

	resp, err := h.Weather.GetCurrentWeather(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWeatherForecast(ctx *gin.Context) {
	location := ctx.Param("location")
	day := ctx.Param("day")

	req := &weatherService.Forecast{Name: location, Day: day}

	resp, err := h.Weather.GetWeatherForecast(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportWeatherCondition(ctx  *gin.Context) {
	location := ctx.Param("location")
	req := &weatherService.Place{Name: location}

	resp, err := h.Weather.ReportWeatherCondition(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}