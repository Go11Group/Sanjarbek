package handler

import (
	"api_gateway/genproto/transportService"
	"api_gateway/genproto/weatherService"
)

type Handler struct {
	Weather   weatherService.WeatherServiceClient
	Transport transportService.TransportServiceClient
}

func NewHandler(Weather weatherService.WeatherServiceClient, Transport transportService.TransportServiceClient) *Handler {
	return &Handler{Weather: Weather, Transport: Transport}
}
