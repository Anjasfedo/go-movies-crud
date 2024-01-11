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
	ID       string    `json: "id"`
	isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json: "firstName`
	LastName  string `json: "lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server on Port 8000/n")

	log.Fatal(http.ListenAndServe(":8000", r))

}
