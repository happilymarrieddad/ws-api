package wsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/happilymarrieddad/ws-api/internal/config"
	"github.com/happilymarrieddad/ws-api/types"
	"github.com/happilymarrieddad/ws-api/utils"

	"github.com/happilymarrieddad/interfaces/httpclient"
)

type TempType string
type WeatherCondition string

const (
	Kelvin     TempType = "standard"
	Celcius    TempType = "metric"
	Fahrenheit TempType = "imperial"
)

const (
	Hot      WeatherCondition = "hot"
	Cold     WeatherCondition = "cold"
	Moderate WeatherCondition = "moderate"
)

const getWeatherAtLongLangURLFormat = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s"

type WeatherResponse struct {
	TempFeelsLike WeatherCondition `json:"temperature_feels_like"`
	Temperature   float64          `json:"temp"`
	Conditions    []*types.Weather `json:"conditions_and_alerts"`
}

//go:generate mockgen -destination=./mocks/WSClient.go -package=mocks github.com/happilymarrieddad/ws-api/internal/wsclient WSClient
type WSClient interface {
	GetWeatherDataAtLongLat(ctx context.Context, latitude, longitude float64, tt *TempType) (*WeatherResponse, error)
	GetWeatherAtLongLat(ctx context.Context, latitude, longitude float64, tt *TempType) (*types.GetWeatherResponse, error)
	GetWeatherConditonFromTempType(tt TempType, tempVal float64) (WeatherCondition, error)
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

func (c *wsClient) GetWeatherDataAtLongLat(ctx context.Context, latitude, longitude float64, tt *TempType) (*WeatherResponse, error) {
	if tt == nil {
		tt = utils.Ref(Fahrenheit)
	}

	res, err := c.GetWeatherAtLongLat(ctx, latitude, longitude, tt)
	if err != nil {
		return nil, err
	}

	cond, err := c.GetWeatherConditonFromTempType(*tt, res.Main.Temp)
	if err != nil {
		return nil, err
	}

	return &WeatherResponse{
		TempFeelsLike: cond,
		Temperature:   res.Main.Temp,
		Conditions:    res.Weather,
	}, nil
}

func (c *wsClient) GetWeatherAtLongLat(ctx context.Context, latitude, longitude float64, tt *TempType) (ws *types.GetWeatherResponse, err error) {
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
		return nil, err
	}

	ws = new(types.GetWeatherResponse)
	if err = json.Unmarshal(resBody, ws); err != nil {
		return nil, err
	}

	if ws.Code/100 != 2 {
		return nil, errors.New(ws.ErrMessage)
	}

	return ws, nil
}

// GetWeatherConditonFromTempType converts to Fahrenheit from temp type and returns the weather condition
func (*wsClient) GetWeatherConditonFromTempType(tt TempType, tempVal float64) (WeatherCondition, error) {
	var valToTest float64

	// TODO: do not use literals here... need to create consts at some point
	//		actually.. probably should make helper funcs here that are tested
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

func ValidateTempType(tt *TempType) error {
	if tt == nil {
		return nil // client will just default so don't worry about it
	}

	switch *tt {
	case Celcius, Fahrenheit, Kelvin:
		return nil
	}

	return errors.New("invalid temp type")
}
