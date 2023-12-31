package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/util"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/model"

	"github.com/gorilla/mux"
)

var NewBook model.Book

// BookResponse is the response object for book API
// swagger:response bookResponse
type BookResponse struct {
	// in:body
	Body model.Book
}

// BooksResponse is the response object for books API
// swagger:response booksResponse
type BooksResponse struct {
	// in:body
	Body []model.Book
}

// BookIDParameter is the bookId parameter in the URL
// swagger:parameters getBookById deleteBook updateBook
type BookIDParameter struct {
	// The id of the book
	// in:path
	// required:true
	BookID int64 `json:"bookId"`
}

// BookRequestBody is the request body for creating/updating a book
// swagger:parameters createBook updateBook
type BookRequestBody struct {
	// in:body
	Body model.Book
}

// GetBook swagger:route GET /books books getBooks
//
// # Get Books
//
// This will get the details of all books.
//
//	Responses:
//	  200: booksResponse
func GetBook(w http.ResponseWriter, r *http.Request) {
	books, _ := model.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById swagger:route GET /books/{bookId} books getBookById
//
// # Get Book By Id
//
// This will get the details of a book by id.
//
//	Responses:
//	  200: bookResponse
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook swagger:route POST /books books createBook
//
// # Create Book
//
// This will create a new book.
//
//	Responses:
//	  200: bookResponse
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	util.ParseBody(r, CreateBook)
	book, _ := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook swagger:route DELETE /books/{bookId} books deleteBook
//
// # Delete Book
//
// This will delete a book by id.
//
//	Responses:
//	  200: bookResponse
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := model.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook swagger:route PUT /books/{bookId} books updateBook
//
// # Update Book
//
// This will update a book by id.
//
//	Responses:
//	  200: bookResponse
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &model.Book{}
	util.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := model.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
