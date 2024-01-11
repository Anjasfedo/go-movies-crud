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

// Movie struct represents the structure of a movie with JSON tags
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents the structure of a director with JSON tags
type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Slice to store movies
var movies []Movie

// Handler function to get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Encode movies slice to JSON and send it in the response
	json.NewEncoder(w).Encode(movies)
}

// Handler function to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request URL
	params := mux.Vars(r)

	// Iterate over movies to find the movie with the specified ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the movie by slicing the movies slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Encode the updated movies slice to JSON and send it in the response
	json.NewEncoder(w).Encode(movies)
}

// Handler function to get a movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request URL
	params := mux.Vars(r)

	// Iterate over movies to find the movie with the specified ID
	for _, item := range movies {
		if item.ID == params["id"] {
			// Encode the found movie to JSON and send it in the response
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Handler function to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Create a new movie instance and decode the request body into it
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Generate a random ID for the new movie
	movie.ID = strconv.Itoa(rand.Intn(10000000))

	// Append the new movie to the movies slice
	movies = append(movies, movie)

	// Encode the new movie to JSON and send it in the response
	json.NewEncoder(w).Encode(movie)
}

// Handler function to update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")

	// Get the parameters from the request URL
	params := mux.Vars(r)

	// Iterate over movies to find the movie with the specified ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the existing movie by slicing the movies slice
			movies = append(movies[:index], movies[index+1:]...)

			// Create a new movie instance and decode the request body into it
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			// Set the ID of the new movie to the original ID
			movie.ID = params["id"]

			// Append the updated movie to the movies slice
			movies = append(movies, movie)

			// Encode the updated movie to JSON and send it in the response
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	// Create a new router instance from the gorilla/mux package
	r := mux.NewRouter()

	// Initialize movies slice with sample data
	movies = append(movies, Movie{ID: "1", Isbn: "10152", Title: "Movie 1", Director: &Director{FirstName: "Anjas", LastName: "Fedo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "10505", Title: "Movie 2", Director: &Director{FirstName: "Anjas", LastName: "Fedo"}})

	// Define routes and corresponding handler functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server on port 8000
	fmt.Printf("Starting Server on Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
