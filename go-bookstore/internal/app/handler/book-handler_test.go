package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/model"
	"github.com/stretchr/testify/require"
)

func setup() func() {
	// Mock the GetAllBooks function to return a predictable response
	getBooks := func() []model.Book {
		return []model.Book{
			{
				Name:        "Test Book",
				Author:      "Test Author",
				Publication: "Test Publication",
			},
		}
	}
	originalGetAllBooks := model.GetAllBooks
	model.GetAllBooks = getBooks

	// Return a function that restores the original GetAllBooks function
	return func() {
		model.GetAllBooks = originalGetAllBooks
	}
}

func TestGetBook(t *testing.T) {
	teardown := setup()
	defer teardown()

	req := httptest.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()

	http.HandlerFunc(GetBook).ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Parse the response body
	var books []model.Book
	var err error = json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the length of books is 1
	require.Len(t, books, 1, "handler returned unexpected number of books")

	// Check the fields of the first book
	require.Equal(t, "Test Book", books[0].Name, "handler returned unexpected book name")
	require.Equal(t, "Test Author", books[0].Author, "handler returned unexpected book author")
	require.Equal(t, "Test Publication", books[0].Publication, "handler returned unexpected book publication")
}
