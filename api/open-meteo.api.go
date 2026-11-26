package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"initial/models"
	"net/http"
	"time"
)

const openMeteoBaseURL = "https://geo.api.gouv.fr/"

var openMeteoClient = &http.Client{Timeout: 15 * time.Second}

var ErrOpenMeteoCityNotFound = errors.New("Unknown player")

func GetOpenMeteo(CityCoords []models.Datas) (models.Reponse, error) {

	fmt.Printf("@%s@\n", CityCoords)

	endpoint := fmt.Sprintf("%scommunes?codePostal=%s&fields=centre", openMeteoBaseURL, CityCode[0].Codes[0])

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Datas{}, err
	}

	fmt.Printf("%v\n", request)

	response, err := openMeteoClient.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Datas{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Datas{}, ErrOpenMeteoCityNotFound

	}

	// var CityCoords []models.Datas
	err = json.NewDecoder(response.Body).Decode(&CityCoords)
	return CityCoords[0], err

}
