package main

import (
	"fmt"
	"initial/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("DEBUG: execution du main\n")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/menu.html")
	})

	http.HandleFunc("/City", controllers.MeteoPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
