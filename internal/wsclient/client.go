package wsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/happilymarrieddad/ws-api/internal/config"
	"github.com/happilymarrieddad/ws-api/types"
	"github.com/happilymarrieddad/ws-api/utils"

	"github.com/happilymarrieddad/interfaces/httpclient"
)

type tempType string
type weatherCondition string

const (
	Kelvin     tempType = "standard"
	Celcius    tempType = "metric"
	Fahrenheit tempType = "imperial"
)

const (
	Hot      weatherCondition = "Hot"
	Cold     weatherCondition = "Cold"
	Moderate weatherCondition = "Moderate"
)

const getWeatherAtLongLangURLFormat = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s"

type WeatherResponse struct {
	Temperature weatherCondition `json:"temperature"`
	Conditions  []*types.Weather
}

//go:generate mockgen -destination=./mocks/WSClient.go -package=mocks ws-api/interfaces/WSClient WSClient
type WSClient interface {
	GetWeatherDataAtLongLat(ctx context.Context, latitude, longitude float64, tt *tempType) (*WeatherResponse, error)
	GetWeatherAtLongLat(ctx context.Context, latitude, longitude float64, tt *tempType) (*types.GetWeatherResponse, error)
	GetWeatherCondition(tt tempType, tempVal float64) (weatherCondition, error)
}

func NewWSClient(cfg *config.Config, httpClient httpclient.HTTPClient) (WSClient, error) {
	if cfg == nil {
		return nil, errors.New("valid config required")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 30, // 30 sec timeout seems reasonable
		}
	}

	return &wsClient{
		apiKey:     cfg.OpenWeatherAPIKey,
		httpClient: httpClient,
	}, nil
}

type wsClient struct {
	apiKey string
	// we use this interface so that it can be mocked should it need to be
	httpClient httpclient.HTTPClient
}

func (c *wsClient) GetWeatherDataAtLongLat(ctx context.Context, latitude, longitude float64, tt *tempType) (*WeatherResponse, error) {
	if tt == nil {
		tt = utils.Ref(Fahrenheit)
	}

	res, err := c.GetWeatherAtLongLat(ctx, latitude, longitude, tt)
	if err != nil {
		return nil, err
	}

	cond, err := c.GetWeatherCondition(*tt, res.Main.Temp)
	if err != nil {
		return nil, err
	}

	return &WeatherResponse{
		Temperature: cond,
		Conditions:  res.Weather,
	}, nil
}

func (c *wsClient) GetWeatherAtLongLat(ctx context.Context, latitude, longitude float64, tt *tempType) (ws *types.GetWeatherResponse, err error) {
	if tt == nil {
		tt = utils.Ref(Fahrenheit)
	}

	url := fmt.Sprintf(getWeatherAtLongLangURLFormat, latitude, longitude, c.apiKey, *tt)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(resBody))
	ws = new(types.GetWeatherResponse)
	if err = json.Unmarshal(resBody, ws); err != nil {
		return nil, err
	}

	return ws, nil
}

// getWeatherCondition - converts to Fahrenheit to test
func (*wsClient) GetWeatherCondition(tt tempType, tempVal float64) (weatherCondition, error) {
	var valToTest float64

	switch tt {
	case Kelvin:
		valToTest = (tempVal-273.15)*9/5 + 32
	case Celcius:
		valToTest = (tempVal * 9 / 5) + 32
	case Fahrenheit:
		valToTest = tempVal
	default:
		return "", fmt.Errorf("invalid temp type '%s' valid=[Kelvin,Celcius,Fahrenheit]", tt)
	}

	if valToTest <= 45.0 {
		return Cold, nil
	} else if valToTest >= 70 {
		return Hot, nil
	}

	return Moderate, nil
}
