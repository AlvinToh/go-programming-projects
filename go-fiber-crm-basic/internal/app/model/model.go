package model

import (
	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with ID"))
		return nil
	}
	db.Delete(&lead)
	return c.Send([]byte("Lead deleted successfully"))
}
