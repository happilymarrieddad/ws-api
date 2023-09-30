package config

import (
	"errors"
	"os"
)

type Config struct {
	OpenWeatherAPIKey string
}

func GetConfig() (*Config, error) {
	openWeatherAPIKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if len(openWeatherAPIKey) == 0 {
		return nil, errors.New("missing env var OPEN_WEATHER_API_KEY")
	}

	return &Config{
		OpenWeatherAPIKey: openWeatherAPIKey,
	}, nil
}
