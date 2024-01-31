package route

import (
	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/handler"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", handler.CreateBook).Methods("POST")
	router.HandleFunc("/book/", handler.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", handler.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", handler.DeleteBook).Methods("DELETE")
}
