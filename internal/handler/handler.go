package handler

import (
	"net/http"

	"github.com/ReyLegar/vkTestProject/internal/service"
)

type Handler struct {
	service *service.Service
}

// Вот здесь фигня, надо исправить
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	/* 	case "/movie/add":
	   		h.service.AddMovie(models.Movie{})
	   	case "/movie/update/":
	   		h.service.UpdateMovie()
	   	case "/movie/delete/":
	   		h.service.DeleteMovie(w, r)
	   	case "/movie/getAll":
	   		h.service.GetAllMovies(w, r)
	   	case "/movie/searchByTitle":
	   		h.service.SearchMoviesByTitle(w, r)
	   	case "/movie/searchByActorName":
	   		h.service.SearchMoviesByActorName(w, r) */
	case "/actor/add":
		h.AddActor(w, r)
	case "/actor/update":
		h.UpdateActor(w, r)
	case "/actor/delete":
		h.DeleteActor(w, r)
	default:
		http.NotFound(w, r)
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() {

	/* http.HandleFunc("/auth/sign-up", h.signUp)
	   http.HandleFunc("/auth/sign-in", h.signIn) */

	http.HandleFunc("/movie/add", h.addMovie)
	http.HandleFunc("/movie/update/", h.updateMovie)
	http.HandleFunc("/movie/delete/", h.deleteMovie)
	http.HandleFunc("/movie/getAll", h.getAllMovies)
	http.HandleFunc("/movie/searchByTitle", h.searchMoviesByTitle)
	http.HandleFunc("/movie/searchByActorName", h.searchMoviesByActorName)

	http.HandleFunc("/actor/add", h.AddActor)
	http.HandleFunc("/actor/update", h.UpdateActor)
	http.HandleFunc("/actor/delete", h.DeleteActor)
}
