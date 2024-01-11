package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete by Index with append, by ....
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", isbn: "10152", title: "Movie 1", Director: &Director{FirstName: "Anjas", LastName: "Fedo"}})
	movies = append(movies, Movie{ID: "2", isbn: "10505", title: "Movie 2", Director: &Director{FirstName: "Anjas", LastName: "Fedo"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server on Port 8000/n")

	log.Fatal(http.ListenAndServe(":8000", r))

}
