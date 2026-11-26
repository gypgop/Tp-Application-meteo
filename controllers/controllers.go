package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"initial/api"
	"log"
	"net/http"
)

func MeteoPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG: Requête reçue sur /menu")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		fmt.Println("DEBUG: Méthode non autorisée :", r.Method)
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	cityName := r.URL.Query().Get("CityName")
	if cityName == "" {
		fmt.Println("DEBUG: Paramètre CityName manquant ou vide")
		http.Error(w, "Veuillez indiquer un nom de ville", http.StatusBadRequest)
		return
	}

	cityCode, err := api.GetCityCode(cityName)

	cityCoords, err := api.GetCityCoords(cityCode)

	cityDatas, err := api.GetOpenMeteo(cityCoords)

	fmt.Printf("\n\n\n\nville trouvé : %s\n %+v %v", cityName, cityDatas, err)

	page, err := template.ParseFiles("templates/meteo.html")

	if err != nil {
		log.Printf("impossible de charger la page Pokémon : %v", err)
		renderError(w, http.StatusInternalServerError, "Impossible de charger la page de la ville demandé")
		return
	}

	var body bytes.Buffer

	if err := page.Execute(&body, cityDatas); err != nil {
		log.Printf("impossible de générer la page de cette ville : %v", err)
		renderError(w, http.StatusInternalServerError, "Impossible de générer la page du bla bla.")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	if _, err := body.WriteTo(w); err != nil {
		log.Printf("impossible d'envoyer la page Pokémon : %v", err)
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	page, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Printf("impossible de charger la page d'erreur : %v", err)
		http.Error(w, "Erreur interne du serveur.", http.StatusInternalServerError)
		return
	}

	var body bytes.Buffer
	if err := page.Execute(&body, struct {
		StatusCode int
		Message    string
	}{StatusCode: status, Message: message}); err != nil {
		log.Printf("impossible de générer la page d'erreur : %v", err)
		http.Error(w, "Erreur interne du serveur.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err := body.WriteTo(w); err != nil {
		log.Printf("impossible d'envoyer la page d'erreur : %v", err)
	}
}
