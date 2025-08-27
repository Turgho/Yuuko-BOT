package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// WeatherData representa o JSON que a API retorna (simplificado)
type WeatherData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

func GetWeather(city string) (*WeatherData, error) {
	// Busca coordenadas
	geoURL := fmt.Sprintf(
		"https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1",
		url.QueryEscape(city),
	)
	resp, err := http.Get(geoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geoData struct {
		Results []struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Name      string  `json:"name"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return nil, err
	}
	if len(geoData.Results) == 0 {
		return nil, fmt.Errorf("cidade não encontrada")
	}

	lat := geoData.Results[0].Latitude
	lon := geoData.Results[0].Longitude

	// Busca previsão
	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current_weather=true&temperature_unit=celsius&windspeed_unit=kmh", lat, lon)
	resp2, err := http.Get(weatherURL)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	var weather WeatherData
	if err := json.NewDecoder(resp2.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

// WeatherCodeToText retorna o emoji e a descrição do código do clima
func WeatherCodeToText(code int) string {
	switch code {
	case 0:
		return "☀️ Céu limpo"
	case 1:
		return "🌤️ Parcialmente nublado"
	case 2:
		return "⛅ Nublado"
	case 3:
		return "☁️ Nuvens densas"
	case 45:
		return "🌫️ Névoa"
	case 48:
		return "🌫️❄️ Névoa com cristais de gelo"
	case 51, 53, 55:
		return "🌦️ Chuvisco"
	case 61, 63, 65:
		return "🌧️ Chuva"
	case 71, 73, 75:
		return "❄️ Neve"
	case 80, 81, 82:
		return "🌧️☔ Pancadas de chuva"
	default:
		return "❓ Condição desconhecida"
	}
}
