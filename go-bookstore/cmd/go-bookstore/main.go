package main

import (
	"log"
	"net/http"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/route"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9001", r))
}
