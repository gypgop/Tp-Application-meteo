package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"initial/models"
	"net/http"
	"time"
)

const baseURL = "https://geo.api.gouv.fr/"

var client = &http.Client{Timeout: 15 * time.Second}

var ErrCityNotFound = errors.New("Unknown player")

func GetCityCoords(CityCode []models.Code) (models.Datas, error) {

	fmt.Printf("-%d-\n", len(CityCode[0].Codes[0]))
	fmt.Printf("@%s@\n", CityCode[0].Codes[0])

	endpoint := fmt.Sprintf("%scommunes?codePostal=%s&fields=centre", baseURL, CityCode[0].Codes[0])

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Datas{}, err
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.Datas{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Datas{}, ErrCityNotFound

	}

	var CityCoords []models.Datas
	err = json.NewDecoder(response.Body).Decode(&CityCoords)
	return CityCoords[0], err

}

func GetCityCode(CityName string) ([]models.Code, error) {

	// communes?nom=Versailles&fields=codesPostaux

	endpoint := fmt.Sprintf("%scommunes?nom=%s&fields=codesPostaux", baseURL, CityName)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return nil, err
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, ErrCityNotFound

	}

	var CityCodes []models.Code

	err = json.NewDecoder(response.Body).Decode(&CityCodes)

	if err != nil {
		return nil, err
	}

	if len(CityCodes) == 0 {
		return nil, ErrCityNotFound
	}
	fmt.Printf("CityCode : %v", CityCodes)

	return CityCodes, err
}
