package main

import (
	"context"
	"fmt"
	"log"
	"services/genproto/transportService"
	t "services/genproto/transportService"
	"services/genproto/weatherService"
	w "services/genproto/weatherService"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}

	defer conn.Close()

	w := w.NewWeatherServiceClient(conn)
	t := t.NewTransportServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()
	// 1-chi function transport
	rGetBus := transportService.Transport{Number: 12}
	resGetBusSch, err := t.GetBusSchedule(ctx, &rGetBus)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resGetBusSch)

	// 2-chi function transport
	resTrackBus, err := t.TrackBusLocation(ctx, &rGetBus)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resTrackBus)

	// 3-chi function transport
	resReportBus, err := t.ReportTrafficJam(ctx, &rGetBus)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resReportBus)


	// 1-chi function weather
	rGetWeather := weatherService.Place{Name: "Tashkent"}
	rGetCurWeather, err := w.GetCurrentWeather(ctx, &rGetWeather)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rGetCurWeather)

	// 2-chi function weather
	rGetForecast := weatherService.Forecast{Name: "Tashkent", Day: "2024-06-27"}
	rGetForecastWeather, err := w.GetWeatherForecast(ctx, &rGetForecast)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rGetForecastWeather)

	// 3-chi function weather
	rGetCondition, err := w.ReportWeatherCondition(ctx, &rGetWeather)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rGetCondition)
}