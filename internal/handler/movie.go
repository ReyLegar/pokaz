package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	movieID, err := h.service.AddMovie(movie)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie added successfully. ID: %d\n", movieID)
}

func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	movieID, err := strconv.Atoi(r.URL.Path[len("/api/movie/update/"):])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateMovie(movieID, movie); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Movie updated successfully. ID: %d\n", movieID)
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID, err := strconv.Atoi(r.URL.Path[len("/api/movie/delete/"):])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteMovie(movieID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Movie deleted successfully. ID: %d\n", movieID)
}

func (h *Handler) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.service.GetAllMovies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch movies: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (h *Handler) searchMoviesByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.service.SearchMoviesByTitle(title)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search movies by title: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (h *Handler) searchMoviesByActorName(w http.ResponseWriter, r *http.Request) {
	actorName := r.URL.Query().Get("actorName")
	if actorName == "" {
		http.Error(w, "Actor name parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.service.SearchMoviesByActorName(actorName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search movies by actor name: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
