package postgres

import (
	"database/sql"
	"log"
	pb "services/genproto/weatherService"
)

type WeatherRepo struct {
	db *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{db: db}
}

func (w *WeatherRepo) CreateWeather(weather pb.WeatherConditionResponse) error {
	_, err := w.db.Exec(`INSERT INTO weather(name, temperature, humidity, windSpeed, condition) VALUES($1)`,
		weather.Place, weather.Temperature, weather.Humidity, weather.WindSpeed, weather.Condition)
		if err != nil {
			return err
		}
		return nil
}

func (w *WeatherRepo) GetCurrentWeather(p *pb.Place) (*pb.CurrentWeatherResponse, error) {
	currentW := pb.CurrentWeatherResponse{}

	query := `select temperature, humidity, wind_speed from weather where name = $1 and date = CURRENT_DATE`
	err := w.db.QueryRow(query, p.Name).Scan(&currentW.Temperature, &currentW.Humidity, &currentW.WindSpeed)
	if err != nil {
		log.Println("Place not found")
		return nil, err
	}

	return &currentW, nil
}

func (w *WeatherRepo) GetWeatherForecast(p *pb.Forecast) (*pb.WeatherForecastResponse, error) {
	forecastW := pb.WeatherForecastResponse{}

	query := `select temperature, humidity, wind_speed from weather where name = $1 and date = $2`
	err := w.db.QueryRow(query, p.Name, p.Day).Scan(&forecastW.Temperature, &forecastW.Humidity, &forecastW.WindSpeed)
	if err != nil {
		log.Println("Place or day not found")
		return nil, err
	}

	return &forecastW, nil
}

func (w *WeatherRepo) WeatherConditionResponse(p *pb.Place) (*pb.WeatherConditionResponse, error) {
	conditionW := pb.WeatherConditionResponse{}

	query := `select temperature, humidity, wind_speed, condition from weather_conditions where name = $1 and date = CURRENT_DATE`
	err := w.db.QueryRow(query, p.Name).Scan(&conditionW.Temperature, &conditionW.Humidity, &conditionW.WindSpeed, &conditionW.Condition)

	if err != nil {
		log.Println("Place not found")
		return nil, err
	}

	return &conditionW, nil
}