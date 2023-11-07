package route

import (
	"go-movies-crud/internal/app/handler"

	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	router.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")
}
