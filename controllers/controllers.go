package controllers

import (
	"fmt"
	"initial/api"
	"net/http"
	"strings"
)

func MeteoPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG: Requête reçue sur /menu")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		fmt.Println("DEBUG: Méthode non autorisée :", r.Method)
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	cityName := strings.TrimSpace(r.URL.Query().Get("CityName"))
	if cityName == "" {
		fmt.Println("DEBUG: Paramètre PlayerName manquant ou vide")
		http.Error(w, "Veuillez indiquer un nom de ville", http.StatusBadRequest)
		return
	}

	cityCode, err := api.GetCityCode(cityName)
	// ne recup apaas le city code

	cityCoords, err := api.GetCityCoords(cityCode)

	cityDatas, err := api.GetOpenMeteo(cityCoords)

	fmt.Fprintf(w, "ville trouvé : %s\n %+v %v", cityName, cityDatas, err)

}
