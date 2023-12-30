package main

import (
	"fmt"
	"log"

	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/model"
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/route"
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/platform/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDatabase()
	database.DBConn.AutoMigrate(&model.Lead{})
	fmt.Println("Database Migrated")

	route.SetupRoutes(app)
	// Serve Swagger UI and Swagger JSON
	app.Static("/swagger/ui", "./swagger/ui")
	app.Static("/swagger", "./swagger")

	log.Fatal(app.Listen(":3000"))
	defer database.DBConn.Close()
}
