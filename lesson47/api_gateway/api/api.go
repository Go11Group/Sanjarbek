package api

import (
	"api_gateway/api/handler"

	"api_gateway/genproto/weatherService"
	"api_gateway/genproto/transportService"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	weather := weatherService.NewWeatherServiceClient(conn)
	transport := transportService.NewTransportServiceClient(conn)

	handler := handler.NewHandler(weather, transport)

	weathe := router.Group("/weather")
	weathe.GET("/get/:location", handler.GetCurrentWeather)
	weathe.GET("/forecast/:location/:day", handler.GetWeatherForecast)
	weathe.GET("/report/:location", handler.ReportWeatherCondition)


	transpor := router.Group("/transport")
	transpor.GET("/get/:number", handler.GetBusSchedule)
	transpor.GET("/location/:number", handler.TrackBusLocation)
	transpor.GET("/report/:number", handler.ReportTrafficJam)

	return router
}
