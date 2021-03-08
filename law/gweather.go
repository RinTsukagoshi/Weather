package law

import (
	"encoding/json"
	"fmt"
	"os"
)

type Matome struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Coord   Coord     `json:"coord"`
	Wind    Wind      `json:"wind"`
	Dt      int64     `json:"dt"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressuer int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

func WtoD(s string) Matome {
	result := Matome{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return result

}
func Honyaku(s string) string {
	switch s {
	case "Clouds":
		return "曇り"
	case "Snow":
		return "雪"
	case "Clear":
		return "晴天"
	case "Drizzle":
		return "霧雨"
	case "Thunderstorm":
		return "雷雨"
	case "Rain":
		return "雨"
	default:
		return "わかりません"
	}
}
