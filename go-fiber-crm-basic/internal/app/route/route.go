package route

import (
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/model"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", model.GetLeads)
	app.Get("/api/v1/lead/:id", model.GetLead)
	app.Post("/api/v1/lead", model.NewLead)
	app.Delete("/api/v1/lead/:id", model.DeleteLead)
}
