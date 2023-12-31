package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/model"
)

func setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	// Create a sqlmock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	// Pass it to gorm
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	// Save the original datastore
	originalDB := model.GetDB()

	// Replace the datastore with the mock
	model.SetDB(gormDB)

	// Return the mock and a function that restores the original datastore and closes the database when defer teardown is called
	return gormDB, mock, func() {
		model.SetDB(originalDB)
		db.Close()
	}
}

func TestGetAllBooks(t *testing.T) {
	_, mock, teardown := setup(t)
	defer teardown()

	// Set up the mock to return a single book when GetAllBooks is called
	rows := sqlmock.NewRows([]string{"ID", "Name", "Author", "Publication"}).
		AddRow(1, "Test Book", "Test Author", "Test Publication")
	mock.ExpectQuery("^SELECT (.+) FROM \"books\" WHERE \"books\".\"deleted_at\" IS NULL$").WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()

	http.HandlerFunc(GetBook).ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Parse the response body
	var books []model.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	require.NoError(t, err)

	// Check that the length of books is 1
	require.Len(t, books, 1, "handler returned unexpected number of books")

	// Check the fields of the first book
	require.Equal(t, "Test Book", books[0].Name, "handler returned unexpected book name")
	require.Equal(t, "Test Author", books[0].Author, "handler returned unexpected book author")
	require.Equal(t, "Test Publication", books[0].Publication, "handler returned unexpected book publication")
}
