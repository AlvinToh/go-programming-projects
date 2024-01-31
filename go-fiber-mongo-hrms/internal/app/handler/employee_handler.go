package handler

import (
	"log"

	"github.com/alvintoh/go-programming-projects/go-fiber-mongo-hrms/internal/app/model"
	"github.com/gofiber/fiber/v2"
)

// swagger:response employeesResponse
type employeesResponseWrapper struct {
	// in:body
	Body []model.Employee
}

// swagger:response employeeResponse
type employeeResponseWrapper struct {
	// in:body
	Body model.Employee
}

// swagger:response genericError
type genericErrorWrapper struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// GetEmployees gets all employees.
// swagger:route GET /employee employees getEmployees
//
// Responses:
//
//	default: genericError
//	200: employeesResponse
func GetEmployees(c *fiber.Ctx) error {
	employees, err := model.GetAllEmployees()
	if err != nil {
		log.Printf("Get Employees error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	return c.JSON(employees)
}

// GetEmployee gets an employee by ID.
// swagger:route GET /employee/{id} employees getEmployee
//
// Responses:
//
//	default: genericError
//	200: employeeResponse
func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := model.GetEmployeeById(id)
	if err != nil {
		log.Printf("Get Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	return c.JSON(employee)
}

// NewEmployee creates a new employee.
// swagger:route POST /employee employees newEmployee
//
// Responses:
//
//	default: genericError
//	200: employeeResponse
func NewEmployee(c *fiber.Ctx) error {
	employee := new(model.Employee)
	if err := c.BodyParser(employee); err != nil {
		log.Printf("New Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	newEmployee, err := employee.CreateEmployee()
	if err != nil {
		log.Printf("Create Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	return c.JSON(newEmployee)
}

// UpdateEmployee updates an employee by ID.
// swagger:route PUT /employee/{id} employees updateEmployee
//
// Responses:
//
//	default: genericError
//	200: employeeResponse
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := model.GetEmployeeById(id)
	if err != nil {
		log.Printf("Get Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	if err := c.BodyParser(employee); err != nil {
		log.Printf("Update Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}

	// Update the employee
	err = employee.UpdateEmployee()
	if err != nil {
		log.Printf("Update Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}

	return c.JSON(employee)
}

// DeleteEmployee deletes an employee by ID.
// swagger:route DELETE /employee/{id} employees deleteEmployee
//
// Responses:
//
//	default: genericError
//	200: employeeResponse
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	err := model.DeleteEmployee(id)
	if err != nil {
		log.Printf("Delete Employee error: %v\n", err)
		c.Status(503).SendString(err.Error())
		return err
	}
	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}
