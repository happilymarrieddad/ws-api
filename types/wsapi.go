package types

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	// The 4 below might be integers but we go with float just in case
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	SeaLevel    float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_level"`
}

type Wind struct {
	Speed  float64 `json:"speed"`
	Degree float64 `json:"degree"`
	Gust   float64 `json:"gust"`
}

type Clouds struct {
	All float64 `json:"all"`
}

type System struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	// Timestamps
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

type GetWeatherResponse struct {
	Coordinates *Coordinates `json:"coord"`
	Weather     []*Weather   `json:"weather"`
	Base        string       `json:"base"`
	Main        *Main        `json:"main"`
	Visibility  float64      `json:"visibility"`
	Wind        *Wind        `json:"wind"`
	Clouds      *Clouds      `json:"clouds"`
	DT          int64        `json:"dt"` // don't know what this is some sort of timestamp
	System      *System      `json:"sys"`
	Timezome    int64        `json:"timezone"`
	ID          int64        `json:"id"`
	City        string       `json:"name"`
	Code        int          `json:"code"`
}
