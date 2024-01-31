package handler

import (
	"log"
	"strconv"

	"github.com/alvintoh/go-programming-projects/go-fiber-crm-basic/internal/app/model"
	"github.com/gofiber/fiber/v2"
)

// swagger:response leadsResponse
type leadsResponseWrapper struct {
	// in:body
	Body []model.Lead
}

// swagger:response leadResponse
type leadResponseWrapper struct {
	// in:body
	Body model.Lead
}

// swagger:response genericError
type genericErrorWrapper struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// GetLeads gets all leads.
// swagger:route GET /lead leads getLeads
//
// Responses:
//
//	default: genericError
//	200: leadsResponse
func GetLeads(c *fiber.Ctx) error {
	var leads []model.Lead
	leads, _ = model.GetAllLeads()
	return c.JSON(leads)
}

// GetLead gets a lead by ID.
// swagger:route GET /lead/{id} leads getLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	leadID, _ := strconv.ParseInt(id, 0, 0)
	lead, _ := model.GetLeadById(leadID)
	return c.JSON(lead)
}

// NewLead creates a new lead.
// swagger:route POST /lead leads newLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func NewLead(c *fiber.Ctx) error {
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		log.Printf("New Lead error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	lead.CreateLead()
	return c.JSON(lead)
}

// DeleteLead deletes a lead by ID.
// swagger:route DELETE /lead/{id} leads deleteLead
//
// Responses:
//
//	default: genericError
//	200: leadResponse
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	leadID, _ := strconv.ParseInt(id, 0, 0)
	_ = model.DeleteLead(leadID)
	return c.JSON(fiber.Map{"message": "Lead deleted successfully"})
}
