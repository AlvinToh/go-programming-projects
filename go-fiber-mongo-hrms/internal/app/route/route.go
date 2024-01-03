package route

import (
	"github.com/alvintoh/go-programming-projects/go-fiber-mongo-hrms/internal/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/employee", handler.GetEmployees)
	app.Get("/api/v1/employee/:id", handler.GetEmployee)
	app.Post("/api/v1/employee", handler.NewEmployee)
	app.Put("/api/v1/employee/:id", handler.UpdateEmployee)
	app.Delete("/api/v1/employee/:id", handler.DeleteEmployee)
}
