package model

import (
	"log"

	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// swagger:parameters lead
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

// swagger:response leadsResponse
type leadsResponseWrapper struct {
	// in:body
	Body []Lead
}

// swagger:response leadResponse
type leadResponseWrapper struct {
	// in:body
	Body Lead
}

// swagger:response genericError
type genericErrorWrapper struct {
	code    string `json:"code"`
	message string `json:"message"`
	details string `json:"email"`
}

// GetLeads gets all leads.
// swagger:route GET /leads leads getLeads
//
// Responses:
//
//	default: genericError
//	200: leadsResponse
func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	log.Printf("Get Leads: %v\n", leads)
	return c.JSON(leads)
}

// GetLead gets a lead by ID.
// swagger:route GET /leads/{id} leads getLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	log.Printf("Get Lead for id: %v %v\n", id, lead)
	return c.JSON(lead)
}

// NewLead creates a new lead.
// swagger:route POST /leads leads newLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		log.Printf("New Lead error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	db.Create(&lead)
	log.Printf("New Lead: %v\n", lead)
	return c.JSON(lead)
}

// DeleteLead deletes a lead by ID.
// swagger:route DELETE /leads/{id} leads deleteLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		log.Printf("Delete Lead error: No lead found with ID %v\n", id)
		c.Status(500).Send([]byte("No lead found with ID"))
		return nil
	}
	db.Delete(&lead)
	log.Printf("Delete Lead: %v\n", lead)
	return c.Send([]byte("Lead deleted successfully"))
}
