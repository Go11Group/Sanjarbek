package transport

import (
	"context"
	t "services/genproto/transportService"
	w "services/genproto/weatherService"
	postgres "services/storage/postgres"
)

type Server struct {
	t.UnimplementedTransportServiceServer
	w.UnimplementedWeatherServiceServer
	Transport *postgres.TransportRepo
	Weather   *postgres.WeatherRepo
}

func (s *Server) GetBusSchedule(ctx context.Context, in *t.Transport) (*t.Schedule, error) {
	resp, err := s.Transport.GetBusSchedule(in)
	return resp, err
}

func (s *Server) TrackBusLocation(ctx context.Context, in *t.Transport) (*t.Location, error) {
	resp, err := s.Transport.TrackBusLocation(in)
	return resp, err
}

func (s *Server) ReportTrafficJam(ctx context.Context, in *t.Transport) (*t.Traffic, error) {
	resp, err := s.Transport.ReportTrafficJam(in)
	return resp, err
}

func (s *Server) GetCurrentWeather(ctx context.Context, in *w.Place) (*w.CurrentWeatherResponse, error) {
	resp, err := s.Weather.GetCurrentWeather(in)
	return resp, err
}

func (s *Server) GetWeatherForecast(ctx context.Context, in *w.Forecast) (*w.WeatherForecastResponse, error) {
	resp, err := s.Weather.GetWeatherForecast(in)
	return resp, err
}

func (s *Server) ReportWeatherCondition(ctx context.Context, in *w.Place) (*w.WeatherConditionResponse, error) {
	resp, err := s.Weather.ReportWeatherCondition(in)
	return resp, err
}
