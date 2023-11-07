package main

import (
	"fmt"
	"log"
	"net/http"

	"go-movies-crud/internal/app/route"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	route.RegisterMovieRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
