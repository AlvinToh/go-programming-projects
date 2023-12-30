package main

import (
	"log"
	"net/http"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/route"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve Swagger UI and Swagger JSON
	r.PathPrefix("/swagger/ui/").Handler(http.StripPrefix("/swagger/ui/", http.FileServer(http.Dir("./swagger/ui"))))
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9001", r))
}
