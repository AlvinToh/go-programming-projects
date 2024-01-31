package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setup(t *testing.T) (sqlmock.Sqlmock, *fiber.App, func()) {
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

	// Create a new Fiber app
	app := fiber.New()

	// Register the GetLeads route
	app.Get("/api/v1/lead", GetLeads)

	// Return the mock, app and a function that restores the original datastore and closes the database when defer teardown is called
	return mock, app, func() {
		model.SetDB(originalDB)
		db.Close()
	}
}

func TestGetLeads(t *testing.T) {
	mock, app, teardown := setup(t)
	defer teardown()

	// Set up the mock to return a single lead when GetLeads is called
	rows := sqlmock.NewRows([]string{"ID", "Name", "Company", "Email", "Phone"}).
		AddRow(1, "Test Lead", "Test Company", "test@email.com", "1234567890")
	mock.ExpectQuery("^SELECT (.+) FROM \"leads\" WHERE \"leads\".\"deleted_at\" IS NULL$").WillReturnRows(rows)

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/lead", nil))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "handler returned wrong status code")

	// Check if there are any unfulfilled expectations
	require.NoError(t, mock.ExpectationsWereMet())

	// Parse the response body
	var leads []model.Lead
	err = json.NewDecoder(resp.Body).Decode(&leads)
	require.NoError(t, err)

	// Check that the length of leads is 1
	require.Len(t, leads, 1, "handler returned unexpected number of leads")

	// Check the fields of the first lead
	require.Equal(t, "Test Lead", leads[0].Name, "handler returned unexpected lead name")
	require.Equal(t, "Test Company", leads[0].Company, "handler returned unexpected lead company")
	require.Equal(t, "test@email.com", leads[0].Email, "handler returned unexpected lead email")
	require.Equal(t, "1234567890", leads[0].Phone, "handler returned unexpected lead phone")
}
