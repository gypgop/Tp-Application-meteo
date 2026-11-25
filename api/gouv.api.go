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

var client = &http.Client{Timeout: 5 * time.Second}

var ErrCityNotFound = errors.New("Unknown player")

func GetCityCoords(CityCode int) (models.Datas, error) {

	endpoint := fmt.Sprintf("%scommunes?codePostal=%d&fields=centre", baseURL, CityCode)

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

func GetCityCode(CityCode string) (int, error) {

	// https://geo.api.gouv.fr/communes?nom=Versailles&fields=codesPostaux,nom,centre

	endpoint := fmt.Sprintf("%s%s&fields=centre", baseURL, CityCode)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Printf("erreur, %v", err)
		return -1, err
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return -1, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return -1, ErrCityNotFound

	}

	var CityCodes int

	err = json.NewDecoder(response.Body).Decode(&CityCodes)
	fmt.Printf("CityCode : %d", CityCodes)
	return CityCodes, err

}
