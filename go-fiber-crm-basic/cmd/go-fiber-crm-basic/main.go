package main

import (
	"fmt"
	"log"

	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/database"
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/model"
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDatabase()
	database.DBConn.AutoMigrate(&model.Lead{})
	fmt.Println("Database Migrated")

	route.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	defer database.DBConn.Close()
}
