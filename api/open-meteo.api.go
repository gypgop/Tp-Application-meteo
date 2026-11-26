package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"initial/models"
	"net/http"
	"time"
)

const openMeteoBaseURL = "https://api.open-meteo.com/v1/forecast?"

// &current=temperature_2m,wind_speed_10m,weather_code&hourly=temperature_2m,precipitation_probability&daily=temperature_2m_max,temperature_2m_min,precipitation_sum&weather_code&timezone=auto&forecast_days=7
var openMeteoClient = &http.Client{Timeout: 15 * time.Second}

var ErrOpenMeteoCityNotFound = errors.New("Unknown player")

func GetOpenMeteo(CityCoords models.Datas) (models.Reponse, error) {

	fmt.Printf("\ncityCoords = @%v@\n", CityCoords.Centre.Coordinates[0])
	fmt.Printf("\ncityCoords = @%v@\n", CityCoords.Centre.Coordinates[1])

	endpoint := fmt.Sprintf("%slatitude=%f&longitude=%f&current=temperature_2m,wind_speed_10m,weather_code&hourly=temperature_2m,precipitation_probability&daily=temperature_2m_max,temperature_2m_min,precipitation_sum&weather_code&timezone=auto&forecast_days=7", openMeteoBaseURL, CityCoords.Centre.Coordinates[0], CityCoords.Centre.Coordinates[1])
	// 													^ a changer

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Reponse{}, err
	}

	// fmt.Printf("%v\n", request)

	response, err := openMeteoClient.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Reponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Reponse{}, ErrOpenMeteoCityNotFound

	}

	var datas models.Reponse

	err = json.NewDecoder(response.Body).Decode(&datas)

	fmt.Printf("%v", datas)

	return datas, err

}
