package route

import (
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", handler.GetLeads)
	app.Get("/api/v1/lead/:id", handler.GetLead)
	app.Post("/api/v1/lead", handler.NewLead)
	app.Delete("/api/v1/lead/:id", handler.DeleteLead)
}
