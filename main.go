package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int16  `json:"year"`
}

type Movies []Movie

func allMovies(w http.ResponseWriter, r *http.Request) {
	articles := Movie{
		Title:       "Teste",
		Description: "Descrição",
		Year:        2018,
	}
	fmt.Println("Endpoint all movies!")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRquests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", allMovies)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRquests()
}
